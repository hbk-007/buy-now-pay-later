package entity

type Transaction struct {
	Id         int //transactionId
	Discount   float64
	Amount     float64
	UserId     int
	MerchantId int
}
