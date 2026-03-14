package logic 

import (
	"context"
)

type CreateAccountParams struct{
	Username string 
	Password string 
}

type user struct {
	ID uint64 
	Username string 
}
type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams ) (User, error)

}

type account struct {

}
func NewAccount() Account {
	return &account{}
}

func (a account) CreateAccount(ctx cntext.Context, params CreateAccountParams) (User, error) {
	return User{}, nil
}
