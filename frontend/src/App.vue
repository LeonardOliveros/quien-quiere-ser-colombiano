<template>
  <div id="app">
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

const authStore = useAuthStore()
const currentYear = new Date().getFullYear()

onMounted(() => {
  authStore.checkAuth()
})
</script>

<style>
/* Global styles will be imported via main.ts */

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
