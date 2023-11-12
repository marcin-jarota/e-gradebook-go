import client from '@/lib/axios'
import type { ApiBaseResponse, UserListResponse, UserInput, SetupPasswordPayload } from '@/types'

export const userResource = {
  async list() {
    const response = await client.get<UserListResponse>('/user/list')
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  },
  async tokenValid() {
    return client.get('/token-valid')
  },
  async setupPassword(payload: SetupPasswordPayload, token: string) {
    const response = await client.post(`/setup-password?token=${token}`, payload)
    if (response.data.error) {
      throw new Error(response.data.error)
    }
    return response.data
  },
  async create(payload: UserInput) {
    const response = await client.post<ApiBaseResponse<{ activationLink: string }>>(
      `/user/${payload.role}/create`,
      payload
    )
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
