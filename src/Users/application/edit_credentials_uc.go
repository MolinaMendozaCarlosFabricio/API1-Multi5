package application

import "api1-multi.com/a/src/Users/domain"

type EditCredentialsUC struct {
	db domain.IUserRepo
}

func NewEditCredentialsUC(db domain.IUserRepo)*EditCredentialsUC{
	return&EditCredentialsUC{db: db}
}

func(uc *EditCredentialsUC)Execute(id int, email, password string)error{
	credentials := &domain.AccessCredentials{Id_user: id, Email: email, Password: password}
	return uc.db.EditCredentialsMethod(*credentials)
}