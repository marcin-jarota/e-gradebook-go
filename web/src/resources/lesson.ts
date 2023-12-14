import client, { unwrapRequestData } from '@/lib/axios'

export const lessonResource = {
  create(payload: {
    teacherID: number
    subjectID: number
    classGroupID: number
    startTime: string
    endTime: string
    dayOfWeek: number
  }) {
    return unwrapRequestData(client.post('/lessons', payload))
  }
}
