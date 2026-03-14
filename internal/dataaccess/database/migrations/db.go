package database 

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/doug-martin/goqu/v9"
	"go-load/internal/configs"
	_ "github.com/go-sql-driver/mysql"   
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

func InitializeDB(databaseConfig configs.Database)(db *sql.DB, cleanup func(), error)  {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
	databaseConfig.accountname, 
	databaseConfig.Password, 
	databaseConfig.Host, 
	databaseConfig.Port, 
	databaseConfig.Database)

	db, err := sql.Operen("mysql", connectionString)
	if err != nil {
		log.Printf("Cannot connect to database", err)
		return nil, nil, err
	}

	cleaup := func() {
		db.Close()
	}
	return db, cleanup, nil
	
} 

func InitializeGoquDB(db *sql.DB) *goqu.Database {
	return goqu.New("mysql", db)
	
}