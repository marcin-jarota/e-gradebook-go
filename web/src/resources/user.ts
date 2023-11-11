import client from '@/lib/axios'
import type { ApiBaseResponse, UserListResponse, UserInput } from '@/types'

export const userResource = {
  async list() {
    const response = await client.get<UserListResponse>('/user/list')
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  },
  async create(payload: UserInput) {
    const response = await client.post(`/user/${payload.role}/create`, payload)
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  },
  async activate(userID: number) {
    const response = await client.get<ApiBaseResponse<{ ok: boolean }>>(`/user/activate/${userID}`)
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  },
  async deactivate(userID: number) {
    const response = await client.get<ApiBaseResponse>(`/user/deactivate/${userID}`)
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  },
  async destroySession(userID: number) {
    const response = await client.get<ApiBaseResponse<any>>(`/user/destroy-session/${userID}`)
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  }
}
