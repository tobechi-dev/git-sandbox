export default function Home({ message }) {
  return (
    <div style={{ fontFamily: 'system-ui', padding: '2rem' }}>
      <h1>Git Sandbox</h1>
      <p>{message}</p>
      <p style={{ color: '#666', fontSize: '0.9rem' }}>
        If you can see this, Go + Inertia + React are talking. 🎉
      </p>
    </div>
  )
}
