import client, { unwrapRequestData } from '@/lib/axios'
import type { ApiBaseResponse } from '@/types'
import {
  type ClassGroupOutput,
  type ClassGroupPayload,
  type ClassGroupStudent,
  type Lesson,
  type TeacherSubject
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
  teachers(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<ClassGroupStudent[]>>(
        '/class-groups/' + classGroupID + '/teachers'
      )
    )
  },
  lessons(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<Lesson[]>>('/class-groups/' + classGroupID + '/lessons')
    )
  },

  teachersSubjects(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<TeacherSubject[]>>(
        '/class-groups/' + classGroupID + '/teacher-subject'
      )
    )
  },
  subjects(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<ClassGroupStudent[]>>(
        '/class-groups/' + classGroupID + '/subjects'
      )
    )
  },
  getOne(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<ClassGroupOutput>>('/class-groups/' + classGroupID)
    )
  },
  delete(classGroupID: number) {
    return unwrapRequestData(
      client.delete<ApiBaseResponse<{ ok: true }>>('/class-groups/' + classGroupID)
    )
  },
  getMarks(classGroupID: number) {
    return unwrapRequestData(
      client.get<ApiBaseResponse<{ value: number; id: number }[]>>(
        '/class-groups/' + classGroupID + '/marks'
      )
    )
  },
  assignStudent({ studentID, classGroupID }: { studentID: number; classGroupID: number }) {
    return unwrapRequestData(
      client.post<ApiBaseResponse<{ success: boolean }>>(`/class-groups/${classGroupID}/students`, {
        studentID
      })
    )
  },
  assignTeacher({ teacherID, classGroupID }: { teacherID: number; classGroupID: number }) {
    return unwrapRequestData(
      client.post<ApiBaseResponse<{ success: boolean }>>(`/class-groups/${classGroupID}/teachers`, {
        teacherID
      })
    )
  },

  assignSubject({ subjectID, classGroupID }: { subjectID: number; classGroupID: number }) {
    return unwrapRequestData(
      client.post<ApiBaseResponse<{ success: boolean }>>(`/class-groups/${classGroupID}/subjects`, {
        subjectID
      })
    )
  },
  assignTeacherSubject({
    subjectID,
    teacherID,
    classGroupID
  }: {
    teacherID: number
    subjectID: number
    classGroupID: number
  }) {
    return unwrapRequestData(
      client.post<ApiBaseResponse<{ success: boolean }>>(`/class-groups/${classGroupID}/subjects`, {
        subjectID,
        teacherID
      })
    )
  },

  async create(payload: ClassGroupPayload) {
    return unwrapRequestData(client.post('/class-groups', payload))
  }
}
