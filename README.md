# 🎞️ Video Frame Extractor Microservice

A high-performance Go microservice that listens to a Kafka topic for video processing jobs, uses **FFMPEG** to extract frames from videos, and leverages **Go's concurrency model** (goroutines) for fast, parallel processing.

---

## 📌 Features

- 🎥 Extracts video frames using FFMPEG
- 📩 Listens to a Kafka topic for video processing jobs
- ⚡ Utilizes goroutines for concurrent video processing
- 📁 Outputs image frames to a specified directory
- 🐳 Docker-compatible for containerized deployments
- ✅ Lightweight and efficient architecture

---

## 🔧 Tech Stack

- **Golang** – for backend logic and concurrency
- **Apache Kafka** – for job/message queue
- **FFMPEG** – for video-to-frames conversion
- **Docker** – optional containerization

---

## 📦 Installation

### 1. Clone the Repository
```bash
git clone https://github.com/zaac04/ffmpeg-frame-extarcter.git
cd ffmpeg-frame-extarcter



