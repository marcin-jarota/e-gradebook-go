export type AddMarkPayload = {
  subjectID: number
  value: number
  studentID: number
  teacherID: number
  comment?: string
  date?: string
}

export type StudentMark = {
  comment: string
  value: number
  subject: {
    id: number
    name: string
  }
  teacher: { id: number; name: string; surname: string }
}

type MarkMap = {
  1: string
  1.5: string
  2: string
  2.5: string
  3: string
  3.5: string
  4: string
  4.5: string
  5: string
  5.5: string
  6: string
}

export const markSimpleValue: MarkMap = {
  1: '1',
  1.5: '1+',
  2: '2',
  2.5: '2+',
  3: '3',
  3.5: '3+',
  4: '4',
  4.5: '4+',
  5: '5',
  5.5: '5+',
  6: '6'
}

export const markColor: MarkMap = {
  1: 'bg-danger',
  1.5: 'bg-danger',
  2: 'bg-warning',
  2.5: 'bg-warning',
  3: 'bg-warning',
  3.5: 'bg-primary',
  4: 'bg-primary',
  4.5: 'bg-pirmary',
  5: 'bg-primary',
  5.5: 'bg-primary',
  6: 'bg-success'
}
