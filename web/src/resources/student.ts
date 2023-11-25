import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import type { StudentResponse } from '@/types/Student'

export const studentResource = {
  getAll() {
    return unwrapRequestData(client.get<ApiBaseResponse<StudentResponse[]>>('/students'))
  }
}
