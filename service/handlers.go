package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Service) incBalanceHandler(c echo.Context) error {

	var params IncRequest
	if err := c.Bind(&params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	tx, err := s.DB.StartTX()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	balance, err := tx.GetBalance(params.User)
	if err != nil {
		tx.CloseTX()
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if balance < int(-params.Sum) {
		tx.CloseTX()
		return c.String(http.StatusBadRequest, "balance cant be < 0")
	}
	if err := tx.IncBalance(params.User, int(params.Sum)); err != nil {
		tx.CloseTX()
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := tx.CommitTX(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)

}

func (s *Service) moveBalanceHandler(c echo.Context) error {

	var params MoveRequest
	if err := c.Bind(&params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	tx, err := s.DB.StartTX()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	balanceFrom, err := tx.GetBalance(params.UserFrom)
	if err != nil {
		tx.RollbackTX()
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if balanceFrom < int(params.Sum) {
		tx.RollbackTX()
		return c.String(http.StatusBadRequest, "not enough money on balance")
	}

	if _, err := tx.GetBalance(params.UserTo); err != nil {
		tx.RollbackTX()
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := tx.IncBalance(params.UserFrom, int(-params.Sum)); err != nil {
		tx.RollbackTX()
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err := tx.IncBalance(params.UserTo, int(params.Sum)); err != nil {
		tx.RollbackTX()
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err := tx.CommitTX(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)

}
