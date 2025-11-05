# 🛡️ Sentinel

**Sentinel** adalah layanan otentikasi dan otorisasi terpusat berbasis **JWT** untuk arsitektur **microservices**.  
Dibangun menggunakan **Golang (Fiber)** untuk backend dan **React** untuk UI administrasi.

---

## 🚀 Fitur Utama

- **Autentikasi berbasis JWT**
  - Login dan logout (dengan revokasi semua token aktif).
  - Validasi token untuk setiap request API.
- **Manajemen Pengguna**
  - Tambah, ubah, dan hapus pengguna.
  - Kelola status aktif/nonaktif pengguna.
- **Manajemen Role dan Permission**
  - Role-based access control (RBAC) untuk menentukan akses endpoint.
  - Granular permission per action (read/write/delete).
- **Integrasi Gateway**
  - Kompatibel dengan **Traefik** melalui mekanisme `forwardAuth`.
- **UI Admin**
  - Dibangun dengan React untuk kemudahan manajemen user & role.

---

## 🧩 Arsitektur Singkat
[Client] → [Traefik Gateway] → [Sentinel] → [JWT Validation] → [Microservice]

Sentinel bertindak sebagai **Auth Service** terpusat yang menerbitkan dan memvalidasi token.  
Microservice lain cukup memverifikasi header JWT dan, bila perlu, memanggil endpoint introspeksi Sentinel.

---

## ⚙️ Teknologi

- **Backend:** Go 1.23+, Fiber, JWT, PostgreSQL/Redis
- **Frontend:** React, TypeScript, TailwindCSS
- **Deployment:** Docker + Kubernetes (opsional Traefik integration)

---

## 🧪 Rencana Modul

| Tahap | Modul | Deskripsi |
|-------|--------|-----------|
| 1 | Auth Core | Login, Logout, Token issuance, Revocation |
| 2 | User Management | CRUD user dan validasi akun |
| 3 | Role Management | RBAC, mapping role → permission |
| 4 | Admin UI | Halaman login dan dashboard |
| 5 | Gateway Integration | Endpoint validasi untuk Traefik |

---

## 🔐 Integrasi dengan Microservice

Setiap service cukup menambahkan middleware JWT validation.  
Jika diperlukan introspeksi token (misalnya cek status revoke), panggil endpoint:
