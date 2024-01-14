import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import type { SchoolYearDetailed } from '@/types/SchoolYear'

export const schoolYearResource = {
  list() {
    return unwrapRequestData(client.get<ApiBaseResponse<SchoolYearDetailed[]>>('/school-year'))
  },
  openSchoolYear(payload: { name: string; start: string; end: string }) {
    return unwrapRequestData(
      client.post<ApiBaseResponse<SchoolYearDetailed[]>>('/school-year', payload)
    )
  }
}
