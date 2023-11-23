import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import { type SubjectCreatePayload, type Subject } from '@/types/Subject'

export const subjectResource = {
  list() {
    return client.get<ApiBaseResponse<Subject[]>>('/subject/all')
  },
  async create(payload: SubjectCreatePayload) {
    return unwrapRequestData(client.post('/subject/create', payload))
  },
  async delete(id: number) {
    return unwrapRequestData(client.get('/subject/delete/' + id))
  }
}
