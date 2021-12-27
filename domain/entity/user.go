package entity

type User struct {
	Name     string
	Id       int
	IsActive bool
	Email    string
	Account
}

type Account struct {
	CreditLimit float64
	CreditUsed  float64 // this will also be the amount left for the user to pay
}
