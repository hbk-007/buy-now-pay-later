package main

import (
	"bufio"
	"fmt"
	"himanshu/poc/bootstrap"
	merchantController "himanshu/poc/business/merchant/controller"
	reportsController "himanshu/poc/business/reporting/controller"
	userController "himanshu/poc/business/users/controller"
	"os"
	"strconv"
	"strings"
)

func main() {
	bootstrap.Init()
	Run()
}

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("<=====================================>")
		fmt.Println("enter the action you want to perform")
		fmt.Println("<=====================================>")
		fmt.Println("1. new user")
		fmt.Println("2. new merchant")
		fmt.Println("3. new transaction")
		fmt.Println("4. payback")
		fmt.Println("5. report-total-dues")
		fmt.Println("6. report-discount")
		fmt.Println("7. report users-at-credit-limit")
		fmt.Println("<=====================================>")
		text, _ := reader.ReadString('\n')
		action, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
		if err != nil {
			fmt.Println("something bad happened")
		}
		switch action {
		case 1:
			userController.RegisterUser(reader)
		case 2:
			merchantController.RegisterMerchant(reader)
		case 3:
			userController.MakeTransaction(reader)
		case 4:
			userController.PayBack(reader)
		case 5:
			reportsController.ReportTotalDues(reader)
		case 6:
			reportsController.ReportTotalDiscount(reader)
		case 7:
			reportsController.ReportUsersAtCreditLimit(reader)
		default:
			fmt.Println("no such action exists. Please Try Again!")
		}
	}
}
