# 🎬 FilmFindr – Temukan, Simpan, dan Review Film Favoritmu

**FilmFindr** adalah platform film interaktif yang dirancang untuk para pencinta film sejati. Temukan film menarik, berikan ulasan pribadi, dan simpan film ke dalam daftar tontonmu hanya dengan satu akun!

---

## 🚀 Fitur Unggulan

✨ **Autentikasi Aman**  
Daftar & login dengan cepat dan aman untuk menikmati fitur penuh.

📝 **Review Film**  
Tulis ulasan dan beri rating untuk setiap film yang kamu tonton.

🎯 **Watch List Pribadi**  
Simpan film yang ingin kamu tonton nanti dalam daftar pribadi.

🔎 **Eksplorasi Film**  
Temukan berbagai film menarik dari berbagai genre dan kategori.

---

## 🛠️ Teknologi yang Digunakan

| Layer       | Teknologi                  |
|-------------|----------------------------|
| Frontend    | ⚛️ React.js                |
| Backend     | 🧠 Golang (Gin + GORM)     |
| Database    | 🐘 PostgreSQL              |

Arsitektur modern yang memisahkan frontend dan backend demi skalabilitas dan fleksibilitas pengembangan.


---

## 🧪 Cara Menjalankan di Lokal

### 1. Clone Proyek
```bash
git clone https://github.com/Anamm15/film-findr.git
cd film-findr
```

### 2. Atur Backend
# a. Masuk ke direktori backend 
```bash
cd backend
```

# b. Buat file .env dan isi dengan
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=film-findr
PORT=8080
```

# c. Jalankan backend
```bash
Jalankan backend
```

### 2. Atur Frontend
# a. Masuk ke direktori frontend 
```bash
cd frontend
```

# b. Install dependency
```bash
npm install
```

# c. Jalankan React dev server
```bash
npm run dev
```

