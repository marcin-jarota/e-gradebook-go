export type ClassGroupOutput = {
  id: number
  name: string
  studentsCount: number
  educationYear?: number
}

export type TeacherSubject = {
  teacher: {
    id: number
    name: string
    surname: string
    email: string
  }
  subject: {
    id: number
    name: string
  }
}
export type ClassGroupStudent = {
  id: number
  name: string
  surname: string
  email: string
  avgMark: number
}

export type ClassGroupPayload = {
  name: string
}

export type Lesson = {
  id: number
  subject: string
  teacher: string
  start: string
  end: string
  dayOfWeek: number
}
