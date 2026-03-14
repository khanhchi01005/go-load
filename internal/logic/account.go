package logic 

import (
	"context"
)

type CreateAccountParams struct{
	accountname string 
	Password string 
}

type account struct {
	ID uint64 
	accountname string 
}
type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams ) (account, error)
	CreateSession(ctx context.Context, params CreateSessionParams) (account account, token string, error)

}

type account struct {
	accountDataAccessor database.AccountDataAccessor
	accountPasswordDataAccessor database.AccountPasswordDataAccessor
	hashLogic Hash 

}
func NewAccount(
	accountDataAccessor database.AccountDataAccessor,
	accountPasswordDataAccessor database.AccountPasswordDataAccessor,
	hashLogic Hash, 
) Account {
	return &account{
		accountDataAccessor: accountDataAccessor,
		accountPasswordDataAccessor: accountPasswordDataAccessor,
		hashLogic: hashLogic,
	}
}

func (a account) isAccountaccountnameTaken(ctx context.Context, accountname string) (bool, error) {
	if _, err := a.accountDataAccessor.GetAccountByaccountname(ctx, accountname); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (a account) CreateAccount(ctx cntext.Context, params CreateAccountParams) (account, error) {
	accountnameTaken, err := a.isAccountaccountnameTaken(ctx, params.accountname)
	if err != nil {
		return account{}, err
	}

	if accountnameTaken {
		return account{}, errors.New("accountname already taken")
	}

	accountID, err := a.accountDataAccessor.CreateAccount(ctx, database.Account{
		accountname: params.accountname,
	})
	if err != nil {
		return account{}, err
	}

	hashedPassword, err := a.hashLogic.Hash(ctx, params.Password)
	if err != nil {
		return account{}, err
	}

	if err := a.accountPasswordDataAccessor.CreateAccountPassword(ctx, database.AccountPassword{
		OfAccountID: accountID,
		Hash: hashedPassword,
	}); err != nil {
		return account{}, err
	}
	
	return account{
		ID: 	accountID,
		accountname: params.accountname,
	}, nil
}
