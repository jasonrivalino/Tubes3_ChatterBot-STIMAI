package searchanswer

import (
	"backend/component/checkdateformat"
	"backend/component/mathoperation"
	"backend/controllers/listquestioncontroller"
	"regexp"
	"strconv"
)

func SearchAnswer(question string) string {
	pattern := regexp.MustCompile(`^\d{1,2}/\d{1,2}/\d{1,4}$`) // Regex untuk mengecek apakah pertanyaan merupakan format tanggal

	// Jika pertanyaan merupakan format tanggal
	if pattern.MatchString(question) {
		// Cek lagi apabila format tanggalnya valid
		if checkdateformat.IsValidDate(question) {
			// Jika benar, maka panggil fungsi searchAnswerByDate
			return "Hari " + checkdateformat.GetDayOfWeek(question)
		}
		return "Maaf, format tanggal tidak valid"
	}

	pattern = regexp.MustCompile(`^[0-9\s\*\+\-\/\^\(\)]*\??$`) // Regex untuk mengecek apakah pertanyaan merupakan format operasi matematika

	// Jika pertanyaan merupakan format operasi matematika secara postfix
	if pattern.MatchString(question) {
		// Cek lagi apabila format postfixnya valid
		if mathoperation.IsInfixValid(question) {
			// Ubah notasi infix menjadi postfix
			question = mathoperation.InfixToPostfix(question)

			// Jika benar, maka hitung hasilnya
			ans, err := mathoperation.CalculatePostfix(question)
			if err != nil {
				return "Maaf, terjadi kesalahan pada sistem"
			}
			// Jika tidak terjadi kesalahan, maka kembalikan hasil perhitungan
			return strconv.FormatFloat(ans, 'f', -1, 64)
		}
		return "Maaf, format operasi matematika tidak valid"
	}

	pattern = regexp.MustCompile(`(?i)tambahkan pertanyaan (.+) dengan jawaban (.+)`)

	// Jika pertanyaan merupakan format untuk menambahkan pertanyaan
	if match := pattern.FindStringSubmatch(question); len(match) > 0 {

		// Cek dahulu apakah sebelumnya ada pertanyaan yang sama di dalam database
		data, err := listquestioncontroller.SearchQuestionOnDatabase(match[1])

		switch err {

		case nil:
			// Kasus pertanyaan sudah ada di dalam database sebelumnya
			if data.Id == 0 {
				if err := listquestioncontroller.AddNewQuestion(match[1], match[2]); err != nil {
					return "Maaf, tidak dapat menambahkan pertanyaan"
				}
				return "Pertanyaan" + match[1] + "berhasil ditambahkan"
			}

			// Kasus pertanyaan belum ada di dalam database sebelumnya
			if err := listquestioncontroller.UpdateAnswer(data, match[2]); err != nil {
				return "Maaf, tidak dapat menambahkan pertanyaan"
			}

			return strconv.Itoa(int(data.Id)) + "Pertanyaan " + match[1] + " sudah ada! Jawaban akan diupdate ke " + match[2]

		// Kasus error
		default:
			return "Maaf, tidak dapat menambahkan pertanyaan"
		}

	}

	pattern = regexp.MustCompile(`(?i)hapus pertanyaan (.+)`)

	// Jika pertanyaan merupakan format untuk menghapus pertanyaan
	if match := pattern.FindStringSubmatch(question); len(match) > 0 {
		if err := listquestioncontroller.DeleteQuestion(match[1]); err != nil {
			return "Maaf, tidak dapat menghapus pertanyaan"
		}
		return "Pertanyaan" + match[1] + "berhasil dihapus"
	}

	// Jika tidak, maka cek apakah pertanyaan merupakan format operasi matematika secara prefix
	return "Maaf, saya tidak mengerti pertanyaan Anda"
}
