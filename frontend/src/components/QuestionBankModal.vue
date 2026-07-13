<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title"><i class="fas fa-database"></i> Base de Datos de Preguntas</h5>
          <button type="button" class="btn-close btn-close-white" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div v-if="loading" aria-busy="true">
            <span class="visually-hidden">Cargando base de preguntas...</span>
            <div class="bank-total">
              <SkeletonBlock width="90px" height="2.4rem" radius="6px" />
              <SkeletonBlock width="140px" height="1rem" radius="4px" />
            </div>
            <div class="category-list">
              <div v-for="n in 4" :key="n" class="category-block">
                <div class="category-toggle">
                  <SkeletonBlock width="16px" height="16px" radius="4px" />
                  <SkeletonBlock class="category-name" width="120px" height="1rem" radius="4px" />
                  <SkeletonBlock width="30px" height="1rem" radius="4px" />
                </div>
                <SkeletonBlock height="10px" radius="5px" />
              </div>
            </div>
          </div>

          <div v-else-if="counts">
            <div class="bank-total">
              <span class="bank-total-value">{{ counts.total }}</span>
              <span class="bank-total-label">preguntas en total</span>
            </div>

            <div class="category-list">
              <div v-for="cat in categoryRows" :key="cat.code" class="category-block">
                <button
                  class="category-toggle"
                  @click="toggleCategory(cat.code)"
                  :aria-expanded="expanded === cat.code"
                >
                  <i class="fas fa-chevron-right toggle-icon" :class="{ open: expanded === cat.code }"></i>
                  <span class="category-name">{{ cat.code }}</span>
                  <span class="category-count">{{ cat.count }}</span>
                </button>
                <div class="category-bar">
                  <div class="category-fill" :style="{ width: cat.share + '%' }"></div>
                </div>

                <ul v-if="expanded === cat.code" class="subcategory-list">
                  <li v-for="sub in cat.subcategories" :key="sub.subcategory" class="subcategory-row">
                    <span>{{ sub.subcategory }}</span>
                    <span class="subcategory-count">{{ sub.count }}</span>
                  </li>
                </ul>
              </div>
            </div>
          </div>

          <div v-else class="text-center text-muted">
            No se pudo cargar el conteo de preguntas.
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import api from '@/services/api'
import SkeletonBlock from '@/components/SkeletonBlock.vue'
import type { QuestionCount } from '@/types'

defineEmits(['close'])

const loading = ref(true)
const counts = ref<QuestionCount | null>(null)
const expanded = ref<string | null>(null)

const categoryRows = computed(() => {
  if (!counts.value) return []
  const bySubcategory = counts.value.by_subcategory ?? []
  const maxCount = Math.max(...Object.values(counts.value.by_category), 1)
  return Object.entries(counts.value.by_category)
    .sort(([a], [b]) => a.localeCompare(b))
    .map(([code, count]) => ({
      code,
      count,
      share: (count / maxCount) * 100,
      subcategories: bySubcategory
        .filter((sub) => sub.category === code)
        .sort((a, b) => b.count - a.count),
    }))
})

function toggleCategory(code: string) {
  expanded.value = expanded.value === code ? null : code
}

onMounted(async () => {
  try {
    counts.value = await api.getQuestionCount()
  } catch (error) {
    console.error('Error loading question counts:', error)
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

.bank-total {
  display: flex;
  align-items: baseline;
  gap: 10px;
  margin-bottom: 20px;
}

.bank-total-value {
  font-size: 2.4rem;
  font-weight: 800;
  color: var(--gold-color);
  line-height: 1;
}

.bank-total-label {
  color: rgba(253, 249, 238, 0.6);
}

.category-list {
  max-height: 420px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.category-toggle {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 10px;
  background: none;
  border: none;
  padding: 0 0 8px;
  color: var(--text-main);
  cursor: pointer;
  text-align: left;
}

.toggle-icon {
  font-size: 0.75rem;
  color: rgba(253, 249, 238, 0.5);
  transition: transform 0.2s ease;
}

.toggle-icon.open {
  transform: rotate(90deg);
}

.category-name {
  font-weight: 700;
  letter-spacing: 0.5px;
  flex: 1;
}

.category-count {
  font-weight: 700;
  font-variant-numeric: tabular-nums;
  color: var(--text-main);
}

.category-bar {
  background: rgba(0, 0, 0, 0.35);
  height: 10px;
  border-radius: 5px;
  overflow: hidden;
}

.category-fill {
  height: 100%;
  border-radius: 5px;
  background: var(--gold-color);
  transition: width 0.6s ease;
}

.subcategory-list {
  list-style: none;
  margin: 10px 0 0;
  padding: 6px 0 0 22px;
}

.subcategory-row {
  display: flex;
  justify-content: space-between;
  gap: 12px;
  padding: 5px 0;
  font-size: 0.92rem;
  color: rgba(253, 249, 238, 0.85);
}

.subcategory-row + .subcategory-row {
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

.subcategory-count {
  font-variant-numeric: tabular-nums;
  color: rgba(253, 249, 238, 0.6);
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
