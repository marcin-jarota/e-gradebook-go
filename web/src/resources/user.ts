import client from '@/lib/axios'
import type { ApiBaseResponse, UserListResponse, UserInput, SetupPasswordPayload } from '@/types'

export const userResource = {
  async list() {
    return handleRequest(client.get<UserListResponse>('/user/list'))
  },
  async tokenValid(token: string) {
    return client.get('/token-valid', {
      headers: {
        Authorization: 'Bearer ' + token
      }
    })
  },
  async setupPassword(payload: SetupPasswordPayload, token: string) {
    return handleRequest(client.post(`/setup-password?token=${token}`, payload))
  },
  async create(payload: UserInput) {
    return handleRequest(
      client.post<ApiBaseResponse<{ activationLink: string }>>(
        `/user/${payload.role}/create`,
        payload
      )
    )
  },
  async activate(userID: number) {
    return handleRequest(client.get<ApiBaseResponse<{ ok: boolean }>>(`/user/activate/${userID}`))
  },
  async deactivate(userID: number) {
    return handleRequest(client.get<ApiBaseResponse>(`/user/deactivate/${userID}`))
  },
  async destroySession(userID: number) {
    return handleRequest(client.get<ApiBaseResponse<any>>(`/user/destroy-session/${userID}`))
  }
}

const handleRequest = async <T extends Promise<any>>(t: T): Promise<Awaited<T>> => {
  const r = await t
  if (r?.data?.error) {
    throw new Error(r?.data?.error)
  }

  return r?.data
}
