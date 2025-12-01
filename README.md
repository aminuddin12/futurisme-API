# Futurisme API

**Futurisme API** adalah ekosistem backend yang dibangun menggunakan bahasa pemrograman **Go (Golang)**. Aplikasi ini dirancang dengan fokus pada kecepatan, keamanan, dan skalabilitas tinggi, menggunakan arsitektur modular yang rapi (Clean Architecture) untuk memudahkan pemeliharaan dan pengembangan fitur baru.

---

## 📖 Deskripsi Singkat

Backend ini berfungsi sebagai pusat logika bisnis dan pengolahan data untuk ekosistem aplikasi Futurisme. Dibangun di atas framework **Go-Fiber** (yang dikenal dengan performanya yang ekstrem), API ini menyediakan layanan RESTful yang aman dengan autentikasi berlapis (*Multilayer Authentication*) dan manajemen database yang fleksibel (*Hybrid Migration*).

---

## 🛠 Terminologi & Tech Stack

Teknologi dan pola desain utama yang digunakan dalam aplikasi ini:

* **Language:** Go (Golang) v1.23+
* **Framework:** [Go-Fiber](https://gofiber.io/) (Express-style, High Performance)
* **Database:** PostgreSQL
* **ORM:** GORM (Development) & SQL Migration (Production)
* **CLI Library:** Cobra (Untuk manajemen perintah terminal)
* **Configuration:** Viper (Manajemen Environment Variable)
* **Authentication:**
    * **Layer 1:** Application Key (Header `X-App-Key`)
    * **Layer 2:** JWT (JSON Web Token) untuk User Session
* **Architecture:** Domain-Driven / Clean Architecture (Modular per fitur)

---

## 📂 Struktur Aplikasi

Struktur direktori dirancang sedinamis mungkin agar mudah dipahami oleh pengembang yang terbiasa dengan framework seperti Laravel, namun tetap mengikuti kaidah *idiomatic* Go.

futurisme-api/
├── cmd/
│   └── commands/        # Definisi perintah CLI (start, seed, dll) - Mirip Artisan Console
├── config/              # Konfigurasi aplikasi & Environment (Viper)
├── internal/
│   ├── middleware/      # Middleware global (Auth, Logger, CORS)
│   ├── modules/         # MODUL FITUR (Domain Based)
│   │   ├── auth/        # Fitur Login & Register
│   │   └── user/        # Fitur Manajemen User
│   │       ├── delivery/    # HTTP Handlers (Controller)
│   │       ├── usecase/     # Business Logic (Service)
│   │       ├── repository/  # Data Access (Query/Model)
│   │       └── entity/      # Struktur Data Database
│   └── server/          # Setup utama Fiber App & Routing
├── pkg/                 # Library bantuan (Shared Code)
│   ├── database/        # Koneksi Database & Logic Hybrid Migration
│   ├── utils/           # Helper functions (Response, JWT, Hash, dll)
├── .env                 # File konfigurasi Environment (JANGAN DI-COMMIT KE GIT)
├── go.mod               # Dependency Manager
├── main.go              # Entry Point Aplikasi
└── Makefile             # Shortcut perintah terminal

---

## 🚀 Cara Menjalankan Aplikasi

### 1. Prasyarat
Pastikan Anda telah menginstal:
* Go (versi terbaru)
* PostgreSQL
* Git

### 2. Instalasi
Clone repositori dan unduh dependency:

```bash
git clone [https://github.com/username/futurisme-api.git](https://github.com/username/futurisme-api.git)
cd futurisme-api
go mod tidy