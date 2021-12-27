package transactiondb

import (
	"container/list"
	"himanshu/poc/domain/entity"
	"sync"
)

var addTxnMutex sync.Mutex

var ListTransactions = list.New()

var transactionCounter = 0

var merchantTransactionsMap = make(map[int][]*entity.Transaction)

var userTransactionMap = make(map[int][]*entity.Transaction)

func AddTransaction(txn entity.Transaction) {
	addTxnMutex.Lock()
	defer addTxnMutex.Unlock()
	txn.Id = transactionCounter
	e := ListTransactions.PushBack(txn)
	v := (e.Value.(entity.Transaction))
	merchantTransactionsMap[txn.MerchantId] = append(merchantTransactionsMap[txn.MerchantId], &v)
	userTransactionMap[txn.UserId] = append(userTransactionMap[txn.UserId], &v)
	transactionCounter += 1
}

func GetAllTransactions() []entity.Transaction {
	var tmp []entity.Transaction
	for e := ListTransactions.Front(); e != nil; e = e.Next() {
		tmp = append(tmp, e.Value.(entity.Transaction))
	}
	return tmp
}

func GetTotalDiscountByMerchant(merchantId int) float64 {
	var totalDiscount float64
	for _, txn := range merchantTransactionsMap[merchantId] {
		totalDiscount += txn.Discount
	}
	return totalDiscount
}
