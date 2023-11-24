import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import {
  type ClassGroupOutput,
  type ClassGroupPayload,
  type ClassGroupStudent
} from '@/types/ClassGroup'

export const classGroupResource = {
  list() {
    return client.get<ApiBaseResponse<ClassGroupOutput[]>>('/class-groups')
  },
  students(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<ClassGroupStudent[]>>(
        '/class-groups/' + classGroupID + '/students'
      )
    )
  },
  getOne(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<ClassGroupOutput>>('/class-groups/' + classGroupID)
    )
  },
  getMarks(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<{ value: number; id: number }[]>>(
        '/class-groups/' + classGroupID + '/marks'
      )
    )
  },
  async create(payload: ClassGroupPayload) {
    return unwrapRequestData(client.post('/class-groups', payload))
  }
}
