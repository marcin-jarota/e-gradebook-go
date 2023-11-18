import client from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import { type ClassGroupOutput, type ClassGroupPayload } from '@/types/ClassGroup'

export const classGroupResource = {
  list() {
    return client.get<ApiBaseResponse<ClassGroupOutput[]>>('/class/all')
  },
  async create(payload: ClassGroupPayload) {
    return handleRequest(client.post('/class/create', payload))
  }
  // async delete(id: number) {
  //   return handleRequest(client.get('/subject/delete/' + id))
  // }
}

const handleRequest = async <T extends Promise<any>>(t: T): Promise<Awaited<T>> => {
  const r = await t
  if (r?.data?.error) {
    throw new Error(r?.data?.error)
  }

  return r?.data
}
