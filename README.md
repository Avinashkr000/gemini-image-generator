# 🎨 Gemini Image Generator

<div align="center">

![Gemini Image Generator](https://img.shields.io/badge/Gemini-AI%20Powered-purple?style=for-the-badge&logo=google)
![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react)
![MongoDB](https://img.shields.io/badge/MongoDB-Atlas-47A248?style=for-the-badge&logo=mongodb)

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
- **MongoDB** - NoSQL database for image storage
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
- MongoDB (local or Atlas)
- Gemini API Key ([Get it here](https://ai.google.dev/))

### Installation

#### 1. Clone the repository

```bash
git clone https://github.com/Avinashkr000/gemini-image-generator.git
cd gemini-image-generator
```

#### 2. Backend Setup

```bash
cd backend

# Copy environment variables
cp .env.example .env

# Edit .env and add your Gemini API key
# GEMINI_API_KEY=your_actual_api_key_here

# Install dependencies
go mod download

# Run the backend
go run main.go
```

Backend will start at `http://localhost:8080`

#### 3. Frontend Setup

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
- MongoDB: `localhost:27017`

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
│   │   └── database.go
│   ├── controllers/
│   │   └── image_controller.go
│   ├── models/
│   │   └── image.go
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
MONGODB_URI=mongodb://localhost:27017/gemini-image-generator
MONGODB_DATABASE=gemini-image-generator
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
```

## 🎯 Features in Detail

### Image Generation
- Uses Gemini 2.0 Flash experimental model
- Supports detailed prompts for precise image generation
- Real-time status updates (pending, completed, failed)
- Base64 encoded image storage

### Image Management
- View all generated images in a responsive grid
- Download images in JPEG format
- Delete unwanted images
- Automatic timestamp tracking

### UI/UX
- Modern gradient design
- Smooth animations and transitions
- Loading states and error handling
- Example prompts for quick start
- Statistics dashboard

## 🐛 Troubleshooting

### Common Issues

**1. MongoDB Connection Error**
```bash
# Make sure MongoDB is running
docker-compose up mongodb
# or
mongod --dbpath /path/to/data
```

**2. Gemini API Error**
- Verify your API key is correct
- Check if you have API quota remaining
- Ensure you're using the correct model name

**3. CORS Error**
- Update `ALLOWED_ORIGINS` in backend .env
- Restart the backend server

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

**Built with ❤️ using Go, React, and Gemini AI**

</div>
