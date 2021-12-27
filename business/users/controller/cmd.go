package controller

import (
	"bufio"
	"fmt"
	"himanshu/poc/bootstrap"
	"himanshu/poc/domain/entity"
	"himanshu/poc/utils"
	"strconv"
)

func RegisterUser(reader *bufio.Reader) {
	fmt.Println("Enter the details of the new user in the format userName, userEmail, creditLimit")
	vals := utils.ReadVals(reader)
	if len(vals) < 3 {
		fmt.Println("invalid user format!")
		return
	}
	creditLimit, err := strconv.ParseFloat(vals[2], 64)
	if err != nil {
		fmt.Println("invalid credit Limit entered. Please try Again!")
		return
	}
	user := entity.User{Name: vals[0], Email: vals[1]}
	user.Account.CreditLimit = creditLimit
	err = bootstrap.UserUCase.RegisterUSer(user)
	if err != nil {
		fmt.Printf("registeration of this user failed. Reason ('%s') \n", err.Error())
	}
	fmt.Println("user successfully registered!")
}

func MakeTransaction(reader *bufio.Reader) {
	fmt.Println("Enter the details of the txn in format userId merchantId txnAmount")
	vals := utils.ReadVals(reader)
	if len(vals) < 3 {
		fmt.Println("invalid inputs")
		return
	}
	userId, err := strconv.Atoi(vals[0])
	if err != nil {
		fmt.Println("Invalid User entered!")
		return
	}
	merchantId, err := strconv.Atoi(vals[1])
	if err != nil {
		fmt.Println("Invalid merchant entered!")
		return
	}
	amount, err := strconv.ParseFloat(vals[2], 64)
	if err != nil {
		fmt.Println("Invalid amount!")
		return
	}
	user := entity.User{Id: userId}
	txn := entity.Transaction{Amount: amount, UserId: userId, MerchantId: merchantId}
	err = bootstrap.UserUCase.Transact(user, txn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txn succesfully processed")
}

func PayBack(reader *bufio.Reader) {
	fmt.Println("Enter yourId and amount for paying back the dues")
	vals := utils.ReadVals(reader)
	if len(vals) < 2 {
		fmt.Println("invalid inputs")
		return
	}
	userId, err := strconv.Atoi(vals[0])
	if err != nil {
		fmt.Println("invalid userId entered. Please Try Again!")
		return
	}
	amount, err := strconv.ParseFloat(vals[1], 64)
	if err != nil {
		fmt.Println("invalid amount entered. Please Try Again!")
		return
	}
	err = bootstrap.UserUCase.PayBack(userId, amount)
	if err != nil {
		fmt.Printf("credit amount pay failed. Reason (%s)\n", err.Error())
		return
	}
	fmt.Println("credit pay successfull")
}
