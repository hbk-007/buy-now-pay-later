package interfaces

import "himanshu/poc/domain/entity"

type IMerchantUCase interface {
	RegisterMerchant(merchant entity.Merchant) (int, error)
	DeleteMerchant(merchant entity.Merchant) error
	UpdateMerchant(merchant entity.Merchant) error
}

type IMerchantRepo interface {
	AddMerchant(merchant entity.Merchant) (int, error)
	DeleteMerchant(merchant entity.Merchant) error
	UpdateMerchant(merchant entity.Merchant) error
	GetMerchantById(mid int) (entity.Merchant, error)
}
