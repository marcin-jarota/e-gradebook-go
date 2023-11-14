import client from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import { type SubjectCreatePayload, type Subject } from '@/types/Subject'

export const subjectResource = {
  list() {
    return client.get<ApiBaseResponse<Subject[]>>('/subject/all')
  },
  async create(payload: SubjectCreatePayload) {
    return handleRequest(client.post('/subject/create', payload))
  }
}

const handleRequest = async <T extends Promise<any>>(t: T): Promise<Awaited<T>> => {
  const r = await t
  if (r?.data?.error) {
    throw new Error(r?.data?.error)
  }

  return r?.data
}
