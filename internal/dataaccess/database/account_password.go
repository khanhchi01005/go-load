package database

import "context"

type accountPassword struct {
	OfAccountID uint 64 `sql: "of_account_id`
	Hash string `sql: "hash"`
}

type AccountPasswordDataAccessor interface {
	CreateAccountPassword(ctx context.Context, accountPassword accountPassword) error
	UpdateAccountPassword(ctx context.Context, accountPassword accountPassword) error
}

type accountPasswordDataAccessor struct {
	database *goqu.Database
}

func NewAccountPasswordAcessor(database *goqu.Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{database: database,}
}

func (a *accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword accountPassword) error {
	return nil
}

func (a *accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPassword accountPassword) error {
	return nil
}