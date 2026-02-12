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
	UpdatedAt time.Time
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
	Iof               decimal.Decimal
	PlatformFee       decimal.Decimal
	DailyInterestRate decimal.Decimal
	FineRate          decimal.Decimal
	GrantedAt         time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Instalment struct {
	ID           int
	LoanID       int
	InvoiceID    int
	Principal    decimal.Decimal
	Interest     decimal.Decimal
	MoraInterest decimal.Decimal
	LateInterest decimal.Decimal
	Fine         decimal.Decimal
}

type Invoice struct {
	ID           int
	BookID       int
	InvoiceID    int
	Principal    decimal.Decimal
	Interest     decimal.Decimal
	MoraInterest decimal.Decimal
	LateInterest decimal.Decimal
	Fine         decimal.Decimal
}
