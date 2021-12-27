package interfaces

import "himanshu/poc/domain/entity"

type IReportingUCase interface {
	GetUsersAtCreditLimit() ([]int, error)
	GetTotalDiscountByMerchant(merchantId int) (float64, error)
	GetTotalDuesByUser() ([]entity.User, error)
}
