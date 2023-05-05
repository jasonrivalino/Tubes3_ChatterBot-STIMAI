package checkdateformat

import (
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func IsValidDate(dateStr string) bool {
	// Memisahkan parameter string menjadi tiga bagian, yaitu tanggal, bulan, dan tahun,
	// dengan menggunakan delimiter garis miring (/) sebagai pemisah.
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return false
	}

	// Parse string menjadi integer untuk mendapatkan nilai tanggal, bulan, dan tahun.
	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return false
	}

	// Validasi tanggal
	if month < 1 || month > 12 || day < 1 {
		return false
	}

	// Validasi bulan dengan 31 hari
	if (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) && day > 31 {
		return false
	}

	// Validasi bulan dengan 30 hari
	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return false
	}

	// Validasi bulan Februari
	if month == 2 {
		if isLeapYear(year) && day > 29 {
			return false
		} else if !isLeapYear(year) && day > 28 {
			return false
		}
	}

	return true
}

func GetDayOfWeek(dateStr string) string {
	// Pisahkan parameter string menjadi tiga bagian, yaitu hari, bulan, dan tahun,
	// dengan menggunakan delimiter garis miring (/) sebagai pemisah.
	parts := strings.Split(dateStr, "/")

	// Parse string menjadi integer untuk mendapatkan nilai hari, bulan, dan tahun.
	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])

	if month < 3 {
		month += 12
		year -= 1
	}
	k := year % 100
	j := year / 100

	h := (day + 13*(month+1)/5 + k + k/4 + j/4 + 5*j) % 7

	// Mengubah nilai h agar merepresentasikan hari yang benar

	// Map index sisa bagi dengan hari dalam seminggu.
	daysOfWeek := map[int]string{
		0: "Sabtu",
		1: "Minggu",
		2: "Senin",
		3: "Selasa",
		4: "Rabu",
		5: "Kamis",
		6: "Jumat",
	}

	return daysOfWeek[h]
}
