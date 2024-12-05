package main

import (
	"github.com/golkity/calc_go/internal/applicant"
	"time"
)

func main() {
	app := application.New(time.Second * 10)

	app.Run()
}
