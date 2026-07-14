<template>
  <div class="admin-container">
    <div class="container py-5">
      <div class="d-flex justify-content-between align-items-center mb-4">
        <h1 class="text-golden mb-0">Panel Admin</h1>
        <button class="btn btn-outline-gold" @click="router.push('/')">
          <i class="fas fa-arrow-left"></i> Volver
        </button>
      </div>

      <div v-if="error" class="alert alert-danger">{{ error }}</div>

      <div v-else-if="loading" class="text-center text-light py-5">
        <i class="fas fa-spinner fa-spin fa-2x"></i>
      </div>

      <template v-else-if="metrics">
        <div class="stat-cards">
          <div class="stat-card">
            <div class="stat-value">{{ metrics.totals.registered_users }}</div>
            <div class="stat-label"><i class="fas fa-user me-1"></i> Usuarios registrados</div>
          </div>
          <div class="stat-card">
            <div class="stat-value">{{ metrics.totals.guest_users }}</div>
            <div class="stat-label"><i class="fas fa-user-clock me-1"></i> Invitados (histórico)</div>
          </div>
          <div class="stat-card">
            <div class="stat-value">{{ metrics.totals.total_games }}</div>
            <div class="stat-label"><i class="fas fa-gamepad me-1"></i> Partidas jugadas</div>
          </div>
        </div>

        <div class="d-flex justify-content-between align-items-center mt-5 mb-3">
          <h4 class="text-light mb-0">Actividad diaria</h4>
          <div class="days-selector">
            <button
              v-for="option in [7, 14, 30]"
              :key="option"
              class="btn btn-sm"
              :class="days === option ? 'btn-gold' : 'btn-outline-gold'"
              @click="changeDays(option)"
            >
              {{ option }} días
            </button>
          </div>
        </div>

        <div class="table-wrapper">
          <table class="metrics-table">
            <thead>
              <tr>
                <th>Fecha</th>
                <th>Jugadores activos</th>
                <th>Partidas</th>
                <th>Nuevos invitados</th>
                <th>Nuevos usuarios</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="day in metrics.daily" :key="day.date">
                <td>{{ day.date }}</td>
                <td>{{ day.active_users }}</td>
                <td>{{ day.games_started }}</td>
                <td>{{ day.new_guests }}</td>
                <td>{{ day.new_users }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/services/api'
import type { AdminMetrics } from '@/types'

const router = useRouter()

const metrics = ref<AdminMetrics | null>(null)
const loading = ref(true)
const error = ref('')
const days = ref(14)

async function loadMetrics() {
  loading.value = true
  error.value = ''
  try {
    metrics.value = await api.getAdminMetrics(days.value)
  } catch (err: any) {
    if (err.response?.status === 404) {
      // The server hides admin endpoints from non-admins
      router.push('/')
      return
    }
    error.value = err.response?.data?.error || 'No se pudieron cargar las métricas'
  } finally {
    loading.value = false
  }
}

function changeDays(value: number) {
  days.value = value
  loadMetrics()
}

onMounted(loadMetrics)
</script>

<style scoped>
.admin-container {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  min-height: 100vh;
}

.text-golden {
  color: var(--gold-color);
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.stat-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.stat-card {
  background: rgba(255, 255, 255, 0.1);
  border: 2px solid var(--gold-color);
  border-radius: 15px;
  padding: 24px;
  text-align: center;
}

.stat-value {
  color: var(--gold-color);
  font-size: 2.5rem;
  font-weight: bold;
}

.stat-label {
  color: var(--text-light);
  margin-top: 4px;
}

.days-selector {
  display: flex;
  gap: 8px;
}

.btn-outline-gold {
  background: transparent;
  border: 2px solid var(--gold-color);
  color: var(--gold-color);
  border-radius: 10px;
  transition: all 0.3s ease;
}

.btn-outline-gold:hover {
  background: var(--gold-color);
  color: var(--primary-color);
}

.btn-gold {
  background: var(--gold-color);
  border: 2px solid var(--gold-color);
  color: var(--primary-color);
  border-radius: 10px;
  font-weight: bold;
}

.table-wrapper {
  overflow-x: auto;
  border: 1px solid var(--gold-color);
  border-radius: 12px;
}

.metrics-table {
  width: 100%;
  color: var(--text-light);
  border-collapse: collapse;
}

.metrics-table th,
.metrics-table td {
  padding: 10px 16px;
  text-align: center;
  white-space: nowrap;
}

.metrics-table th {
  background: rgba(255, 215, 0, 0.12);
  color: var(--gold-color);
}

.metrics-table tbody tr:nth-child(odd) {
  background: rgba(255, 255, 255, 0.05);
}

.text-light {
  color: var(--text-light);
}
</style>
