export type ClassGroupOutput = {
  id: number
  name: string
  studentsCount: number
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
