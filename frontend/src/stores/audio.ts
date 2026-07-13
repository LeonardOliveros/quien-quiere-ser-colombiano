import { defineStore } from 'pinia'
import { ref } from 'vue'

const MUTED_STORAGE_KEY = 'bgMusicMuted'
// Same-origin path: CloudFront routes /media/* to a dedicated S3 bucket in
// prod (see infra/lib/quiz-app-stack.ts); Vite's dev proxy forwards it to
// the deployed site so `npm run dev` also has music (see vite.config.ts).
const AUDIO_SRC = '/media/audio/himno-nacional-instrumental-v1.mp3'

export const useAudioStore = defineStore('audio', () => {
  const muted = ref(localStorage.getItem(MUTED_STORAGE_KEY) === 'true')
  let audio: HTMLAudioElement | null = null

  function init(): void {
    if (audio) return
    audio = new Audio(AUDIO_SRC)
    audio.loop = true
    audio.volume = 0.35
    audio.muted = muted.value
  }

  async function start(): Promise<void> {
    init()
    if (muted.value) return
    try {
      await audio!.play()
    } catch {
      // Blocked by the browser's autoplay-with-sound policy until a user
      // gesture happens; the caller retries this on the next one.
    }
  }

  function toggleMute(): void {
    init()
    muted.value = !muted.value
    audio!.muted = muted.value
    localStorage.setItem(MUTED_STORAGE_KEY, String(muted.value))
    if (!muted.value) void start()
  }

  return { muted, init, start, toggleMute }
})
