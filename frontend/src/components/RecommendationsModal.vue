<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Recomendaciones de Estudio</h5>
          <button type="button" class="btn-close btn-close-white" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div v-if="loading" class="text-center">
            <div class="spinner-border text-warning" role="status"></div>
          </div>

          <div v-else-if="recommendations && recommendations.length > 0">
            <div class="recommendations-list">
              <div v-for="(rec, index) in recommendations" :key="index" class="recommendation-item">
                <h6 class="text-warning">{{ rec.category }} - {{ rec.subcategory }}</h6>
                <p>{{ rec.description }}</p>
                <div class="d-flex justify-content-between align-items-center">
                  <span class="badge bg-warning text-dark">Prioridad: {{ rec.priority }}/5</span>
                  <small class="text-muted" v-if="rec.resources">{{ rec.resources }}</small>
                </div>
              </div>
            </div>
          </div>

          <div v-else class="text-center text-muted">
            No hay recomendaciones disponibles. Completa algunos juegos para obtener recomendaciones personalizadas.
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
import type { Recommendation } from '@/types'

const emit = defineEmits(['close'])

const authStore = useAuthStore()
const loading = ref(true)
const recommendations = ref<Recommendation[]>([])

onMounted(async () => {
  try {
    const data = await api.getRecommendations(authStore.userId!)
    recommendations.value = data
  } catch (error) {
    console.error('Error loading recommendations:', error)
  }
  loading.value = false
})
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

.recommendations-list {
  max-height: 500px;
  overflow-y: auto;
}

.recommendation-item {
  background: linear-gradient(145deg, var(--accent-color), var(--secondary-color));
  border-left: 4px solid var(--warning-color);
  padding: 15px;
  margin-bottom: 15px;
  border-radius: 5px;
}

.recommendation-item h6 {
  color: var(--warning-color);
  margin-bottom: 10px;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
