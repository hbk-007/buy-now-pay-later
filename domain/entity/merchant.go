package entity

type Merchant struct {
	Name     string
	Id       int
	EmailId  string
	Promo    Promotion
	IsActive bool
}

type Promotion struct {
	Discount float64
}
