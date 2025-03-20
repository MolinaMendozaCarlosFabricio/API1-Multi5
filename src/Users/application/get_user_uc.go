package application

import "api1-multi.com/a/src/Users/domain"

type ViewUserUC struct {
	db domain.IUserRepo
}

func NewViewUserUC(db domain.IUserRepo)*ViewUserUC{
	return&ViewUserUC{db: db}
}

func(uc *ViewUserUC)Execute(id int)([]domain.User, error){
	return uc.db.ViewUserMethod(id)
}