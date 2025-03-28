package application

import "api1-multi.com/a/src/Users/domain"

type RegisterUserUC struct {
	db domain.IUserRepo
}

func NewRegisterUserUC(db domain.IUserRepo)*RegisterUserUC{
	return&RegisterUserUC{db: db}
}

func(uc *RegisterUserUC)Execute(first_name, last_name, email, password string)error{
	user := &domain.User{ID: 0, First_name: first_name, Last_name: last_name}
	credentials := &domain.AccessCredentials{Id_user: 0, Email: email, Password: password}
	return uc.db.RegisterUserMethod(*user, *credentials)
}