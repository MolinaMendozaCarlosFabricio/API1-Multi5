package application

import "api1-multi.com/a/src/Users/domain"

type EditUserUC struct {
	db domain.IUserRepo
}

func NewEditUserUC(db domain.IUserRepo)*EditUserUC{
	return&EditUserUC{db: db}
}

func(uc *EditUserUC)Execute(id int, first_name, last_name string)error{
	user := &domain.User{ID: id, First_name: first_name, Last_name: last_name}
	return uc.db.EditUserMethod(*user)
}