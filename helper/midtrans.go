package helper

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	uuid "github.com/satori/go.uuid"
)

func MidtransCreateTransaction() interface{} {
	var s = snap.Client{}
	s.New("SB-Mid-server-nP8oOrzwnFwp8UTSeDXEhm7v", midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).
	// 2. Initiate Snap request param
	orderID := uuid.NewV4().String()
	orderID = "GROUP-3-ORDER-ID-" + orderID
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
	}
	snapResp, _ := s.CreateTransaction(req)
	return snapResp
}
