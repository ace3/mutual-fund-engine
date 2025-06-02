## ❗️Baca baik baik petunjuk dibawah ini hingga selesai sebelum memulai test ini. ❗️

---

**NOBI Investment** adalah sebuah produk investasi yang memiliki mekanisme sebagai berikut:

1. Saat nasabah bergabung dalam suatu produk investasi, balance nasabah akan ditukarkan dengan Unit.   
   1. Unit didapat dari rumus: Balance Nasabah / NAB  
2. NAB adalah Nilai Aktiva Bersama, suatu angka yang menggambarkan performa produk investasi tersebut.  
   1. Angka NAB dihitung dengan rumus: Total Balance / Total Unit   
   2. Angka NAB dinamis dan tergantung kedua faktor diatas.  
   3. Angka NAB dibulatkan 4 angka dibelakang koma dengan metode round-down.  
   4. NAB saat belum ada nasabah nilainya \= 1  
3. Unit yang dimiliki Nasabah memiliki sifat:  
   1. Saat penyetoran balance, jumlah unit yang dimiliki nasabah bertambah.   
   2. Saat penarikan balance, jumlah unit yang dimiliki nasabah berkurang  
   3. Diluar penyetoran & penarikan dana, jumlah unit yang dimiliki nasabah selalu tetap.  
   4. Jumlah unit yang berkurang / bertambah dihitung dengan rumus: Balance / NAB  
   5. Unit dibulatkan 4 angka dibelakang koma dengan metode round-down.  
4. Balance adalah angka investasi dalam Rupiah.   
   1. Nasabah selalu melihat balance dalam Rupiah, tidak dalam unit & NAB   
   2. Balance yang disetor & ditarik juga selalu dalam Rupiah  
   3. Balance nasabah dihitung dengan rumus: Unit Nasabah x NAB  
   4. Balance dibulatkan 2 angka dibelakang koma dengan metode round-down.

### 

### Ilustrasi Mekanisme Investasi.

Tanggal 1

Arman adalah nasabah baru yang melakukan penyetoran balance sebesar Rp 5000\.   
Angka NAB pada tanggal 1 adalah 1.4000   
Unit yang didapat Arman adalah: 5000 / 1.4000 \= 3571.42857143 \= 3571.4285 unit 

Jadi, secara sistem Arman memiliki 3571.4285 unit yang bernilai Rp 5000 di tanggal 1

Tanggal 2

Budi adalah nasabah baru yang melakukan penyetoran balance sebesar Rp 15000  
Angka NAB pada tanggal 2 adalah 2.2250  
Unit yang didapat Budi adalah: 15000 / 2.2250 \= 6741.573033 \= 6741.5730 unit

Jadi, secara sistem Budi memiliki 6741.5730 unit yang bernilai Rp 15000 di tanggal 2

Tanggal 3

Angka NAB pada tanggal 3 adalah 2.4000

Arman melihat balance & mendapati balance dia Rp 8571.42   
Angka didapat dari unit yang dia miliki dikali NAB saat ini:   
2.4 x 3571.4285 \= 8571.4284 \=  8571.42

Arman melakukan penarikan balance sebesar Rp 6000  
Unit ditarik: 6000 / 2.4000 \= 2500  
Sisa Unit Arman \= 3571.4285 \- 2500 \= 1071.4285

### 

### Spesifikasi Teknis 

Bangun sebuah backend sederhana berbasis REST API yang mengimplementasikan sistem NOBI Investment dengan spesifikasi & ketentuan dibawah ini.

Pilihan bahasa pemrograman

- Golang atau Javascript atau Typescript atau framework berbasis Golang atau Javascript. Golang lebih preferable

Mengenai Database

- Implementasi menggunakan varian database MySQL   
- Desain database bebas   
- Setiap tabel harus memiliki kolom ID yang unique (auto increment) & menjadi primary key

Penilaian & submission

- Untuk kemudahan penilaian, buatlah Postman API List dari API yang dikembangkan. Export Postman API List, sertakan di submission.  
- Yang perlu disertakan di submission:  
  - MySQL Database dump   
  - Source code  
  - README yang berisikan dokumentasi & instruksi build   
- Metode submission:  
  - **Github** share github repository ke user GitHub **edibez** dan  **verzth**

Faktor Penilaian 

- Pemahaman requirements  
  - Feature completion  
  - Logika dan Validasi  
  - Code cleanliness  
  - Creativity  
  - Dokumentasi   
  - Point extra untuk extra effort (silahkan dijelaskan di README)

REST API Endpoint yang perlu dibangun: 

| \#1 POST  /api/v1/user/add |                                                                  |
| :------------------------- | :--------------------------------------------------------------- |
| Input                      | name, username                                                   |
| Output                     | id user (running number) yang baru saja ditambahkan              |
| Logic                      | Menambahkan user. Username sifatnya unique jadi tidak boleh sama |

| \#2 POST /api/v1/ib/updateTotalBalance |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| :------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Input                                  | current\_balance                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| Output                                 | nab\_amount                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| Logic                                  | Update balance dengan angka yang diberkan di input. Karena balance berubah, akibatnya NAB baru akan terbentuk.  Generate NAB dari input parameter current\_balance (float, 5 digit dibelakang koma). Jika total seluruh unit user saat ini adalah 0, maka NAB akan selalu 1\. Jika tidak akan mengikuti rumus yang sudah dijelaskan diatas.  Simpan angka NAB di database, setiap kali endpoint ini dipanggil maka angka ini akan yang akan menjadi NAB  yang akan dipakai dalam perhitungan penyetoran balance / penarikan balance nasabah Tampilkan NAB hasil perhitungan di response API |

| \#3 GET /api/v1/ib/listNAB |                                                                                                                                                  |
| :------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------- |
| Input                      | tidak ada                                                                                                                                        |
| Output                     | List NAB dan perubahanya eg :  {{nab:1.22,date:”2021-02-01 10:00:00}, {nab:2.33,date:”2021-02-01 09:00:00}, {nab:1.4,date:”2021-02-01 08:00:00}} |
| Logic                      | Jika endpoint ini dipanggil akan mereturn list perubahan NAB diurutkan dari perubahan terakhir.                                                  |

| \#4 POST /api/v1/ib/topup |                                                                                                                                                                   |
| :------------------------ | :---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Input                     | user\_id, amount\_rupiah                                                                                                                                          |
| Output                    | {nilai\_unit\_hasil\_topup, nilai\_unit\_total, saldo\_rupiah\_total}                                                                                             |
| Logic                     | Jika endpoint ini dipanggil akan mencatat di database bahwa user\_id tersebut melakukan penyetoran balance dan mendapatkan unit sesuai dengan ketentuan rumus NAB |

| \#5 POST /api/v1/ib/withdraw |                                                                                                                                                                                                                                         |
| :--------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Input                        | user\_id, amount\_rupiah                                                                                                                                                                                                                |
| Output                       | {nilai\_unit\_setelah\_withdraw, nilai\_unit\_total, saldo\_rupiah\_total}                                                                                                                                                              |
| Logic                        | Jika endpoint ini dipanggil akan mencatat di database bahwa user\_id tersebut melakukan penarikan balance  dan terjadi pengurangan  unit sesuai dengan ketentuan rumus NAB . User tidak boleh withdraw lebih dari asset yang dia miliki |

| \#6 GET /api/v1/ib/member |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| :------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Input                     | userid (optional), page (optional) , limit (optional)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| Output                    | list seluruh nasabah yang saat ini investasi dengan menampilkan informasi userid, total\_unit\_per\_uid, total\_amountrupiah\_per\_uid, Current NAB                                                                                                                                                                                                                                                                                                                                                                                               |
| Logic                     | Menampilkan list user yang saat ini sedang mengikuti Nobi Investment beserta total nilai asset dia baik secara rupiah maupun secara unit, limit adalah total baris yang ingin ditampilkan, page adalah pagination dari limit. Jika param page / limit tidak diikut sertakan maka nilai default page adalah \= page 0 (page pertama) dan limit \= 20 . Ditampilkan sort by userid 1 user hanya akan muncul 1x di hasil output, jadi misalkan user pernah top up 2x dan exit 1x yang ditampilkan hanya posisi balance terakhir bukannya jadi 3 row. |

