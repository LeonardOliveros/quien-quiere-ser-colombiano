<template>
  <div id="app">
    <button
      class="music-toggle"
      type="button"
      :aria-label="audioStore.muted ? 'Activar música' : 'Silenciar música'"
      @click="audioStore.toggleMute"
    >
      <i :class="audioStore.muted ? 'fas fa-volume-mute' : 'fas fa-volume-up'"></i>
    </button>
    <router-view />
    <footer class="app-footer">
      <div class="footer-content">
        <p class="footer-credits">
          © {{ currentYear }} <a href="https://github.com/LeonardOliveros" target="_blank" rel="noopener noreferrer">Leonardo Oliveros</a>. Todos los derechos reservados.
        </p>
        <p class="footer-tagline">
          Hecho con <span class="heart">❤</span> en Colombia 🇨🇴
        </p>
        <p class="footer-links">
          <a href="https://github.com/LeonardOliveros/quien-quiere-ser-colombiano" target="_blank" rel="noopener noreferrer">
            <i class="fab fa-github"></i> Código fuente
          </a>
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuthStore } from './stores/auth'
import { useAudioStore } from './stores/audio'

const authStore = useAuthStore()
const audioStore = useAudioStore()
const currentYear = new Date().getFullYear()

onMounted(() => {
  authStore.checkAuth()
  audioStore.init()
  // Browsers block audio-with-sound until a user gesture; start on the
  // first click/tap anywhere in the app instead of gating it behind a
  // specific button.
  document.addEventListener('pointerdown', () => audioStore.start(), { once: true })
})
</script>

<style>
/* Global styles will be imported via main.ts */

.music-toggle {
  position: fixed;
  top: 16px;
  right: 16px;
  z-index: 1000;
  width: 44px;
  height: 44px;
  border-radius: 50%;
  border: none;
  background: rgba(0, 0, 0, 0.35);
  color: var(--flag-yellow);
  font-size: 1.1rem;
  cursor: pointer;
  transition: background 0.2s ease;
}

.music-toggle:hover {
  background: rgba(0, 0, 0, 0.55);
}

.app-footer {
  background: rgba(0, 0, 0, 0.35);
  border-top: 4px solid transparent;
  border-image: linear-gradient(
    to right,
    var(--flag-yellow) 0 50%,
    var(--flag-blue-light) 50% 75%,
    var(--flag-red) 75% 100%
  ) 1;
  padding: 24px 20px;
  text-align: center;
  color: var(--text-main);
}

.app-footer p {
  margin: 4px 0;
}

.footer-credits {
  font-weight: 600;
}

.footer-credits a {
  color: var(--flag-yellow);
  text-decoration: none;
}

.footer-credits a:hover {
  text-decoration: underline;
}

.footer-tagline {
  font-size: 0.9rem;
  opacity: 0.85;
}

.footer-tagline .heart {
  color: var(--flag-red-light);
}

.footer-links {
  font-size: 0.9rem;
}

.footer-links a {
  color: var(--flag-blue-light);
  text-decoration: none;
  transition: color 0.2s ease;
}

.footer-links a:hover {
  color: var(--flag-yellow);
}
</style>
