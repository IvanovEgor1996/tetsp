package db

import (
	"context"
	"github.com/go-pg/pg/extra/pgdebug/v10"
	"github.com/go-pg/pg/v10"
	"os"
)

type DBConnector struct {
	DB pg.DBI
}

func NewDBConnector() (DBConnector, error) {

	DB := pg.Connect(&pg.Options{
		Addr:     os.Getenv("PG_HOST") + ":" + os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DBNAME"),
	})
	DB.AddQueryHook(pgdebug.NewDebugHook())
	if err := DB.Ping(context.Background()); err != nil {
		return DBConnector{}, err
	}

	return DBConnector{DB: DB}, nil

}

func (con DBConnector) Ping() error {
	if err := con.DB.(*pg.DB).Ping(context.Background()); err != nil {
		return err
	}
	return nil
}

func (con DBConnector) StartTX() (DBConnector, error) {

	tx, err := con.DB.Begin()
	if err != nil {
		return DBConnector{}, err
	}
	_, err = tx.Exec("Set TRANSACTION ISOLATION LEVEL SERIALIZABLE;")
	if err != nil {
		return DBConnector{}, err
	}

	return DBConnector{DB: tx}, nil
}

func (con DBConnector) CommitTX() error {
	return con.DB.(*pg.Tx).Commit()
}

func (con DBConnector) RollbackTX() error {
	return con.DB.(*pg.Tx).Rollback()
}

func (con DBConnector) CloseTX() error {
	return con.DB.(*pg.Tx).Close()
}
