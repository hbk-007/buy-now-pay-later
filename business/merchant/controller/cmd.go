package controller

import (
	"bufio"
	"fmt"
	"himanshu/poc/bootstrap"
	"himanshu/poc/domain/entity"
	"himanshu/poc/utils"
	"strconv"
)

func RegisterMerchant(reader *bufio.Reader) {
	fmt.Println("Enter the details of the new merchant in the format merchantName, merchantEmail, merchantDiscountOffered")
	vals := utils.ReadVals(reader)
	if len(vals) < 3 {
		fmt.Println("invalid merchant format!")
		return
	}
	merchantDiscount, err := strconv.ParseFloat(vals[2], 64)
	if err != nil {
		fmt.Println("invalid merchantDiscount entered. Please try Again!")
		return
	}
	merchant := entity.Merchant{Name: vals[0], EmailId: vals[1]}
	merchant.Promo.Discount = merchantDiscount
	mid, err := bootstrap.MerchantUCase.RegisterMerchant(merchant)
	if err != nil {
		fmt.Printf("registeration of this merchant failed. Reason ('%s') \n", err.Error())
	}
	fmt.Printf("new merchant created at id=%d\n", mid)
}
