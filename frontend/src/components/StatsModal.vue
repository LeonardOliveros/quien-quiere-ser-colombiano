<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Mis Estadísticas</h5>
          <button type="button" class="btn-close btn-close-white" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div v-if="loading" class="text-center">
            <div class="spinner-border text-warning" role="status"></div>
          </div>

          <div v-else-if="stats">
            <!-- Overall Stats -->
            <div class="stats-section">
              <h6 class="text-golden">Estadísticas Generales</h6>
              <div class="row">
                <div class="col-md-3 col-6 mb-3">
                  <div class="stat-card">
                    <div class="stat-value">{{ stats.total_games }}</div>
                    <div class="stat-label">Juegos</div>
                  </div>
                </div>
                <div class="col-md-3 col-6 mb-3">
                  <div class="stat-card">
                    <div class="stat-value">{{ stats.average_score.toFixed(1) }}%</div>
                    <div class="stat-label">Promedio</div>
                  </div>
                </div>
                <div class="col-md-3 col-6 mb-3">
                  <div class="stat-card">
                    <div class="stat-value">{{ stats.best_score.toFixed(1) }}%</div>
                    <div class="stat-label">Mejor</div>
                  </div>
                </div>
                <div class="col-md-3 col-6 mb-3">
                  <div class="stat-card">
                    <div class="stat-value">{{ stats.total_questions }}</div>
                    <div class="stat-label">Preguntas</div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Category Stats -->
            <div class="stats-section mt-4" v-if="stats.category_stats">
              <h6 class="text-golden">Por Categoría</h6>
              <div v-for="(catStat, category) in stats.category_stats" :key="category" class="category-stat-item">
                <div class="d-flex justify-content-between align-items-center mb-2">
                  <strong>{{ category }}</strong>
                  <span class="badge bg-info">{{ catStat.percentage.toFixed(1) }}%</span>
                </div>
                <div class="progress" style="height: 20px;">
                  <div
                    class="progress-bar"
                    :style="{ width: catStat.percentage + '%' }"
                    :class="{
                      'bg-success': catStat.percentage >= 75,
                      'bg-warning': catStat.percentage >= 50 && catStat.percentage < 75,
                      'bg-danger': catStat.percentage < 50
                    }"
                  >
                    {{ catStat.correct }}/{{ catStat.total }}
                  </div>
                </div>
              </div>
            </div>

            <!-- Weak Areas -->
            <div class="stats-section mt-4" v-if="stats.weak_areas && stats.weak_areas.length > 0">
              <h6 class="text-golden">Áreas Débiles</h6>
              <ul class="list-unstyled">
                <li v-for="(area, index) in stats.weak_areas" :key="index" class="weak-area-item">
                  <i class="fas fa-exclamation-triangle text-danger"></i>
                  {{ area.category }} - {{ area.subcategory }}
                  <span class="badge bg-danger">{{ area.accuracy.toFixed(1) }}%</span>
                </li>
              </ul>
            </div>

            <!-- Strong Areas -->
            <div class="stats-section mt-4" v-if="stats.strong_areas && stats.strong_areas.length > 0">
              <h6 class="text-golden">Áreas Fuertes</h6>
              <ul class="list-unstyled">
                <li v-for="(area, index) in stats.strong_areas" :key="index" class="strong-area-item">
                  <i class="fas fa-check-circle text-success"></i>
                  {{ area.category }} - {{ area.subcategory }}
                  <span class="badge bg-success">{{ area.accuracy.toFixed(1) }}%</span>
                </li>
              </ul>
            </div>

            <!-- Reset Stats Button -->
            <div class="text-center mt-4">
              <button class="btn btn-danger" @click="confirmReset">
                <i class="fas fa-trash"></i> Resetear Estadísticas
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import api from '@/services/api'
import type { UserStats } from '@/types'

const emit = defineEmits(['close'])

const authStore = useAuthStore()
const loading = ref(true)
const stats = ref<UserStats | null>(null)

onMounted(async () => {
  try {
    const data = await api.getUserStats(authStore.userId!)
    stats.value = data
  } catch (error) {
    console.error('Error loading stats:', error)
  }
  loading.value = false
})

async function confirmReset() {
  if (confirm('¿Estás seguro de que quieres resetear todas tus estadísticas?')) {
    try {
      await api.resetStats(authStore.userId!)
      alert('Estadísticas reseteadas correctamente')
      emit('close')
    } catch (error) {
      alert('Error al resetear estadísticas')
    }
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1050;
  animation: fadeIn 0.3s ease;
}

.modal-content {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  border: 2px solid var(--gold-color);
}

.modal-header {
  border-bottom: 1px solid var(--gold-color);
}

.modal-title {
  color: var(--gold-color);
}

.text-golden {
  color: var(--gold-color);
  font-weight: bold;
}

.stats-section {
  margin-bottom: 20px;
}

.stat-card {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid var(--gold-color);
  border-radius: 10px;
  padding: 15px;
  text-align: center;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: var(--gold-color);
}

.stat-label {
  color: var(--text-light);
  font-size: 0.9rem;
  margin-top: 5px;
}

.category-stat-item {
  background: rgba(255, 255, 255, 0.05);
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 10px;
}

.weak-area-item, .strong-area-item {
  background: rgba(255, 255, 255, 0.05);
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 8px;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
