import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse, UserOutput, UserInput, SetupPasswordPayload } from '@/types'

export const userResource = {
  async list() {
    return unwrapRequestData(client.get<ApiBaseResponse<UserOutput[]>>('/user/list'))
  },
  async tokenValid(token: string) {
    return client.get('/token-valid', {
      headers: {
        Authorization: 'Bearer ' + token
      }
    })
  },
  async setupPassword(payload: SetupPasswordPayload, token: string) {
    return unwrapRequestData(client.post(`/setup-password?token=${token}`, payload))
  },
  async create(payload: UserInput) {
    const { role, ...restPayload } = payload
    return unwrapRequestData(
      client.post<ApiBaseResponse<{ activationLink: string }>>(`/user/create/${role}`, restPayload)
    )
  },
  async activate(userID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<{ ok: boolean }>>(`/user/activate/${userID}`)
    )
  },
  async deactivate(userID: number) {
    return unwrapRequestData(client.get<ApiBaseResponse>(`/user/deactivate/${userID}`))
  },
  async destroySession(userID: number) {
    return unwrapRequestData(client.get<ApiBaseResponse<any>>(`/user/destroy-session/${userID}`))
  }
}
