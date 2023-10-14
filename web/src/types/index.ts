export enum Role {
  Student = 'student',
  Admin = 'admin'
}

export type SessionUser = {
  id: number
  name: string
  surname: string
  email: string
  role: Role
}
