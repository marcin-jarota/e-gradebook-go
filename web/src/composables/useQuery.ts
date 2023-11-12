export const useQuery = () => {
  const params = new URL(location.href).searchParams

  const getParam = (name: string): string | null => {
    return params.get(name)
  }

  return {
    getParam
  }
}
