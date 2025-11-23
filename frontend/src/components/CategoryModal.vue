<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Selecciona una Categoría</h5>
          <button type="button" class="btn-close btn-close-white" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div class="category-grid">
            <button
              v-for="category in categories"
              :key="category"
              class="btn-category"
              @click="selectCategory(category)"
            >
              <i :class="getCategoryIcon(category)"></i>
              <div>{{ category }}</div>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const emit = defineEmits(['close', 'select'])

const categories = ['CULTURA', 'GEOGRAFIA', 'HISTORIA', 'CONSTITUCION']

function selectCategory(category: string) {
  emit('select', category)
  emit('close')
}

function getCategoryIcon(category: string): string {
  const icons: Record<string, string> = {
    'CULTURA': 'fas fa-palette',
    'GEOGRAFIA': 'fas fa-globe-americas',
    'HISTORIA': 'fas fa-landmark',
    'CONSTITUCION': 'fas fa-balance-scale'
  }
  return icons[category] || 'fas fa-question'
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

.category-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 15px;
}

.btn-category {
  background: linear-gradient(145deg, var(--accent-color), var(--secondary-color));
  color: var(--text-light);
  border: 2px solid var(--gold-color);
  padding: 30px 20px;
  border-radius: 15px;
  transition: all 0.3s ease;
  cursor: pointer;
  font-weight: bold;
}

.btn-category:hover {
  background: linear-gradient(145deg, var(--secondary-color), var(--accent-color));
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(255, 215, 0, 0.3);
  color: var(--gold-color);
}

.btn-category i {
  font-size: 2rem;
  display: block;
  margin-bottom: 10px;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
