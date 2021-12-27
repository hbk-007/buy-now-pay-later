package repository

import (
	"errors"
	userdb "himanshu/poc/db/userDB"
	"himanshu/poc/domain/entity"
	"himanshu/poc/domain/interfaces"
)

type userRepo struct {
}

func NewUserRepository() interfaces.IUserRepo {
	return &userRepo{}
}

func (urepo *userRepo) AddUser(user entity.User) error {
	userdb.AddUser(user)
	return nil
}

func (urepo *userRepo) DeleteUser(user entity.User) error {
	userdb.DeleteUser(user)
	return nil
}
func (urepo *userRepo) UpdateUser(user entity.User) error {
	return userdb.UpdateUser(user)
}

func (urepo *userRepo) GetUserById(id int) (entity.User, error) {
	user := userdb.GetUserById(id)
	if user == nil {
		return entity.User{}, errors.New("no such user found")
	}
	return *user, nil
}

func (urepo *userRepo) GetAllUsersAtCreditLimit() ([]int, error) {
	ListUsersIds := userdb.GetUserAtCreditLimit()
	return ListUsersIds, nil
}

func (urepo *userRepo) GetUsersWithDues() ([]entity.User, error) {
	users, err := userdb.GetUsersWithDues()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (urepo *userRepo) Payback(userId int, amount float64) error {
	err := userdb.UpdateUserCrediUsed(userId, amount)
	if err != nil {
		return err
	}
	return nil
}
