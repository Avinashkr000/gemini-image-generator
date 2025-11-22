# 🎨 Gemini Image Generator

<div align="center">

![Gemini Image Generator](https://img.shields.io/badge/Gemini-AI%20Powered-purple?style=for-the-badge&logo=google)
![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react)
![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=for-the-badge&logo=mysql)

A full-stack AI-powered image generation application using Google Gemini API, built with Go backend and React frontend.

[Features](#-features) • [Tech Stack](#-tech-stack) • [Getting Started](#-getting-started) • [API Endpoints](#-api-endpoints) • [Screenshots](#-screenshots)

</div>

---

## ✨ Features

- 🎨 **AI Image Generation** - Generate stunning images using Google Gemini 2.0 Flash
- 💾 **Image History** - Store and manage all generated images
- 📊 **Real-time Stats** - Track generation status and statistics
- 🎯 **Smart Prompts** - Example prompts for inspiration
- ⬇️ **Download Images** - Save generated images locally
- 🗑️ **Image Management** - Delete unwanted images
- 📱 **Responsive Design** - Beautiful UI that works on all devices
- 🚀 **Fast Performance** - Go backend for lightning-fast responses

## 🛠 Tech Stack

### Backend
- **Go 1.21** - High-performance backend
- **Gin Framework** - Fast HTTP web framework
- **GORM** - The fantastic ORM library for Golang
- **MySQL 8.0** - Reliable relational database
- **Gemini API** - Google's latest AI model for image generation

### Frontend
- **React 18** - Modern UI library
- **Vite** - Lightning-fast build tool
- **Tailwind CSS** - Utility-first CSS framework
- **Axios** - HTTP client for API calls
- **Lucide React** - Beautiful icon library

### DevOps
- **Docker** - Containerization
- **Docker Compose** - Multi-container orchestration

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- MySQL 8.0 or higher
- Gemini API Key ([Get it here](https://ai.google.dev/))

### Installation

#### 1. Clone the repository

```bash
git clone https://github.com/Avinashkr000/gemini-image-generator.git
cd gemini-image-generator
```

#### 2. Setup MySQL Database

```bash
# Login to MySQL
mysql -u root -p

# Create database
CREATE DATABASE gemini_image_generator;

# Create user (optional)
CREATE USER 'gemini'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON gemini_image_generator.* TO 'gemini'@'localhost';
FLUSH PRIVILEGES;
```

#### 3. Backend Setup

```bash
cd backend

# Copy environment variables
cp .env.example .env

# Edit .env and configure:
# - GEMINI_API_KEY=your_actual_api_key_here
# - DB_PASSWORD=your_mysql_password
# - DB_USER=root (or your mysql user)

# Install dependencies
go mod download

# Run the backend
go run main.go
```

Backend will start at `http://localhost:8080`

#### 4. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Copy environment variables
cp .env.example .env

# Run the frontend
npm run dev
```

Frontend will start at `http://localhost:3000`

### 🐳 Docker Setup (Recommended)

```bash
# Copy environment file
cp .env.example .env

# Edit .env and add your Gemini API key
# GEMINI_API_KEY=your_actual_key_here
# DB_PASSWORD=your_secure_password

# Build and run all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

Services:
- Frontend: `http://localhost:3000`
- Backend: `http://localhost:8080`
- MySQL: `localhost:3306`

## 📡 API Endpoints

### Health Check
```http
GET /health
```

### Generate Image
```http
POST /api/images/generate
Content-Type: application/json

{
  "prompt": "A serene sunset over mountains"
}
```

### Get All Images
```http
GET /api/images
```

### Get Image by ID
```http
GET /api/images/:id
```

### Delete Image
```http
DELETE /api/images/:id
```

## 📁 Project Structure

```
gemini-image-generator/
├── backend/
│   ├── config/
│   │   └── database.go       # MySQL connection with GORM
│   ├── controllers/
│   │   └── image_controller.go
│   ├── models/
│   │   └── image.go          # GORM model
│   ├── routes/
│   │   └── routes.go
│   ├── main.go
│   ├── go.mod
│   └── .env.example
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── ImageCard.jsx
│   │   │   └── ImageGenerator.jsx
│   │   ├── App.jsx
│   │   ├── main.jsx
│   │   └── index.css
│   ├── package.json
│   └── vite.config.js
├── docker-compose.yml
└── README.md
```

## 🔧 Environment Variables

### Backend (.env)
```env
PORT=8080
GEMINI_API_KEY=your_gemini_api_key
GEMINI_API_URL=https://generativelanguage.googleapis.com/v1beta

# MySQL Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=gemini_image_generator

ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
```

### Docker Compose (.env)
```env
GEMINI_API_KEY=your_gemini_api_key
DB_USER=gemini
DB_PASSWORD=secure_password
```

## 🎯 Features in Detail

### Image Generation
- Uses Gemini 2.0 Flash experimental model
- Supports detailed prompts for precise image generation
- Real-time status updates (pending, completed, failed)
- Base64 encoded image storage in MySQL (LONGTEXT)

### Database Schema
```sql
CREATE TABLE images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    prompt TEXT NOT NULL,
    image_url LONGTEXT,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Image Management
- View all generated images in a responsive grid
- Download images in JPEG format
- Delete unwanted images
- Automatic timestamp tracking with GORM

### UI/UX
- Modern gradient design
- Smooth animations and transitions
- Loading states and error handling
- Example prompts for quick start
- Statistics dashboard

## 🐛 Troubleshooting

### Common Issues

**1. MySQL Connection Error**
```bash
# Check if MySQL is running
sudo systemctl status mysql
# or
mysqladmin -u root -p ping

# Start MySQL
sudo systemctl start mysql

# For Docker:
docker-compose up mysql
```

**2. Database Migration Error**
```bash
# GORM auto-migrates on startup
# If issues persist, manually create the database:
mysql -u root -p -e "CREATE DATABASE gemini_image_generator;"
```

**3. Gemini API Error**
- Verify your API key is correct
- Check if you have API quota remaining
- Ensure you're using the correct model name

**4. CORS Error**
- Update `ALLOWED_ORIGINS` in backend .env
- Restart the backend server

**5. Port Already in Use**
```bash
# Check what's using port 3306
sudo lsof -i :3306
# or
sudo netstat -tulpn | grep 3306

# Kill the process or change port in .env
```

## 📝 License

MIT License - feel free to use this project for learning and development!

## 👨‍💻 Author

**Avinash Kumar**
- GitHub: [@Avinashkr000](https://github.com/Avinashkr000)
- Backend Developer specializing in Go & Java

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ⭐ Show your support

Give a ⭐️ if you like this project!

---

<div align="center">

**Built with ❤️ using Go, React, MySQL, and Gemini AI**

</div>
