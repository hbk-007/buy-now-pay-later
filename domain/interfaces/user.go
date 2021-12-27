package interfaces

import "himanshu/poc/domain/entity"

type IUserUCase interface {
	RegisterUSer(user entity.User) error
	DeleteUser(user entity.User) error
	UpdateUser(user entity.User) error
	Transact(user entity.User, txn entity.Transaction) error
	PayBack(userId int, amount float64) error
}

type IUserRepo interface {
	AddUser(user entity.User) error
	DeleteUser(user entity.User) error
	UpdateUser(user entity.User) error
	GetUserById(id int) (entity.User, error)
	GetAllUsersAtCreditLimit() ([]int, error)
	GetUsersWithDues() ([]entity.User, error)
	Payback(userId int, amount float64) error
}
