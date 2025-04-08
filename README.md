# 🛒 E-Commerce Backend (Go)

A lightweight, containerized e-commerce backend written in Go, powered by MySQL, Redis, and Kafka.

---

## 🚀 Getting Started

### 1. Clone the repo & enter project
```
cd e-commerce
```

### 2. Start infrastructure services (MySQL, Redis, Kafka)
```
docker compose up -d
```

### 3. Install dependencies (make & goose)
```
sudo apt update
sudo apt install make -y
curl -fsSL https://github.com/block/goose/releases/download/stable/download_cli.sh | bash
```
### 4. Configure your environment
```
mv config/production.yaml.example config/production.yaml
```

### 5. Run database migration
```
make upse
```

### 6. Build backend Docker image
```
docker build -t ecommerce:latest .
```

### 7. Run the backend container
```
docker run --network ecommerce_network -p 8082:8082 -it ecommerce:latest
```

## 📝 Tech Stack
- **Language**: Go 1.20+
- **Database**: MySQL 5.7
- **Cache**: Redis
- **Queue**: Apache Kafka (KRaft mode)
- **Migration**: goose
- **Containerization**: Docker + Docker Compose

## 🤝 Contributing
Feel free to fork and submit PRs if you'd like to improve or add features!

## 🧪 License
MIT © 2025