package application

import "api1-multi.com/a/src/Users/domain"

type DeleteUserUC struct {
	db domain.IUserRepo
}

func NewDeleteUserUC(db domain.IUserRepo)*DeleteUserUC{
	return&DeleteUserUC{db: db}
}

func(uc *DeleteUserUC)Execute(id int)error{
	return uc.db.DeleteUserMethod(id)
}