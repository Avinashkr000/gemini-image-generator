import { useState } from 'react'
import { Sparkles, Loader2 } from 'lucide-react'

function ImageGenerator({ onGenerate, loading }) {
  const [prompt, setPrompt] = useState('')

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!prompt.trim()) return

    try {
      await onGenerate(prompt)
      setPrompt('')
    } catch (error) {
      console.error('Generation failed:', error)
    }
  }

  const examplePrompts = [
    'A serene sunset over mountains with vibrant colors',
    'A futuristic city with flying cars and neon lights',
    'A magical forest with glowing mushrooms and fairies',
    'An underwater scene with colorful coral and tropical fish'
  ]

  const handleExampleClick = (example) => {
    setPrompt(example)
  }

  return (
    <div className="bg-white rounded-2xl shadow-lg p-6 md:p-8">
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="prompt" className="block text-sm font-medium text-gray-700 mb-2">
            Describe your image
          </label>
          <textarea
            id="prompt"
            value={prompt}
            onChange={(e) => setPrompt(e.target.value)}
            placeholder="Enter a detailed description of the image you want to generate..."
            rows={4}
            className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all resize-none"
            disabled={loading}
          />
        </div>

        <button
          type="submit"
          disabled={loading || !prompt.trim()}
          className="w-full bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 text-white font-medium py-3 px-6 rounded-xl transition-all duration-200 flex items-center justify-center space-x-2 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {loading ? (
            <>
              <Loader2 className="w-5 h-5 animate-spin" />
              <span>Generating...</span>
            </>
          ) : (
            <>
              <Sparkles className="w-5 h-5" />
              <span>Generate Image</span>
            </>
          )}
        </button>
      </form>

      {/* Example Prompts */}
      <div className="mt-6">
        <p className="text-sm font-medium text-gray-700 mb-3">Try these examples:</p>
        <div className="flex flex-wrap gap-2">
          {examplePrompts.map((example, index) => (
            <button
              key={index}
              onClick={() => handleExampleClick(example)}
              className="text-xs bg-purple-50 hover:bg-purple-100 text-purple-700 px-3 py-2 rounded-lg transition-colors duration-200"
              disabled={loading}
            >
              {example}
            </button>
          ))}
        </div>
      </div>
    </div>
  )
}

export default ImageGenerator
