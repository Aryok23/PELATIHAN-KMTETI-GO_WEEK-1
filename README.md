# PELATIHAN-KMTETI-GO_WEEK-1

Berikut adalah penjelasan alur kode dari program konversi suhu yang telah dibuat dalam bahasa pemrograman Go:

1. Deklarasi Package dan Import
Program ini dimulai dengan deklarasi package "main", yang menandakan bahwa program ini adalah sebuah eksekusi mandiri. Package "fmt" di-import untuk menyediakan fungsi input/output seperti "Println", "Print", dan "Scan".

2. Fungsi "main()"
Fungsi "main()" adalah titik awal eksekusi program. Semua logika dan alur eksekusi berada di dalam fungsi ini.

3. Menampilkan Menu Konversi
Bagian ini mencetak judul dan menu konversi suhu ke layar. Pengguna akan diminta memilih skala suhu yang diinginkan dengan memasukkan angka yang sesuai (1, 2, atau 3).

4. Input Pilihan Skala Suhu
Program menunggu input dari pengguna untuk memilih skala suhu yang akan dikonversi. Input disimpan dalam variabel "temperatureScale" bertipe "int".

5. Menggunakan "switch" untuk Mengevaluasi Pilihan
Penggunaan "switch" di sini untuk memproses tiga kemungkinan pilihan pengguna. Setiap "case" mewakili konversi suhu yang berbeda:
Case 1: Konversi dari Celsius ke Fahrenheit
Case 2: Konversi dari Celsius ke Kelvin
Case 3: Konversi dari Celsius ke Reamur
Default: Pesan ditampilkan jika input tidak valid.

6. Konversi Celsius ke Fahrenheit (Case 1)
Jika pengguna memilih "1", program meminta input suhu dalam Celsius. Program kemudian mengonversi suhu Celsius ke Fahrenheit dengan rumus: Fahrenheit = (Celsius × 9/5) + 32 Hasilnya dicetak dengan dua angka di belakang koma.

7. Konversi Celsius ke Kelvin (Case 2)
Jika pengguna memilih "2", program mengonversi suhu Celsius ke Kelvin dengan rumus: Kelvin = Celsius + 273 Hasil konversi dicetak dengan dua angka di belakang koma.

8. Konversi Celsius ke Reamur (Case 3)
Jika pengguna memilih "3", program akan mengonversi suhu Celsius ke Reamur dengan rumus: Reamur=Celsius × 4/5 Hasil konversi dicetak ke layar.

9. Pesan Kesalahan untuk Input yang Tidak Valid
Jika pengguna memasukkan angka selain "1", "2", atau "3", program menampilkan pesan bahwa input skala suhu tidak valid.

Alur Kerja Singkat:

1. Program menampilkan menu konversi suhu dan meminta input pilihan pengguna.
2. Berdasarkan input, program akan meminta pengguna memasukkan suhu dalam Celsius.
3. Program menghitung konversi suhu ke skala yang dipilih (Fahrenheit, Kelvin, atau Reamur).
4. Hasil konversi ditampilkan, atau jika input salah, pesan kesalahan akan diberikan.



Catatan : Penjelasan alur kode dibantu menggunakan tool AI ChatGPT
