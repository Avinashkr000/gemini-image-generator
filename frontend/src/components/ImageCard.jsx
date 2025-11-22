import { Trash2, Calendar, Download } from 'lucide-react'

function ImageCard({ image, onDelete, statusIcon }) {
  const handleDownload = () => {
    const link = document.createElement('a')
    link.href = image.imageUrl
    link.download = `gemini-${image.id}.jpg`
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  return (
    <div className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300">
      {/* Image */}
      <div className="relative aspect-square bg-gray-100 overflow-hidden">
        {image.status === 'completed' && image.imageUrl ? (
          <img
            src={image.imageUrl}
            alt={image.prompt}
            className="w-full h-full object-cover"
          />
        ) : (
          <div className="w-full h-full flex items-center justify-center">
            <div className="text-center p-6">
              {statusIcon}
              <p className="mt-2 text-sm text-gray-500 capitalize">{image.status}</p>
            </div>
          </div>
        )}
        
        {/* Status Badge */}
        <div className="absolute top-3 right-3">
          <div className="bg-white/90 backdrop-blur-sm rounded-full px-3 py-1 flex items-center space-x-1">
            {statusIcon}
            <span className="text-xs font-medium capitalize">{image.status}</span>
          </div>
        </div>
      </div>

      {/* Content */}
      <div className="p-4">
        <p className="text-gray-800 font-medium line-clamp-2 mb-3">
          {image.prompt}
        </p>
        
        <div className="flex items-center text-xs text-gray-500 mb-4">
          <Calendar className="w-3 h-3 mr-1" />
          {formatDate(image.createdAt)}
        </div>

        {/* Actions */}
        <div className="flex space-x-2">
          {image.status === 'completed' && (
            <button
              onClick={handleDownload}
              className="flex-1 bg-purple-600 hover:bg-purple-700 text-white py-2 px-4 rounded-lg transition-colors duration-200 flex items-center justify-center space-x-2"
            >
              <Download className="w-4 h-4" />
              <span>Download</span>
            </button>
          )}
          <button
            onClick={() => onDelete(image.id)}
            className="bg-red-50 hover:bg-red-100 text-red-600 p-2 rounded-lg transition-colors duration-200"
          >
            <Trash2 className="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>
  )
}

export default ImageCard
