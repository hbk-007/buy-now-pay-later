package repository

import (
	transactiondb "himanshu/poc/db/transactionDB"
	"himanshu/poc/domain/entity"
	"himanshu/poc/domain/interfaces"
)

type transactionRepo struct {
}

func NewTransactionRepository() interfaces.ITransactionRepo {
	return &transactionRepo{}
}

func (txnRepo *transactionRepo) CreateTransaction(txn entity.Transaction) error {
	transactiondb.AddTransaction(txn)
	return nil
}

func (txnRepo *transactionRepo) GetTotalDiscountByMerchantId(merchantId int) (float64, error) {
	totalDiscount := transactiondb.GetTotalDiscountByMerchant(merchantId)
	return totalDiscount, nil
}
