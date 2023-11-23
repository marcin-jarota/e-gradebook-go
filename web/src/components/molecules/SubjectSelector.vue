<template>
  <div class="position-relative">
    <InputText v-if="!selectedSubject" type="text" name="subject" label="Przedmiot" @focusin="listVisible = true"
      v-model="query" @blur="closeList" placeholder="Nazwa przedmiotu" />
    <div v-else class="d-flex align-items-end">
      <InputText label="Przedmiot" class="flex-grow-1" :disabled="true" v-model="selectedSubject.name" type="text"
        name="subject-value" placeholder="Nazwa przedmiotu" />

      <button type="button" class="btn btn-danger d-inline clear-btn text-left" @click="selectedSubject = null">
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

<script setup lang="ts">
import InputText from '@/components/atoms/InputText.vue'
import { subjectResource } from '@/resources/subject'
import type { Subject } from '@/types/Subject'
import { computed, ref } from 'vue'

const $emit = defineEmits(['on-subject-select'])
const query = ref('')
const selectedSubject = ref<Subject | null>(null)
const subjects = ref<Subject[]>()
const listVisible = ref(false)

const selectSubject = (sub: Subject) => {
  selectedSubject.value = sub
  $emit('on-subject-select', sub)
}

const closeList = () => {
  setTimeout(() => {
    listVisible.value = false
  }, 200)
}

const subjectsList = computed(() => {
  return subjects.value?.filter((s) => s.name.toLowerCase().includes(query.value.toLowerCase()))
})

const getSubjects = async () => {
  const { data } = await subjectResource.list()

  subjects.value = data
}

getSubjects()
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
