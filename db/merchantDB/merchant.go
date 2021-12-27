package merchantdb

import (
	"container/list"
	"himanshu/poc/domain/entity"
	"sync"
)

var merchants = list.New()

var merchantCounter = 0

var addmutex sync.Mutex

var merchantIdMap = make(map[int]*entity.Merchant)

func AddMerchant(merchant entity.Merchant) int {
	addmutex.Lock()
	defer addmutex.Unlock()
	merchant.Id = merchantCounter
	merchants.PushBack(merchant)
	e := merchants.Back().Value.(entity.Merchant)
	merchantIdMap[merchant.Id] = &e
	merchantCounter += 1
	return merchant.Id
}

func DeleteMerchant(merchant entity.Merchant) {
	merchantIdMap[merchant.Id].IsActive = false
}

func UpdateMerchant(merchant entity.Merchant) {
	if merchant.Name != "" {
		//name cannot be empty
		merchantIdMap[merchant.Id].Name = merchant.Name
	}
	merchantIdMap[merchant.Id].Promo.Discount = merchant.Promo.Discount
}

func GetMerchantById(id int) *entity.Merchant {
	return merchantIdMap[id]
}
