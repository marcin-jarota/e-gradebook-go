import type { ApiBaseResponse } from '@/types'
import axios, { type AxiosResponse } from 'axios'

const client = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 1000,
  headers: {
    Authorization: `Bearer ${localStorage.getItem('token')}`,
    'Content-Type': 'application/json',
    asdas: '12'
  }
})

const unwrapRequestData = async <T>(
  promise: Promise<AxiosResponse<ApiBaseResponse<T>>>
): Promise<Awaited<ApiBaseResponse<T>>> => {
  const response = await promise
  if (response?.data?.error) {
    throw new Error(response?.data?.error)
  }

  return response?.data
}

export { unwrapRequestData }
export default client
