package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"time"
)

func main() {
	log.Println(gofakeit.Name())
	log.Println(gofakeit.Email())

	startDate, _ := time.Parse("2006-01-02", "1980-01-01")
	endDate, _ := time.Parse("2006-01-02", "2010-01-01")
	log.Println(gofakeit.DateRange(startDate, endDate))
}
