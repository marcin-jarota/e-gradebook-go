import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import type { TeacherOutput } from '@/types/Teacher'

export const teacherResource = {
  list() {
    return unwrapRequestData(client.get<ApiBaseResponse<TeacherOutput[]>>('/teachers'))
  }
}
