<template>
  <div class="login-container">
    <div class="container">
      <div class="row justify-content-center align-items-center min-vh-100">
        <div class="col-md-6">
          <div class="auth-card">
            <h1 class="text-center mb-4 text-golden">¿Quién Quiere Ser Colombiano?</h1>

            <!-- Login Form -->
            <div v-if="!showRegister" class="auth-form">
              <h3 class="text-center mb-4">Iniciar Sesión</h3>
              <form @submit.prevent="handleLogin">
                <div class="mb-3">
                  <label for="loginUsername" class="form-label">Usuario</label>
                  <input
                    type="text"
                    class="form-control"
                    id="loginUsername"
                    v-model="loginForm.username"
                    required
                  >
                </div>
                <div class="mb-3">
                  <label for="loginPassword" class="form-label">Contraseña</label>
                  <div class="password-wrapper">
                    <input
                      :type="showLoginPassword ? 'text' : 'password'"
                      class="form-control"
                      id="loginPassword"
                      v-model="loginForm.password"
                      required
                    >
                    <button
                      type="button"
                      class="password-toggle"
                      @click="showLoginPassword = !showLoginPassword"
                      :aria-label="showLoginPassword ? 'Ocultar contraseña' : 'Mostrar contraseña'"
                      tabindex="-1"
                    >
                      <i :class="showLoginPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                    </button>
                  </div>
                </div>
                <div v-if="guestExpired" class="alert alert-warning">
                  Tu sesión de invitado expiró. Puedes jugar de nuevo como invitado o crear una cuenta para conservar tu progreso.
                </div>
                <div v-if="loginError" class="alert alert-danger">{{ loginError }}</div>
                <button type="submit" class="btn btn-primary w-100" :disabled="isLoading">
                  <span v-if="isLoading">Cargando...</span>
                  <span v-else>Iniciar Sesión</span>
                </button>
              </form>
              <button
                type="button"
                class="btn btn-guest w-100 mt-3"
                :disabled="isLoading"
                @click="handleGuestLogin"
              >
                <span v-if="isLoading">Cargando...</span>
                <span v-else><i class="fas fa-user-clock me-2"></i>Jugar como invitado</span>
              </button>
              <p class="text-center guest-hint mt-2">
                Sin registro. Tu progreso como invitado se guarda solo por 24 horas.
              </p>
              <p class="text-center mt-3">
                ¿No tienes cuenta?
                <a href="#" @click.prevent="showRegister = true" class="text-golden">Regístrate</a>
              </p>
            </div>

            <!-- Register Form -->
            <div v-else class="auth-form">
              <h3 class="text-center mb-4">Registrarse</h3>
              <form @submit.prevent="handleRegister">
                <div class="mb-3">
                  <label for="registerUsername" class="form-label">Usuario</label>
                  <input
                    type="text"
                    class="form-control"
                    id="registerUsername"
                    v-model="registerForm.username"
                    required
                  >
                </div>
                <div class="mb-3">
                  <label for="registerPassword" class="form-label">Contraseña</label>
                  <div class="password-wrapper">
                    <input
                      :type="showRegisterPassword ? 'text' : 'password'"
                      class="form-control"
                      id="registerPassword"
                      v-model="registerForm.password"
                      minlength="8"
                      required
                    >
                    <button
                      type="button"
                      class="password-toggle"
                      @click="showRegisterPassword = !showRegisterPassword"
                      :aria-label="showRegisterPassword ? 'Ocultar contraseña' : 'Mostrar contraseña'"
                      tabindex="-1"
                    >
                      <i :class="showRegisterPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                    </button>
                  </div>
                  <small class="form-hint">Mínimo 8 caracteres</small>
                </div>
                <div class="mb-3">
                  <label for="registerConfirmPassword" class="form-label">Confirmar contraseña</label>
                  <div class="password-wrapper">
                    <input
                      :type="showConfirmPassword ? 'text' : 'password'"
                      class="form-control"
                      id="registerConfirmPassword"
                      v-model="confirmPassword"
                      required
                    >
                    <button
                      type="button"
                      class="password-toggle"
                      @click="showConfirmPassword = !showConfirmPassword"
                      :aria-label="showConfirmPassword ? 'Ocultar contraseña' : 'Mostrar contraseña'"
                      tabindex="-1"
                    >
                      <i :class="showConfirmPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                    </button>
                  </div>
                  <small v-if="confirmPassword && confirmPassword !== registerForm.password" class="text-mismatch">
                    Las contraseñas no coinciden
                  </small>
                </div>
                <div v-if="registerError" class="alert alert-danger">{{ registerError }}</div>
                <div v-if="registerSuccess" class="alert alert-success">{{ registerSuccess }}</div>
                <button type="submit" class="btn btn-primary w-100" :disabled="isLoading">
                  <span v-if="isLoading">Cargando...</span>
                  <span v-else>Registrarse</span>
                </button>
              </form>
              <p class="text-center mt-3">
                ¿Ya tienes cuenta?
                <a href="#" @click.prevent="showRegister = false" class="text-golden">Inicia sesión</a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const guestExpired = computed(() => route.query.expired === 'guest')

const showRegister = ref(false)
const isLoading = ref(false)
const showLoginPassword = ref(false)
const showRegisterPassword = ref(false)
const showConfirmPassword = ref(false)
const confirmPassword = ref('')

const loginForm = ref({
  username: '',
  password: ''
})

const registerForm = ref({
  username: '',
  password: ''
})

const loginError = ref('')
const registerError = ref('')
const registerSuccess = ref('')

async function handleLogin() {
  isLoading.value = true
  loginError.value = ''

  const result = await authStore.login(loginForm.value)

  if (result.success) {
    router.push('/')
  } else {
    loginError.value = result.message
  }

  isLoading.value = false
}

async function handleGuestLogin() {
  isLoading.value = true
  loginError.value = ''

  const result = await authStore.loginAsGuest()

  if (result.success) {
    router.push('/')
  } else {
    loginError.value = result.message
  }

  isLoading.value = false
}

async function handleRegister() {
  registerError.value = ''
  registerSuccess.value = ''

  if (registerForm.value.password.length < 8) {
    registerError.value = 'La contraseña debe tener al menos 8 caracteres'
    return
  }

  if (registerForm.value.password !== confirmPassword.value) {
    registerError.value = 'Las contraseñas no coinciden'
    return
  }

  isLoading.value = true

  const result = await authStore.register(registerForm.value)

  if (result.success) {
    confirmPassword.value = ''

    const loginResult = await authStore.login({
      username: registerForm.value.username,
      password: registerForm.value.password
    })

    if (loginResult.success) {
      router.push('/')
    } else {
      registerSuccess.value = result.message + ' Ahora puedes iniciar sesión.'
      setTimeout(() => {
        showRegister.value = false
        registerSuccess.value = ''
      }, 2000)
    }
  } else {
    registerError.value = result.message
  }

  isLoading.value = false
}
</script>

<style scoped>
.login-container {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  min-height: 100vh;
}

.auth-card {
  background: rgba(255, 255, 255, 0.1);
  border: 2px solid var(--gold-color);
  border-radius: 20px;
  padding: 40px;
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.5);
}

.auth-form {
  animation: fadeIn 0.5s ease-in;
}

.text-golden {
  color: var(--gold-color);
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.form-control {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid var(--gold-color);
  color: var(--text-light);
  padding: 12px;
  border-radius: 8px;
}

.form-control:focus {
  background: rgba(255, 255, 255, 0.15);
  border-color: var(--gold-color);
  color: var(--text-light);
  box-shadow: 0 0 10px rgba(255, 215, 0, 0.3);
}

.form-label {
  color: var(--text-light);
  font-weight: 500;
}

.password-wrapper {
  position: relative;
}

.password-wrapper .form-control {
  padding-right: 44px;
}

.password-toggle {
  position: absolute;
  top: 50%;
  right: 12px;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--gold-color);
  cursor: pointer;
  padding: 4px;
  line-height: 1;
}

.password-toggle:hover {
  color: var(--text-light);
}

.text-mismatch {
  color: #ff8a8a;
  display: block;
  margin-top: 4px;
}

.form-hint {
  color: rgba(255, 255, 255, 0.6);
  display: block;
  margin-top: 4px;
}

.btn-primary {
  background: linear-gradient(145deg, var(--accent-color), var(--secondary-color));
  border: 2px solid var(--gold-color);
  color: var(--text-light);
  padding: 12px;
  border-radius: 10px;
  font-weight: bold;
  transition: all 0.3s ease;
}

.btn-primary:hover:not(:disabled) {
  background: linear-gradient(145deg, var(--secondary-color), var(--accent-color));
  box-shadow: 0 5px 20px rgba(255, 215, 0, 0.3);
  transform: translateY(-2px);
}

.btn-primary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-guest {
  background: transparent;
  border: 2px solid var(--gold-color);
  color: var(--gold-color);
  padding: 12px;
  border-radius: 10px;
  font-weight: bold;
  transition: all 0.3s ease;
}

.btn-guest:hover:not(:disabled) {
  background: rgba(255, 215, 0, 0.15);
  box-shadow: 0 5px 20px rgba(255, 215, 0, 0.2);
  transform: translateY(-2px);
  color: var(--gold-color);
}

.btn-guest:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.guest-hint {
  color: rgba(255, 255, 255, 0.6);
  font-size: 0.85rem;
  margin-bottom: 0;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
