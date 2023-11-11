import { ref } from 'vue'
import { defineStore } from 'pinia'

type SnackbarMessage = {
  id: string | number
  variant?: 'success' | 'danger' | 'info'
  content: string
  title?: string
  autoHideDelay?: number
}

export const useSnackbarStore = defineStore('snackbar', () => {
  const snackbars = ref<SnackbarMessage[]>([])

  const createsnackbar = (
    content: string,
    cfg: Partial<Pick<SnackbarMessage, 'title' | 'autoHideDelay' | 'variant'>>
  ) => {
    const message = {
      id: Math.floor(Math.random() * 100),
      content,
      title: cfg?.title
    }
    snackbars.value = [...snackbars.value, message]

    if (cfg.autoHideDelay) {
      setTimeout(() => {
        closeSnackbar(message.id)
      }, cfg.autoHideDelay)
    }
  }

  const closeSnackbar = (id: number | string) => {
    snackbars.value = snackbars.value.filter((s) => s.id !== id)
  }

  return { createsnackbar, snackbars, closeSnackbar }
})
