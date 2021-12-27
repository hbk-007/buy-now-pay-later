package usecase

import (
	"himanshu/poc/domain/entity"
	"himanshu/poc/domain/interfaces"
)

type merchantUCase struct {
	merchantRepo interfaces.IMerchantRepo
}

func NewMerchantUCase(mrepo interfaces.IMerchantRepo) interfaces.IMerchantUCase {
	return &merchantUCase{merchantRepo: mrepo}
}

func (muc *merchantUCase) RegisterMerchant(merchant entity.Merchant) (int, error) {
	return muc.merchantRepo.AddMerchant(merchant)
}

func (muc *merchantUCase) DeleteMerchant(merchant entity.Merchant) error {
	return muc.merchantRepo.DeleteMerchant(merchant)
}

func (muc *merchantUCase) UpdateMerchant(merchant entity.Merchant) error {
	return muc.merchantRepo.UpdateMerchant(merchant)
}
