import { useSnackbarStore } from '@/stores/snackbar'

export const useSnackbar = () => {
  const snakcbarStore = useSnackbarStore()

  const successSnackbar = (content: string, autoHide?: number) => {
    snakcbarStore.createsnackbar(content, {
      variant: 'success',
      ...(autoHide ? { autoHideDelay: autoHide } : {})
    })
  }

  const errorSnackbar = (content: string, autoHide?: number) => {
    snakcbarStore.createsnackbar(content, {
      variant: 'danger',
      ...(autoHide ? { autoHideDelay: autoHide } : {})
    })
  }

  return {
    errorSnackbar,
    successSnackbar
  }
}
