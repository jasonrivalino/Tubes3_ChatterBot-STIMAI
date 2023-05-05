# TUGAS BESAR 3 STRATEGI ALGORITMA - IF2211
> Tugas Besar 3: Penerapan String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana

## Anggota Kelompok
<table>
    <tr>
        <td colspan="3", align = "center"><center>Nama Kelompok: ChatterBot-STIM(AI)</center></td>
    </tr>
    <tr>
        <td>No.</td>
        <td>Nama</td>
        <td>NIM</td>
    </tr>
    <tr>
        <td>1.</td>
        <td>Jason Rivalino</td>
        <td>13521008</td>
    </tr>
    <tr>
        <td>2.</td>
        <td>Muhhamad Syauqi Jannatan</td>
        <td>13521014</td>
    </tr>
    <tr>
        <td>3.</td>
        <td>Muhammad Zulfiansyah Bayu Pratama</td>
        <td>13521028</td>
    </tr>
</table>

## Table of Contents
* [Deskripsi Singkat](#deskripsi-singkat)
* [Struktur File](#struktur-file)
* [Requirements](#requirements)
* [Cara Menjalankan Program](#cara-menjalankan-program)
* [Cara Mengoperasikan Program](#cara-mengoperasikan-program)
* [Tampilan Interface Program](#tampilan-interface-program)
* [Acknowledgements](#acknowledgements)
* [Foto Bersama](#foto-bersama)

## Deskripsi Singkat
Repository ini berisi aplikasi ChatGPT sederhana berbasis web. Aplikasi ChatGPT sederhana ini dibangun dengan mengimplementasikan algoritma string matching (Knuth-Morris-Pratt dan Boyer-Moore) dan juga regular expression. Aplikasi ini dapat menjawab pertanyaan berdasarkan masukan dari pengguna. Jenis pertanyaan yang dapat dijawab antara lain:
1. Pertanyaan yang terdapat didalam database (menjawab menggunakan algoritma KMP atau BM)
2. Kalkulator sederhana
3. Menanyakan hari berdasarkan tanggal
4. Menambahkan pertanyaan dan jawaban ke dalam database
5. Menghapus pertanyaan dari database

Selain itu, history dari pertanyaan yang dimasukkan akan disimpan ke dalam sebuah section pertanyaan untuk sebuah sesi. Pengguna juga dapat menambahkan section baru dengan mengklik add new history. 

Untuk pembuatan program ini, pembuatan frontend dilakukan dengan menggunakan bahasa pemrograman HTML, CSS, dan juga Javascript dengan menggunakan Framework React. Untuk Backend dari program dibuat dengan menggunakan bahasa pemrograman Golang. Sedangkan untuk penyimpanan database menggunakan MySQL.

## Struktur File
```bash
📦Tubes3_ChatterBot-STIMAI
 ┣ 📂.vscode
 ┃ ┗ 📜settings.json
 ┣ 📂doc
 ┃ ┃ 📜ChatterBot-STIM(AI)-22.pdf
 ┃ ┗ 📜Spesifikasi Tugas Besar 3 IF2211 Strategi Algoritma 2023.pdf
 ┣ 📂src
 ┃ ┣ 📂backend
 ┃ ┃ ┣ 📂component
 ┃ ┃ ┃ ┣ 📂bm
 ┃ ┃ ┃ ┃ ┗ 📜bm.go
 ┃ ┃ ┃ ┣ 📂checkdateformat
 ┃ ┃ ┃ ┃ ┗ 📜checkdateformat.go
 ┃ ┃ ┃ ┣ 📂kmp
 ┃ ┃ ┃ ┃ ┗ 📜kmp.go
 ┃ ┃ ┃ ┣ 📂lcs
 ┃ ┃ ┃ ┃ ┗ 📜lcs.go
 ┃ ┃ ┃ ┣ 📂mathoperation
 ┃ ┃ ┃ ┃ ┗ 📜mathoperation.go
 ┃ ┃ ┃ ┗ 📂searchanswer
 ┃ ┃ ┃ ┃ ┗ 📜searchanswer.go
 ┃ ┃ ┣ 📂controllers
 ┃ ┃ ┃ ┣ 📂listquestioncontroller
 ┃ ┃ ┃ ┃ ┗ 📜listquestioncontroller.go
 ┃ ┃ ┃ ┗ 📂riwayatcontroller
 ┃ ┃ ┃ ┃ ┗ 📜riwayatcontroller.go
 ┃ ┃ ┣ 📂models
 ┃ ┃ ┃ ┣ 📂stack
 ┃ ┃ ┃ ┃ ┗ 📜stack.go
 ┃ ┃ ┃ ┣ 📂stackfloat
 ┃ ┃ ┃ ┃ ┗ 📜stackfloat.go
 ┃ ┃ ┃ ┣ 📜listquestion.go
 ┃ ┃ ┃ ┣ 📜riwayat.go
 ┃ ┃ ┃ ┗ 📜setup.go
 ┃ ┃ ┣ 📂output
 ┃ ┃ ┃ ┗ 📜regex.exe
 ┃ ┃ ┣ 📜go.mod
 ┃ ┃ ┣ 📜go.sum
 ┃ ┃ ┗ 📜main.go
 ┃ ┣ 📂data
 ┃ ┃ ┗ 📜data.sql
 ┃ ┗ 📂frontend
 ┃ ┃ ┣ 📂frontendsrc
 ┃ ┃ ┃ ┣ 📂components
 ┃ ┃ ┃ ┃ ┗ 📂assets
 ┃ ┃ ┃ ┗ 📂styles
 ┃ ┃ ┣ 📂node_modules
 ┃ ┃ ┣ 📂public
 ┃ ┃ ┃ ┣ 📜favicon.ico
 ┃ ┃ ┃ ┣ 📜index.html
 ┃ ┃ ┃ ┣ 📜manifest.json
 ┃ ┃ ┃ ┣ 📜robots.txt
 ┃ ┃ ┃ ┗ 📜webLogo.png
 ┃ ┃ ┣ 📂src
 ┃ ┃ ┃ ┣ 📂components
 ┃ ┃ ┃ ┃ ┣ 📂assets
 ┃ ┃ ┃ ┃ ┃ ┣ 📜blank.jpg
 ┃ ┃ ┃ ┃ ┃ ┗ 📜robot.jpeg
 ┃ ┃ ┃ ┃ ┣ 📜ChatHistory.js
 ┃ ┃ ┃ ┃ ┣ 📜ChatInput.js
 ┃ ┃ ┃ ┃ ┣ 📜ChatItem.js
 ┃ ┃ ┃ ┃ ┣ 📜ChatList.js
 ┃ ┃ ┃ ┃ ┗ 📜RadioButton.js
 ┃ ┃ ┃ ┣ 📂styles
 ┃ ┃ ┃ ┃ ┣ 📜ChatHistory.css
 ┃ ┃ ┃ ┃ ┣ 📜ChatInput.css
 ┃ ┃ ┃ ┃ ┣ 📜ChatItem.css
 ┃ ┃ ┃ ┃ ┣ 📜ChatList.css
 ┃ ┃ ┃ ┃ ┣ 📜LandingPage.css
 ┃ ┃ ┃ ┃ ┗ 📜RadioButton.css
 ┃ ┃ ┃ ┣ 📜App.css
 ┃ ┃ ┃ ┣ 📜App.js
 ┃ ┃ ┃ ┣ 📜index.css
 ┃ ┃ ┃ ┗ 📜index.js
 ┃ ┃ ┣ 📜.gitignore
 ┃ ┃ ┣ 📜package-lock.json
 ┃ ┃ ┗ 📜package.json
 ┗ 📜README.md
 ```
 
## Requirements
1. Visual Studio Code
2. Bahasa Pemrograman Golang
3. Node.js untuk running frontend

## Cara Menjalankan Program
Langkah-langkah proses setup program adalah sebagai berikut:
1. Clone repository ini
2. Menjalankan frontend yang ada dengan mengetikkan `npm start` (dijalankan pada direction folder src yang terdapat dalam frontend)

## Cara Mengoperasikan Program
1. Pengguna dapat memasukkan pesan kedalam aplikasi
2. Pengguna dapat memilih Algoritma untuk menjawab pertanyaan antara KMP (Knuth-Morris-Pratt) dan juga BM (Boyer-Moore)
3. Aplikasi akan memberikan respon jawaban dari pertanyaan
4. Setiap pertanyaan yang ditanyakan akan disimpan ke dalam section tertentu
5. Pengguna juga dapat menambahkan section baru untuk memulai section baru dan pertanyaan lama akan tersimpan dalam section sebelumnya

## Tampilan Interface Program
![Screenshot (181)](https://user-images.githubusercontent.com/91790457/236441708-c1997cae-a9da-48e6-9d13-e9e8af6d38b8.png)

## Acknowledgements
- Tuhan Yang Maha Esa
- Dosen Mata Kuliah yaitu Pak Rinaldi (K1), Bu Ulfa (K2), dan Pak Rila (K3)
- Kakak-Kakak Asisten Mata Kuliah Strategi Algoritma IF2211

## Foto Bersama
