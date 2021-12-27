package bootstrap

import (
	mrepo "himanshu/poc/business/merchant/repository"
	mucase "himanshu/poc/business/merchant/usecase"
	"himanshu/poc/business/reporting/usecase"
	urepo "himanshu/poc/business/users/repository"
	uucase "himanshu/poc/business/users/usecase"
	"himanshu/poc/domain/interfaces"
)

var MerchantUCase interfaces.IMerchantUCase
var UserUCase interfaces.IUserUCase
var ReportingUCase interfaces.IReportingUCase

func Init() {
	merchantRepo := mrepo.NewMerchantRepository()
	userRepo := urepo.NewUserRepository()
	txnRepo := urepo.NewTransactionRepository()
	MerchantUCase = mucase.NewMerchantUCase(merchantRepo)
	UserUCase = uucase.NewUserUCase(userRepo, txnRepo, merchantRepo)
	ReportingUCase = usecase.NewReportingUCase(txnRepo, userRepo)
}
