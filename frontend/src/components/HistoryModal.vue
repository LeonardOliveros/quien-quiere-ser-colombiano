<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Historial de Juegos</h5>
          <button type="button" class="btn-close btn-close-white" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div v-if="loading" class="text-center">
            <div class="spinner-border text-warning" role="status"></div>
          </div>

          <div v-else-if="history && history.length > 0">
            <div class="history-list">
              <div v-for="(game, index) in history" :key="index" class="history-item">
                <div class="row align-items-center">
                  <div class="col-md-3">
                    <small class="text-muted">{{ formatDate(game.created_at) }}</small>
                  </div>
                  <div class="col-md-3">
                    <span class="badge bg-info">{{ game.mode }}</span>
                  </div>
                  <div class="col-md-3">
                    {{ game.correct_answers }}/{{ game.total_questions }}
                  </div>
                  <div class="col-md-3">
                    <span class="badge" :class="{
                      'bg-success': game.score >= 70,
                      'bg-warning': game.score >= 50 && game.score < 70,
                      'bg-danger': game.score < 50
                    }">
                      {{ game.score }}%
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div v-else class="text-center text-muted">
            No hay historial de juegos todavía.
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
import type { GameHistory } from '@/types'

const emit = defineEmits(['close'])

const authStore = useAuthStore()
const loading = ref(true)
const history = ref<GameHistory[]>([])

onMounted(async () => {
  try {
    const data = await api.getUserHistory(authStore.userId!)
    history.value = data
  } catch (error) {
    console.error('Error loading history:', error)
  }
  loading.value = false
})

function formatDate(dateString: string): string {
  const date = new Date(dateString)
  return date.toLocaleDateString('es-ES', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
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

.history-list {
  max-height: 500px;
  overflow-y: auto;
}

.history-item {
  background: rgba(255, 255, 255, 0.05);
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 10px;
  border-left: 4px solid var(--gold-color);
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
