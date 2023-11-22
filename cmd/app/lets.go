package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type User struct {
	name     string
	age      int
	password string
	y        []int
}

func (u *User) setName(a string) {
	u.name = a
}

func (u User) getName() string {
	return u.name
}

func change(a *string) {
	*a = "asd"
}

func main() {

	s := echo.New()

	s.GET("/status", Handler)
	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

func Handler(ctx echo.Context) error {
	d := time.Date(2025, time.June, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	s := fmt.Sprintf("%d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
