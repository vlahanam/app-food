# App Food - Docker & Nginx Environment

Dự án **App Food** sử dụng kiến trúc container hóa với:
- **Golang 1.26** (Hỗ trợ live-reload bằng Air & quản lý migration qua `golang-migrate`)
- **PostgreSQL 18** (Alpine image kèm healthcheck)
- **Nginx 1.26** (Reverse Proxy hiệu năng cao, tối ưu Gzip, buffers và bảo mật)

---

## 🚀 Hướng dẫn khởi chạy nhanh

### 1. Chuẩn bị file môi trường
Tạo file `.env` từ file mẫu `.env.example`:
```bash
cp .env.example .env
```

### 2. Khởi chạy bằng Makefile hoặc Docker Compose
```bash
# Khởi chạy môi trường Dev
make up
# hoặc: docker compose up -d
```

### 3. Kiểm tra trạng thái và logs
```bash
make ps
make logs
```

---

## 🌐 Danh sách Cổng (Ports) & Điểm truy cập

| Dịch vụ | Địa chỉ truy cập | Mô tả |
| :--- | :--- | :--- |
| **Nginx (Reverse Proxy)** | `http://localhost:8088` | Điểm vào chính của ứng dụng |
| **Nginx Healthcheck** | `http://localhost:8088/healthz` | Kiểm tra trạng thái hoạt động của Nginx |
| **Backend API (qua Nginx)** | `http://localhost:8088/api` | Reverse proxy chuyển tiếp vào backend server |
| **Backend Trực tiếp** | `http://localhost:8090` | Cổng backend (Go Server container) |
| **PostgreSQL 18** | `localhost:5432` | Cổng truy cập database từ máy host |

---

## 🛠️ Các lệnh tiện ích (Makefile)

| Lệnh | Chức năng |
| :--- | :--- |
| `make up` | Khởi chạy môi trường Dev ở chế độ background |
| `make down` | Dừng và hạ toàn bộ container môi trường Dev |
| `make restart` | Khởi động lại các container Dev |
| `make build` | Build lại toàn bộ container |
| `make logs` | Xem stream log của tất cả services |
| `make ps` | Xem danh sách và trạng thái containers |
| `make server-sh` | Mở terminal sh trong container Go Server |
| `make nginx-sh` | Mở terminal sh trong container Nginx |
| `make nginx-reload` | Reload Nginx configuration không cần restart container |
| `make psql-cli` | Mở CLI tương tác trực tiếp với PostgreSQL 18 (`food_db`) |
| `make migration name=<tên>` | Tạo file migration mới trong thư mục `server/migrations/` |
| `make migrate-up` | Chạy toàn bộ migration mới nhất vào database |
| `make migrate-down` | Rollback 1 bước migration gần nhất |
| `make clean` | Dừng container và xóa toàn bộ volumes, dữ liệu |

---

## 📁 Cấu trúc thư mục

```
app-food/
├── docker-compose.yml          # Compose môi trường Dev / Local
├── docker-compose.staging.yml  # Compose môi trường Staging
├── docker-compose.product.yml  # Compose môi trường Production
├── Makefile                    # Phím tắt điều khiển và thao tác dự án
├── .env.example                # Cấu hình biến môi trường mẫu
├── .env                        # File biến môi trường thực tế (đã gitignore)
├── nginx/
│   ├── Dockerfile              # Dockerfile Nginx môi trường Dev
│   ├── Dockerfile.staging      # Dockerfile Nginx Staging
│   ├── Dockerfile.product      # Dockerfile Nginx Production
│   ├── nginx.conf              # Cấu hình Nginx chính (Gzip, Logging, Buffers)
│   └── conf.d/
│       ├── default.conf        # Virtual host Dev (Proxy /api và /)
│       ├── staging.conf        # Virtual host Staging (Headers bảo mật)
│       └── product.conf        # Virtual host Production
└── server/
    ├── cmd/                    # Entrypoint ứng dụng Golang
    ├── internal/               # Mã nguồn nội bộ (controllers, services, ...)
    ├── migrations/             # Thư mục lưu migration SQL
    ├── .air.toml               # File cấu hình Air hot-reloading
    ├── Dockerfile              # Dockerfile Go 1.26 Dev (Air + migrate)
    ├── Dockerfile.staging      # Multi-stage Staging
    ├── Dockerfile.product      # Multi-stage Hardened Production
    └── go.mod                  # Go module definition (Golang 1.26)
```
