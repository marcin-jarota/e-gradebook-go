<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import { useSessionStore } from '@/stores/session'
import { computed } from 'vue'

const sessionStore = useSessionStore()

fetch(import.meta.env.VITE_API_BASE_URL + '/students/1/marks', {
  headers: {
    Authorization: `Bearer ${localStorage.getItem('token')}`,
    'Content-Type': 'application/json'
  }
})
  .then((r) => r.json())
  .then(console.log)
const welcomeMessage = computed(() => {
  const hour = new Date().getHours()
  if (hour >= 5 && hour < 12) {
    return `Dzień dobry, ${sessionStore.user?.name} 🌞`
  } else if (hour >= 12 && hour < 18) {
    return `Witaj, ${sessionStore.user?.name} ☕`
  } else {
    return `Dobry wieczór, ${sessionStore.user?.name} 🌛`
  }
})

const options = {
  chart: {
    width: 380,
    type: 'pie'
  },
  labels: ['Team A', 'Team B', 'Team C', 'Team D', 'Team E'],
  responsive: [
    {
      breakpoint: 480,
      options: {
        chart: {
          width: 200
        },
        legend: {
          position: 'bottom'
        }
      }
    }
  ]
}

const series = [44, 55, 13, 43, 22]
</script>

<template>
  <MainLayout>
    <h1>{{ welcomeMessage }}</h1>
    <apex-chart width="500" type="pie" :options="options" :series="series"></apex-chart>
  </MainLayout>
</template>
