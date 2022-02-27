package db

import (
	"fmt"
	"github.com/google/uuid"
)

func (con DBConnector) GetBalance(userUUID uuid.UUID) (int, error) {
	balance := Balances{User: userUUID}
	_, err := con.DB.Model(&balance).
		OnConflict("DO NOTHING"). // optional
		WherePK().
		SelectOrInsert()
	if err != nil {
		return 0, err
	}
	return balance.Balance, nil
}

func (con DBConnector) IncBalance(userUUID uuid.UUID, inc int) error {
	res, err := con.DB.Model((*Balances)(nil)).
		Set("balance = balance + ?", inc).
		Where("\"user\" = ?", userUUID).
		Update()
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}
