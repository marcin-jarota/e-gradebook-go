import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import { type AddMarkPayload } from '@/types/Mark'

export const markResource = {
  async addMark(payload: AddMarkPayload) {
    return unwrapRequestData(client.post<ApiBaseResponse<{ success: boolean }>>(`/marks`, payload))
  }
}
