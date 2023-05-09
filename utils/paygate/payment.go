package paygate

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	mid "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type payService struct {
	snapData snap.Client
	midApi   coreapi.Client
}

func initializeSnapData(serverkey string) snap.Client {
	var snp snap.Client
	snp.New(serverkey, mid.Sandbox)
	return snp
}

func initializeCoreData(serverkey string) coreapi.Client {
	var coreData coreapi.Client

	coreData.New(serverkey, mid.Sandbox)
	return coreData
}

func NewPaymentService(serverkey string) PayService {
	return &payService{
		snapData: initializeSnapData(serverkey),
		midApi:   initializeCoreData(serverkey),
	}
}

func (s *payService) GeneratePayReq(ctx context.Context, name string, email string, orderid string, total int64) (string, string, error) {

	s.snapData.Options.SetPaymentAppendNotification("http://127.0.0.1:8080/append")

	// Optional : here is how if you want to set override payment notification for this request
	s.snapData.Options.SetPaymentOverrideNotification("http://127.0.0.1:8080/override")

	snapReq := &snap.Request{
		CustomerDetail: &mid.CustomerDetails{
			Email: email,
			FName: name,
		},
		TransactionDetails: mid.TransactionDetails{
			OrderID:  orderid,
			GrossAmt: total,
		},
	}

	resp, err := s.snapData.CreateTransaction(snapReq)
	// Initiate Snap Request
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
		return "failed", "", err
	}

	return resp.Token, resp.RedirectURL, nil
}

func (s *payService) GetNotification(ctx echo.Context, orderid string) string {
	// 1. Initialize empty map
	var notificationPayload map[string]interface{}

	// 2. Parse JSON request body and use it to set json to payload
	//err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	err := json.NewDecoder(ctx.Request().Body).Decode(&notificationPayload)

	if err != nil {
		// do something on error when decode
		return "failed"
	}
	// 3. Get order-id from payload
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return "failed"
	}

	var status string
	//generate core api
	//payment.InitializeCoreClient()

	// 4. Check transaction to Midtrans with param orderId
	//transactionStatusResp, e := getCore().CheckTransaction(orderId)
	transactionStatusResp, e := s.midApi.CheckTransaction(orderId)
	if e != nil {
		//http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		ctx.JSON(http.StatusInternalServerError, e.GetMessage())
		return "failed"
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
					status = "challenge"
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
					status = "success"
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
				status = "settlement"
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
				status = "deny"
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
				status = "cancel"
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
				status = "pending"
			}
		}
	}

	return status
}
