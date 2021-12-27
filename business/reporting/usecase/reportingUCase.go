package usecase

import (
	"himanshu/poc/domain/entity"
	"himanshu/poc/domain/interfaces"
)

type reportingUCase struct {
	transactionRepo interfaces.ITransactionRepo
	userRepo        interfaces.IUserRepo
}

func NewReportingUCase(txnRepo interfaces.ITransactionRepo, urepo interfaces.IUserRepo) interfaces.IReportingUCase {
	return &reportingUCase{transactionRepo: txnRepo, userRepo: urepo}
}

func (ruc *reportingUCase) GetUsersAtCreditLimit() ([]int, error) {
	return ruc.userRepo.GetAllUsersAtCreditLimit()
}

func (ruc *reportingUCase) GetTotalDiscountByMerchant(merchantId int) (float64, error) {
	return ruc.transactionRepo.GetTotalDiscountByMerchantId(merchantId)
}

func (ruc *reportingUCase) GetTotalDuesByUser() ([]entity.User, error) {
	users, err := ruc.userRepo.GetUsersWithDues()
	if err != nil {
		return nil, err
	}
	return users, nil
}
