<template>
  <div class="week-schedule">
    <!-- Time Column -->
    <div class="col time-column">
      <div class="time-slot" v-for="time in times" :key="time">{{ time }}</div>
    </div>

    <!-- Day Columns -->
    <div class="col day-column" v-for="day in daysOfWeek" :key="day">
      <div class="day-label">{{ day }}</div>
      <div class="lessons-container">
        <div class="lesson" v-for="lesson in filteredLessons(daysOfWeek.indexOf(day))" :key="lesson.id"
          :style="getLessonStyle(lesson)" @mouseenter="showFullLesson(lesson.id)" @mouseleave="hideFullLesson">
          <strong>{{ lesson.subject }}</strong><br />
          <div v-if="activeLesson === lesson.id" class="full-lesson">
            <p>{{ lesson.teacher }}</p>
            {{ lesson.start }} - {{ lesson.end }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { type Lesson } from '@/types/ClassGroup'

const props = defineProps<{ lessons: Lesson[] }>()

const times = ref(Array.from({ length: 11 }, (_, i) => `${7 + i}:00`))
const daysOfWeek = ['Poniedziałek', 'Wtorek', 'Środa', 'Czwartek', 'Piątek']
const activeLesson = ref<number | null>(null)

const filteredLessons = (dayOfWeek: number) => {
  return (props.lessons || []).filter((lesson: Lesson) => lesson.dayOfWeek === dayOfWeek + 1)
}
const getLessonStyle = (lesson: Lesson) => {
  const startHour = parseInt(lesson.start.split(':')[0])
  const startMinutes = parseInt(lesson.start.split(':')[1])
  const endHour = parseInt(lesson.end.split(':')[0])
  const endMinutes = parseInt(lesson.end.split(':')[1])
  const duration = endHour + endMinutes / 60 - (startHour + startMinutes / 60)
  const top = (startHour - 7 + startMinutes / 60) * 60 // assuming each hour block is 60px height
  const height = duration * 60 // assuming each hour block is 60px height

  return {
    position: 'absolute',
    top: `${top}px`,
    height: `${height}px`,
    backgroundColor: '#f8f9fa',
    border: '1px solid #dee2e6',
    padding: '0.5rem',
    borderRadius: '0.25rem',
    width: '90%',
    left: '50%',
    transform: 'translateX(-50%)'
  }
}

const showFullLesson = (id: number) => {
  activeLesson.value = id
}

const hideFullLesson = () => {
  activeLesson.value = null
}
</script>

<style scoped lang="scss">
.week-schedule {
  display: flex;
}

.time-column,
.day-column {
  flex: 1;
  position: relative;
}

.time-column {
  top: 25px;
}

.time-slot {
  text-align: right;
  padding-right: 10px;
  border-bottom: 1px solid #dee2e6;
  height: 60px;
  line-height: 60px;
}

.time-slot:hover {
  height: max-content;
}

.day-label {
  text-align: center;
  font-weight: bold;
  border-bottom: 1px solid #dee2e6;
}

.lessons-container {
  min-height: 660px;
  /* 11 hours * 60px */
  position: relative;
}

.lesson {
  /* ... Previous styles */
  cursor: pointer;
  transition: all 0.3s ease;
}

.lesson:hover {
  z-index: 10;
  background-color: #e9ecef;
}

.full-lesson {
  position: absolute;
  top: 0;
  left: 105%;
  right: -100%;
  background-color: white;
  border: 1px solid #dee2e6;
  border-radius: 0.25rem;
  padding: 0.5rem;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
  z-index: 9;
  /* Below the hovered lesson */
}

/* Add borders between days */
.day-column:not(:last-child) {
  border-right: 1px solid #dee2e6;
}
</style>
