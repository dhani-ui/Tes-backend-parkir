# 🏢 Sistem Manajemen Parkir Modern & CRM Otomatis

Sistem Manajemen Parkir berbasis web yang dirancang untuk efisiensi operasional tinggi. Dilengkapi dengan denah interaktif multi-lantai dan fitur **CRM WhatsApp Otomatis** untuk meningkatkan pengalaman pelanggan.

![React](https://img.shields.io/badge/Frontend-ReactJS-61DAFB?style=for-the-badge&logo=react)
![Golang](https://img.shields.io/badge/Backend-Golang-00ADD8?style=for-the-badge&logo=go)
![Tailwind](https://img.shields.io/badge/Styling-TailwindCSS-38B2AC?style=for-the-badge&logo=tailwind-css)
![Konva](https://img.shields.io/badge/Graphics-KonvaJS-333333?style=for-the-badge)

---

## 🌟 Fitur Unggulan

### 1. 🏗️ Arsitektur Multi-Lantai (Enterprise Scale)
- Mendukung **5 Lantai** dengan total **250 Spot Parkir**.
- Denah interaktif menggunakan **Konva.js** untuk performa rendering yang mulus.
- Sistem navigasi tab lantai yang responsif.

### 2. ⏱️ Real-Time Monitoring
- Timer parkir presisi dengan format **Jam : Menit : Detik**.
- Dashboard ringkasan: Total Kapasitas, Spot Tersedia, dan Spot Terisi.
- Indikator warna dinamis (Hijau: Tersedia, Merah: Terisi, Pink: Overtime).

### 3. 🤖 CRM WhatsApp Automation (Integrasi Fonnte)
Sistem secara proaktif berkomunikasi dengan pelanggan melalui WhatsApp:
- **Digital Ticket:** Mengirim rincian spot dan plat nomor saat mobil masuk.
- **Auto-Reminder:** Mengirim peringatan otomatis tepat **5 menit** sebelum durasi parkir habis.
- **Manual Ping:** Fitur tombol lonceng (🔔) bagi petugas untuk memanggil pelanggan dalam kondisi darurat.
- **Digital Receipt:** Mengirim struk tagihan otomatis saat sesi parkir berakhir (Checkout).

### 4. 💰 Billing System
- Perhitungan tarif otomatis: **Rp 3.000 / Jam**.
- Pembulatan ke atas (Fixed hourly billing).

---

## 🛠️ Teknologi yang Digunakan

- **Frontend:** React.js, TypeScript, Vite.
- **Graphics:** React-Konva (Canvas API).
- **Backend (CRM Proxy):** Golang (Native Net/HTTP).
- **Styling:** Tailwind CSS.
- **Storage:** LocalStorage (Persistensi data antar sesi).
- **API Gateway:** Fonnte WhatsApp API.

---

## 🚀 Cara Menjalankan

### Persyaratan
- Node.js & NPM
- Go (Golang)
- Akun Fonnte (untuk fitur WA)

### 1. Jalankan Backend (Server akan berjalan di http://localhost:8080CRM Service)
```bash
cd backend-parkir
# Edit main.go dan masukkan Token Fonnte Anda
go run main.go
```
Server akan berjalan di http://localhost:8080

2. Jalankan Frontend
cd sistem-parkir
npm install
npm run dev

Aplikasi dapat diakses di http://localhost:5173

​📸 Antarmuka Pengguna
​Denah: Grid interaktif 5x10 per lantai.
​Modal Form: Input data kendaraan yang bersih dan fokus.
​Timer Card: Kartu kontrol aktif dengan fitur peringatan visual.
