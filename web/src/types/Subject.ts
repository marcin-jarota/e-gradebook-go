export type Subject = {
  name: string
  id: number
  teachers: { id: number; name: string; surname: string }[]
}

export type SubjectCreatePayload = {
  name: string
}
