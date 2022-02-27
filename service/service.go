package service

import (
	"github.com/labstack/echo/v4"
	"os"
	"testP/db"
)

type Service struct {
	Echo *echo.Echo
	DB   db.DBConnector
}

func NewService() (*Service, error) {

	s := &Service{
		Echo: echo.New(),
	}
	conn, err := db.NewDBConnector()
	if err != nil {
		return nil, err
	}
	s.DB = conn

	s.addHandlers()

	return s, nil
}

func (s *Service) addHandlers() {

	s.Echo.POST("/incBalance", s.incBalanceHandler)
	s.Echo.POST("/moveBalance", s.moveBalanceHandler)

}

func (s *Service) Run() {
	s.Echo.Logger.Fatal(s.Echo.Start(":" + os.Getenv("APP_PORT")))
}
