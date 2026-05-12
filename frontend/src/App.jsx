import { useEffect, useState } from 'react'

const choices = ['rock', 'paper', 'scissors']

function App() {
  const [choice, setChoice] = useState('rock')
  const [result, setResult] = useState(null)
  const [stats, setStats] = useState(null)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const fetchStats = async () => {
    try {
      const response = await fetch('/api/stats')
      if (!response.ok) throw new Error('Failed to load stats')
      const data = await response.json()
      setStats(data)
    } catch (err) {
      setError(err.message)
    }
  }

  useEffect(() => {
    fetchStats()
  }, [])

  const handlePlay = async event => {
    event.preventDefault()
    setLoading(true)
    setError('')
    setResult(null)

    try {
      const response = await fetch('/api/play', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ choice })
      })

      if (!response.ok) {
        const payload = await response.text()
        throw new Error(payload || 'Play request failed')
      }

      const payload = await response.json()
      setResult(payload)
      await fetchStats()
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="app-shell">
      <header>
        <h1>Rock Paper Scissors</h1>
        <p>Play against the computer and view game stats.</p>
      </header>

      <main>
        <section className="card">
          <h2>Play</h2>
          <form onSubmit={handlePlay}>
            <label htmlFor="choice">Choose your move</label>
            <select id="choice" value={choice} onChange={e => setChoice(e.target.value)}>
              {choices.map(item => (
                <option key={item} value={item}>
                  {item.charAt(0).toUpperCase() + item.slice(1)}
                </option>
              ))}
            </select>

            <button type="submit" disabled={loading}>
              {loading ? 'Playing...' : 'Play'}
            </button>
          </form>

          {error && <p className="error">{error}</p>}

          {result && (
            <div className="result">
              <p>
                <strong>You:</strong> {result.player_choice}
              </p>
              <p>
                <strong>Computer:</strong> {result.computer_choice}
              </p>
              <p>
                <strong>Result:</strong> {result.result}
              </p>
            </div>
          )}
        </section>

        <section className="card stats-card">
          <h2>Statistics</h2>
          {stats ? (
            <div>
              <p>Games Played: {stats.games_played}</p>
              <p>Wins: {stats.wins}</p>
              <p>Losses: {stats.losses}</p>
              <p>Draws: {stats.draws}</p>
              <p>Win Rate: {stats.win_percentage?.toFixed(1)}%</p>
            </div>
          ) : (
            <p>Loading stats...</p>
          )}
        </section>
      </main>
    </div>
  )
}

export default App
