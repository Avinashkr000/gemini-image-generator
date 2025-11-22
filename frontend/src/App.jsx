import { useState, useEffect } from 'react'
import axios from 'axios'
import ImageGenerator from './components/ImageGenerator'
import ImageCard from './components/ImageCard'
import { Sparkles, Clock, CheckCircle, XCircle } from 'lucide-react'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

function App() {
  const [images, setImages] = useState([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const fetchImages = async () => {
    try {
      const response = await axios.get(`${API_URL}/images`)
      setImages(response.data.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt)))
    } catch (err) {
      console.error('Failed to fetch images:', err)
    }
  }

  useEffect(() => {
    fetchImages()
  }, [])

  const handleGenerateImage = async (prompt) => {
    setLoading(true)
    setError(null)
    try {
      const response = await axios.post(`${API_URL}/images/generate`, { prompt })
      await fetchImages()
      return response.data
    } catch (err) {
      const errorMsg = err.response?.data?.error || 'Failed to generate image'
      setError(errorMsg)
      throw new Error(errorMsg)
    } finally {
      setLoading(false)
    }
  }

  const handleDeleteImage = async (id) => {
    try {
      await axios.delete(`${API_URL}/images/${id}`)
      await fetchImages()
    } catch (err) {
      console.error('Failed to delete image:', err)
    }
  }

  const getStatusIcon = (status) => {
    switch (status) {
      case 'completed':
        return <CheckCircle className="w-4 h-4 text-green-500" />
      case 'pending':
        return <Clock className="w-4 h-4 text-yellow-500 animate-spin" />
      case 'failed':
        return <XCircle className="w-4 h-4 text-red-500" />
      default:
        return null
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-50 via-pink-50 to-blue-50">
      {/* Header */}
      <header className="bg-white/80 backdrop-blur-sm shadow-sm sticky top-0 z-10">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
          <div className="flex items-center justify-center space-x-3">
            <Sparkles className="w-8 h-8 text-purple-600" />
            <h1 className="text-3xl font-bold bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent">
              Gemini Image Generator
            </h1>
          </div>
          <p className="text-center text-gray-600 mt-2">Transform your ideas into stunning AI-generated images</p>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Image Generator */}
        <div className="mb-12">
          <ImageGenerator onGenerate={handleGenerateImage} loading={loading} />
          {error && (
            <div className="mt-4 bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg">
              <p className="font-medium">Error:</p>
              <p className="text-sm">{error}</p>
            </div>
          )}
        </div>

        {/* Stats */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <div className="bg-white rounded-xl shadow-sm p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-500 text-sm">Total Images</p>
                <p className="text-2xl font-bold text-gray-800">{images.length}</p>
              </div>
              <div className="bg-purple-100 p-3 rounded-lg">
                <Sparkles className="w-6 h-6 text-purple-600" />
              </div>
            </div>
          </div>
          <div className="bg-white rounded-xl shadow-sm p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-500 text-sm">Completed</p>
                <p className="text-2xl font-bold text-green-600">
                  {images.filter(img => img.status === 'completed').length}
                </p>
              </div>
              <div className="bg-green-100 p-3 rounded-lg">
                <CheckCircle className="w-6 h-6 text-green-600" />
              </div>
            </div>
          </div>
          <div className="bg-white rounded-xl shadow-sm p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-gray-500 text-sm">Pending</p>
                <p className="text-2xl font-bold text-yellow-600">
                  {images.filter(img => img.status === 'pending').length}
                </p>
              </div>
              <div className="bg-yellow-100 p-3 rounded-lg">
                <Clock className="w-6 h-6 text-yellow-600" />
              </div>
            </div>
          </div>
        </div>

        {/* Images Grid */}
        <div>
          <h2 className="text-2xl font-bold text-gray-800 mb-6">Generated Images</h2>
          {images.length === 0 ? (
            <div className="text-center py-16 bg-white rounded-xl shadow-sm">
              <Sparkles className="w-16 h-16 text-gray-300 mx-auto mb-4" />
              <p className="text-gray-500 text-lg">No images yet. Start creating!</p>
              <p className="text-gray-400 text-sm mt-2">Enter a prompt above to generate your first image</p>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {images.map((image) => (
                <ImageCard
                  key={image.id}
                  image={image}
                  onDelete={handleDeleteImage}
                  statusIcon={getStatusIcon(image.status)}
                />
              ))}
            </div>
          )}
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-white/80 backdrop-blur-sm mt-16 py-6">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center text-gray-600">
          <p>Powered by Google Gemini AI • Built with React & Go</p>
        </div>
      </footer>
    </div>
  )
}

export default App
