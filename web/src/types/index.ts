export type ApiBaseResponse<T = any> = {
  error: string | null
  data: T
}

export enum Role {
  Student = 'student',
  Admin = 'admin',
  Teacher = 'teacher'
}

export type SessionUser = {
  id: number
  name: string
  surname: string
  email: string
  role: Role
}

export type UserOutput = {
  id: number
  name: string
  surname: string
  email: string
  role: Role
  isActive?: boolean
  sessionActive?: boolean
}

export type UserInput = {
  name: string
  surname: string
  email: string
  role: Role
}

export type SetupPasswordPayload = {
  password: string
  passwordConfirm: string
}

export type UserListResponse = ApiBaseResponse<UserOutput[]>
