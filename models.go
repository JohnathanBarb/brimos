package main

import (
	"github.com/shopspring/decimal"
	"time"
)

type Person struct {
	ID        int
	Name      string
	Document  string
	CreatedAt time.Time
}

type Book struct {
	ID        int
	PersonID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Loan struct {
	ID                int
	BookID            int
	Principal         decimal.Decimal
	DailyInterestRate decimal.Decimal
	GrantedAt         time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
