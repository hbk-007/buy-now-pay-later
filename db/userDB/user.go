package userdb

import (
	"container/list"
	"errors"
	"himanshu/poc/domain/entity"
	"sync"
)

var users = list.New()

var userCounter = 0

var addUserMutex sync.Mutex

var userIdMap = make(map[int]*entity.User)

func AddUser(user entity.User) {
	addUserMutex.Lock()
	defer addUserMutex.Unlock()
	user.Id = userCounter
	users.PushBack(user)
	u := users.Back().Value.(entity.User)
	userIdMap[user.Id] = &u
	userCounter += 1
}

func DeleteUser(user entity.User) {
	user.IsActive = false
}

func UpdateUser(user entity.User) error {
	if userIdMap[user.Id] == nil {
		return errors.New("no such user exist")
	}
	userIdMap[user.Id].CreditLimit = user.CreditLimit
	userIdMap[user.Id].Name = user.Name
	userIdMap[user.Id].Email = user.Email
	userIdMap[user.Id].CreditUsed = user.CreditUsed
	return nil
}

func GetUserById(id int) *entity.User {
	return userIdMap[id]
}

func GetUserAtCreditLimit() []int {
	var ListUserAtCreditLimit []int
	for _, user := range userIdMap {
		if user.CreditLimit == user.CreditUsed {
			ListUserAtCreditLimit = append(ListUserAtCreditLimit, user.Id)
		}
	}
	return ListUserAtCreditLimit
}

func GetUsersWithDues() ([]entity.User, error) {
	var ListUsers []entity.User
	for _, user := range userIdMap {
		if user.CreditUsed > 0 {
			ListUsers = append(ListUsers, *user)
		}
	}
	return ListUsers, nil
}

func UpdateUserCrediUsed(userId int, amount float64) error {
	if userIdMap[userId] == nil {
		return errors.New("no such user exists")
	}
	if userIdMap[userId].CreditUsed < amount {
		// the following constraint can be added to the sql constraint or equivalent at application layer
		return errors.New("repay amount cannot be greater than the credit limit")
	}
	userIdMap[userId].CreditUsed -= amount
	return nil
}
