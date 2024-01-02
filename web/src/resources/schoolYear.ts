import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import type { SchoolYearDetailed } from '@/types/SchoolYear'

export const schoolYearResource = {
  list() {
    return unwrapRequestData(client.get<ApiBaseResponse<SchoolYearDetailed[]>>('/school-year'))
  }
}
