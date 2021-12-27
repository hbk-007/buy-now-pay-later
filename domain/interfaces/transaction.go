package interfaces

import "himanshu/poc/domain/entity"

type ITransactionRepo interface {
	CreateTransaction(txn entity.Transaction) error
	GetTotalDiscountByMerchantId(merchantId int) (float64, error)
}
