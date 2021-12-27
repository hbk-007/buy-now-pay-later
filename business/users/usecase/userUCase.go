package usecase

import (
	"errors"
	"himanshu/poc/domain/entity"
	"himanshu/poc/domain/interfaces"
)

type userUCase struct {
	userRepo        interfaces.IUserRepo
	merchantRepo    interfaces.IMerchantRepo
	transactionRepo interfaces.ITransactionRepo
}

func NewUserUCase(urepo interfaces.IUserRepo, txnRepo interfaces.ITransactionRepo, mRepo interfaces.IMerchantRepo) interfaces.IUserUCase {
	return &userUCase{userRepo: urepo, merchantRepo: mRepo, transactionRepo: txnRepo}
}

func (uucase *userUCase) RegisterUSer(user entity.User) error {
	return uucase.userRepo.AddUser(user)
}

func (uucase *userUCase) DeleteUser(user entity.User) error {
	return uucase.userRepo.DeleteUser(user)
}

func (uucase *userUCase) UpdateUser(user entity.User) error {
	return uucase.userRepo.UpdateUser(user)
}

func (uucase *userUCase) Transact(user entity.User, txn entity.Transaction) error {
	merchantDetails, err := uucase.merchantRepo.GetMerchantById(txn.MerchantId)
	if err != nil {
		return err
	}
	userDetails, err := uucase.userRepo.GetUserById(user.Id)
	if err != nil {
		return err
	}
	if userDetails.CreditLimit-user.CreditUsed == 0 {
		return errors.New("txn rejected! (reason: credit limit reached)")
	}
	if userDetails.CreditLimit-userDetails.CreditUsed < txn.Amount {
		return errors.New("txn rejected! (reason: not sufficient credit)")
	}
	// need to add failsafe mechanism
	txn.Discount = (float64(txn.Amount) * merchantDetails.Promo.Discount * 100) / 10000.00
	err = uucase.transactionRepo.CreateTransaction(txn)
	if err != nil {
		return err
	}
	userDetails.CreditUsed = userDetails.CreditUsed + txn.Amount
	err = uucase.userRepo.UpdateUser(userDetails)
	if err != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (uucase *userUCase) PayBack(userId int, amount float64) error {
	err := uucase.userRepo.Payback(userId, amount)
	if err != nil {
		return err
	}
	return nil
}
