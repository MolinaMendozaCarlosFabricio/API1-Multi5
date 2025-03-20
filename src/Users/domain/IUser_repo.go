package domain

type IUserRepo interface{
	RegisterUserMethod(user User, Credentials AccessCredentials)error
	ViewUserMethod(id int)([]User, error)
	EditUserMethod(user User)error
	EditCredentialsMethod(credentials AccessCredentials)error
	DeleteUserMethod(id int)error
}