import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import { type SubjectCreatePayload, type Subject } from '@/types/Subject'

export const subjectResource = {
  list() {
    return unwrapRequestData(client.get<ApiBaseResponse<Subject[]>>('/subjects'))
  },
  assignTeacher({ teacherID, subjectID }: { teacherID: number; subjectID: number }) {
    return unwrapRequestData(
      client.post<ApiBaseResponse<{ success: boolean }>>(`/subjects/${subjectID}/teachers`, {
        teacherID
      })
    )
  },
  async create(payload: SubjectCreatePayload) {
    return unwrapRequestData(client.post('/subjects', payload))
  },
  async delete(id: number) {
    return unwrapRequestData(client.delete('/subjects/' + id))
  }
}
