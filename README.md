# ddd-notification

Use Case: Sistem Notifikasi (Notification Service)

Deskripsi:
Buatlah sebuah sistem pengelolaan notifikasi yang dapat mengirimkan pemberitahuan ke pengguna berdasarkan preferensi mereka.
Sistem ini akan menggunakan arsitektur Domain Driven Design (DDD) dan framework Go Fiber untuk pengelolaan HTTP request.

Fitur yang harus diimplementasikan:

1. Model Domain:
Buatlah model domain untuk Notification yang menyimpan informasi tentang pengguna dan jenis notifikasi yang ingin dikirimkan (misalnya: Email, SMS).
Buatlah agregat dan layanan domain untuk memproses dan mengirimkan notifikasi.

2. API Endpoints:
Buat API untuk membuat notifikasi baru.
API untuk mendapatkan daftar notifikasi yang belum dikirimkan.
API untuk menandai notifikasi sebagai terkirim.

3. Validasi dan Pengolahan Notifikasi:
Implementasikan validasi pada input, seperti memastikan format email yang benar.
Implementasikan logika untuk memproses notifikasi (misalnya, apakah notifikasi yang dikirimkan adalah Email atau SMS).

4. Integrasi Framework Go Fiber:
Gunakan Go Fiber untuk membuat routing dan menyiapkan server HTTP.