import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { writeFileSync, rmSync, existsSync, mkdirSync } from 'node:fs'
import { dirname } from 'node:path'

const HOT_FILE = '.vite/hot'

// Writes a "hot" file when the dev server starts, deletes it on exit.
// gonertia reads this file to know when to point the browser at Vite's dev server.
function hotFilePlugin() {
  return {
    name: 'hot-file',
    apply: 'serve',
    configureServer(server) {
      server.httpServer?.once('listening', () => {
        const addr = server.httpServer.address()
        const host = addr.address === '::' ? 'localhost' : addr.address
        mkdirSync(dirname(HOT_FILE), { recursive: true })
        writeFileSync(HOT_FILE, `http://${host}:${addr.port}`)
        console.log(`\n  ➜  hot file written: ${HOT_FILE}\n`)
      })
      server.httpServer?.once('close', () => {
        if (existsSync(HOT_FILE)) rmSync(HOT_FILE)
      })
    },
  }
}

export default defineConfig({
  plugins: [react(), hotFilePlugin()],
  server: {
    port: 5173,
    host: true, // listen on 0.0.0.0 so the browser (same device) can reach it
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    manifest: true,
    rollupOptions: {
      input: 'src/app.jsx',
    },
  },
})
