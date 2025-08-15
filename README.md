## 📋 **Panduan Instalasi dan Menjalankan Aplikasi**

### **📥 Prerequisites**
Pastikan sudah terinstall:
- **Node.js** (v16 atau lebih baru)
- **Go** (v1.19 atau lebih baru) 
- **MySQL** atau **PostgreSQL**
- **Git**

---

### **🔧 Step 1: Clone Repository**
```bash
git clone https://github.com/username/repository-name.git
cd Capstoneeee
```

---

### **🗄️ Step 2: Setup Database**
1. **Buat database baru** di MySQL/PostgreSQL:
   ```sql
   CREATE DATABASE transmigrasi_db;
   ```

2. **Konfigurasi environment** di folder backend:
   ```bash
   cd backend
   cp .env.example .env  # Jika ada file example
   ```

3. **Edit file `.env`** dengan kredensial database Anda:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=transmigrasi_db
   JWT_SECRET=your_jwt_secret_key
   ```

---

### **🚀 Step 3: Setup Backend (Go)**
```bash
# Masuk ke folder backend
cd backend

# Download dependencies
go mod download

# Jalankan migrasi database dan seeder (otomatis)
go run main.go
```

**Backend akan berjalan di**: `http://localhost:8070`

---

### **🎨 Step 4: Setup Frontend (Vue.js)**
**Buka terminal baru**, kemudian:
```bash
# Masuk ke folder frontend
cd frontend

# Install dependencies
npm install

# Jalankan development server
npm run dev
```

**Frontend akan berjalan di**: `http://localhost:5173`

---

### **🔑 Step 5: Login Credentials**
Setelah backend berjalan, gunakan akun dummy berikut:

#### **👤 User/Masyarakat:**
- **Email**: `user@example.com`
- **Password**: `password123`

#### **👨‍💼 Admin:**
- **Email**: `admin@example.com`
- **Password**: `password123`

#### **⚡ Super Admin:**
- **Email**: `superadmin@example.com`
- **Password**: `password123`

---

### **📂 Step 6: Struktur Aplikasi**
```
Capstoneeee/
├── backend/          # Go + Fiber API Server (Port 8070)
├── frontend/         # Vue.js 3 + Vite (Port 5173)
└── README.md
```

---

### **🔧 Commands Berguna**

#### **Backend:**
```bash
# Jalankan server
go run main.go

# Build aplikasi
go build -o app main.go

# Jalankan dengan hot reload (jika punya air)
air
```

#### **Frontend:**
```bash
# Development mode
npm run dev

# Build untuk production
npm run build

# Preview build hasil
npm run preview
```

---

### **🌐 Cara Akses Aplikasi**

1. **Buka browser** dan akses: `http://localhost:5173`
2. **Daftar akun baru** atau **login** dengan kredensial di atas
3. **Explore fitur**:
   - User: Pendaftaran transmigrasi, track status
   - Admin: Verifikasi data, update status
   - Super Admin: Manajemen sistem

---

### **🐛 Troubleshooting**

#### **Database Connection Error:**
- Pastikan MySQL/PostgreSQL sudah running
- Cek kredensial di file `.env`
- Pastikan database sudah dibuat

#### **Port Already in Use:**
```bash
# Cek proses yang menggunakan port
lsof -i :8070  # Backend
lsof -i :5173  # Frontend

# Kill proses jika perlu
kill -9 <PID>
```

#### **Go Module Issues:**
```bash
go mod tidy
go mod download
```

#### **Node Dependencies Issues:**
```bash
rm -rf node_modules package-lock.json
npm install
```

---

### **📱 Fitur Utama**
- ✅ **Multi-role Authentication** (User, Admin, Super Admin)
- ✅ **Pendaftaran Transmigrasi Online**
- ✅ **Upload & Verifikasi Dokumen**
- ✅ **Real-time Status Tracking**
- ✅ **Dashboard Admin & Analytics**
- ✅ **Log Aktivitas Sistem**

---

### **🔗 API Endpoints**
- **Frontend**: `http://localhost:5173`
- **Backend API**: `http://localhost:8070/api`
- **API Documentation**: `http://localhost:8070/api/hello`

Aplikasi siap digunakan! 🎉
