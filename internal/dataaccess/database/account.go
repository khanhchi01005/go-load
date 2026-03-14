pacakge database

type Account struct {
	accountId uint64 `sql:"account_id"`
	accountname string `sql:"accountname"`
}

type AccountDataAccessor interface{
	CreateAccount(ctx context.Context, account Account) (uint64,error)
	GetAccountByID(ctx context.Context, id uint64) (Account, error)
	GetAccountByaccountname(ctx context.Context, accountname string) (Account, error)
}

type accountDataAccessor struct {
	database *goqu.Database
}

func NewAccountAcessor(database *goqu.Database) AccountDataAccessor {
	return &accountDataAccessor{database: database}
}

func (a accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64,error) {
	result, err := a.database.
		Insert("accounts").
		Rows(goqu.Record{
			"accountname": account.accountname,
		}).
		Executor().
		ExeContext(ctx)
	if err != nil {
		log.Printf("Failed to create account", err)
		return 0, err
	}

	lastInsertedId , err :=result.LastInsertId()
	if err != nil {
		log.Printf("Failed to create account", err)
		return 0, err
	}
	return uint64(lastInsertedId), nil
}

func (a *accountDataAccessor) GetAccountByID(ctx context.Context, id uint64) (Account, error) {
	return Account{}, nil
}

func (a *accountDataAccessor) GetAccountByaccountname(ctx context.Context, accountname string) (Account, error) {
	return Account{}, nil
}
