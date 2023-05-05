package searchanswer

import (
	"backend/component/bm"
	"backend/component/checkdateformat"
	"backend/component/kmp"
	"backend/component/lcs"
	"backend/component/mathoperation"
	"backend/controllers/listquestioncontroller"
	"regexp"
	"sort"
	"strconv"
)

func SearchAnswer(question string, stringMatching int8) string {
	// Jika stringMatching tidak bernilai 1 atau 2, maka kembalikan pesan error
	// if stringMatching != 1 || stringMatching != 2 {
	// 	return "Maaf, pencocokan string tidak valid"
	// }

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
				return "Pertanyaan " + match[1] + " berhasil ditambahkan"
			}

			// Kasus pertanyaan belum ada di dalam database sebelumnya
			if err := listquestioncontroller.UpdateAnswer(data, match[2]); err != nil {
				return "Maaf, tidak dapat menambahkan pertanyaan"
			}

			return "Pertanyaan " + match[1] + " sudah ada! Jawaban akan diupdate ke " + match[2]

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
		return "Pertanyaan " + match[1] + " berhasil dihapus"
	}

	// Jika input bukan merupakan format tanggal, operasi matematika, maupun perintah untuk menambahkan/hapus pertanyaan
	// Maka lakukan pencarian pertanyaan yang exact match di dalam database dengan Algoritma KMP atau Boyer-Moore
	listQuestion, err := listquestioncontroller.GetAllListQuestion()

	switch err {
	case nil:
		// Jika tidak terjadi error, maka lakukan pencarian pertanyaan yang exact match di dalam database dengan Algoritma KMP atau Boyer-Moore
		for _, data := range listQuestion {
			if stringMatching == 1 {
				if len(question) > len(data.Pertanyaan) {
					if kmp.KMPSearch(data.Pertanyaan, question) >= 0 {
						return data.Jawaban
					}
				}
				if kmp.KMPSearch(question, data.Pertanyaan) >= 0 {
					return data.Jawaban
				}
			} else if stringMatching == 2 {
				if len(question) < len(data.Pertanyaan) {
					if bm.BoyerMooreSearch(data.Pertanyaan, question) >= 0 {
						return data.Jawaban
					}
				}
				if bm.BoyerMooreSearch(question, data.Pertanyaan) >= 0 {
					return data.Jawaban
				}
			}

		}
	default:
		return "Maaf, terjadi kesalahan pada sistem"
	}

	// Jika tidak menemukan pertanyaan yang cocok, maka cari pertanyaan yang nilai kemripiannya paling tinggi
	// nilai kemiripan diantara 0 sampai 1 dan nilai kemiripan dihitung dengan LCS
	// Jika nilai kemiripan lebih dari 0.9, maka kembalikan jawaban pertanyaan tersebut

	type ListQuestionSimilarity struct {
		Id         int64   `json:"id"`
		Pertanyaan string  `json:"pertanyaan"`
		Jawaban    string  `json:"jawaban"`
		Similarity float64 `json:"similarity"`
	}

	var listQuestionSimilarity []ListQuestionSimilarity
	// Samakan data Id, Pertanyaan, dan Jawaban dari listQuestion ke listQuestionSimilarity
	for _, data := range listQuestion {
		listQuestionSimilarity = append(listQuestionSimilarity, ListQuestionSimilarity{Id: data.Id, Pertanyaan: data.Pertanyaan, Jawaban: data.Jawaban})
	}

	// Hitung nilai kemiripan pertanyaan dengan pertanyaan yang ada di dalam database
	for i := 0; i < len(listQuestionSimilarity); i++ {
		listQuestionSimilarity[i].Similarity = lcs.Similarity(question, listQuestionSimilarity[i].Pertanyaan)
	}

	// Sorting listQuestionSimilarity berdasarkan nilai kemiripan dari yang tertinggi ke terendah
	sort.Slice(listQuestionSimilarity, func(i, j int) bool {
		return listQuestionSimilarity[i].Similarity > listQuestionSimilarity[j].Similarity
	})

	// Jika nilai kemiripan pertanyaan dengan pertanyaan yang ada di dalam database lebih dari 0.9, maka kembalikan jawaban pertanyaan tersebut
	if listQuestionSimilarity[0].Similarity > 0.9 {
		return listQuestionSimilarity[0].Jawaban
	}

	// Jika tidak menemukan pertanyaan yang cocok, maka kembalikan jawaban default
	result := "Maaf, pertanyaan Anda tidak ada di dalam database.\n"
	result += "Apakah maksud Anda:\n"
	result += "1. " + listQuestionSimilarity[0].Pertanyaan + "\n"
	result += "2. " + listQuestionSimilarity[1].Pertanyaan + "\n"
	result += "3. " + listQuestionSimilarity[2].Pertanyaan + "\n"

	return result
}
