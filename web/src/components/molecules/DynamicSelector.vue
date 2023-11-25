<template>
  <div class="position-relative">
    <InputText v-if="!selected" type="text" name="subject" :label="title" @focusin="listVisible = true" v-model="query"
      @blur="closeList" :placeholder="placeholder" />
    <div v-else class="d-flex align-items-end">
      <InputText :label="title" class="flex-grow-1" :disabled="true" type="text" v-model="selected.name"
        name="subject-value" :placeholder="placeholder" />

      <button type="button" class="btn btn-danger d-inline clear-btn text-left" @click="selected = null">
        <font-awesome-icon icon="fa-solid fa-xmark" />
      </button>
    </div>

    <div v-show="listVisible" class="position-absolute end-0 start-0 dropdown">
      <ul v-if="subjectsList?.length" class="list-group">
        <li v-for="subject in subjectsList" :key="subject.id" class="list-group-item">
          <button @click="selectSubject(subject)" class="bg-white border-0 w-100 text-start">
            {{ subject.name }}
          </button>
        </li>
      </ul>
      <ul v-else class="list-group">
        <li class="list-group-item">Brak przedmiotu</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts" generic="T extends { id: number; name: string }">
import InputText from '@/components/atoms/InputText.vue'
import { computed, ref } from 'vue'

const $emit = defineEmits(['on-select'])
const query = ref('')
const selected = ref<T | null>()
const props = defineProps<{ items: T[]; title: string; placeholder: string }>()
const listVisible = ref(false)

const selectSubject = (sub: T) => {
  selected.value = sub
  $emit('on-select', sub)
}

const closeList = () => {
  setTimeout(() => {
    listVisible.value = false
  }, 200)
}

const subjectsList = computed(() => {
  return props.items?.filter((s) => s.name.toLowerCase().includes(query.value.toLowerCase()))
})
</script>

<style scoped lang="scss">
.clear-btn {
  max-height: 40px;
  margin-left: 12px;
}

.dropdown {
  // bottom: 10px;
}

.subject-btn {}
</style>
