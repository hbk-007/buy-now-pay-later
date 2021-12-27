package repository

import (
	"errors"
	merchantdb "himanshu/poc/db/merchantDB"
	"himanshu/poc/domain/entity"
	"himanshu/poc/domain/interfaces"
)

type merchantRepo struct {
}

func NewMerchantRepository() interfaces.IMerchantRepo {
	return &merchantRepo{}
}

func (mrepo *merchantRepo) AddMerchant(merchant entity.Merchant) (int, error) {
	merchantId := merchantdb.AddMerchant(merchant)
	if merchantId == 0 {
		return 0, errors.New("merchant couldn't be added")
	}
	return merchantId, nil
}

func (mrepo *merchantRepo) DeleteMerchant(merchant entity.Merchant) error {
	merchantdb.DeleteMerchant(merchant)
	return nil
}

func (mrepo *merchantRepo) UpdateMerchant(merchant entity.Merchant) error {
	merchantdb.UpdateMerchant(merchant)
	return nil
}

func (mrepo *merchantRepo) GetMerchantById(mid int) (entity.Merchant, error) {
	merchant := merchantdb.GetMerchantById(mid)
	if merchant == nil {
		return entity.Merchant{}, errors.New("no merchant found")
	}
	return *merchant, nil
}
