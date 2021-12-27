package controller

import (
	"bufio"
	"fmt"
	"himanshu/poc/bootstrap"
	"himanshu/poc/utils"
	"strconv"
)

func ReportTotalDues(reader *bufio.Reader) {
	users, err := bootstrap.ReportingUCase.GetTotalDuesByUser()
	if err != nil {
		fmt.Printf("fetching users' dues failed. Reason is (%s)\n", err.Error())
		return
	}
	if len(users) == 0 {
		fmt.Println("no users with pending dues exist.")
		return
	}
	for _, user := range users {
		fmt.Printf("userid=%d has pending dues = %f\n", user.Id, user.CreditUsed)
	}
}

func ReportTotalDiscount(reader *bufio.Reader) {
	fmt.Println("Enter the merchantId for which you want to calculate Discount")
	vals := utils.ReadVals(reader)
	if len(vals) < 1 {
		fmt.Println("invalid inputs")
		return
	}
	merchantId, err := strconv.Atoi(vals[0])
	if err != nil {
		fmt.Println("invalid merchantId entered. Please Try Again!")
		return
	}
	discount, err := bootstrap.ReportingUCase.GetTotalDiscountByMerchant(merchantId)
	if err != nil {
		fmt.Printf("report fetching failed. Reason is (%s)\n", err.Error())
		return
	}
	fmt.Printf("Total Discount at merchantId %d is %f\n", merchantId, discount)
}

func ReportUsersAtCreditLimit(reader *bufio.Reader) {
	users, err := bootstrap.ReportingUCase.GetUsersAtCreditLimit()
	if err != nil {
		fmt.Printf("fetching users at credit limit failed. Reason is (%s)\n", err.Error())
	}
	fmt.Println("users at credit limit are: ", users)
}
