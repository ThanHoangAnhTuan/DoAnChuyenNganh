package impl

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/crypto"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/ip"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/payment"
	"go.uber.org/zap"
)

type PaymentImpl struct {
	sqlc *database.Queries
	db   *sql.DB
}

// TODO: admin used to refund user
func (p *PaymentImpl) PostRefund(ctx *gin.Context, in *vo.PostRefundInput) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)

	vnpRequestId := now.Format("150405")
	vnpCreateDate := now.Format("20060102150405")
	ipAddr := ip.GetClientIP(ctx)
	vnpOrderInfo := "Hoan tien GD ma:" + in.OrderId
	vnpTransactionNo := "0"

	// Create data string for signature
	data := strings.Join([]string{
		vnpRequestId,
		"2.1.0",
		"refund",
		global.Config.Payment.VnpTmnCode,
		in.TransType,
		in.OrderId,
		strconv.Itoa(in.Amount * 100),
		vnpTransactionNo,
		in.TransDate,
		in.User,
		vnpCreateDate,
		ipAddr,
		vnpOrderInfo,
	}, "|")

	vnpSecureHash := crypto.CreateHMAC(data, global.Config.Payment.VnpHashSecret)

	dataObj := vo.RefundDataObj{
		VnpRequestId:       vnpRequestId,
		VnpVersion:         "2.1.0",
		VnpCommand:         "refund",
		VnpTmnCode:         global.Config.Payment.VnpTmnCode,
		VnpTransactionType: in.TransType,
		VnpTxnRef:          in.OrderId,
		VnpAmount:          in.Amount * 100,
		VnpTransactionNo:   vnpTransactionNo,
		VnpCreateBy:        in.User,
		VnpOrderInfo:       vnpOrderInfo,
		VnpTransactionDate: in.TransDate,
		VnpCreateDate:      vnpCreateDate,
		VnpIpAddr:          ipAddr,
		VnpSecureHash:      vnpSecureHash,
	}

	// Make HTTP request to VNPay API
	response, err := payment.MakeAPIRequest(global.Config.Payment.VnpApi, dataObj)
	if err != nil {
		log.Printf("Error processing refund: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process refund"})
		return
	}

	log.Printf("Refund response: %s", response)
	ctx.JSON(http.StatusOK, gin.H{"message": "Refund request sent successfully"})
}

// TODO: manager used to look up orders
func (p *PaymentImpl) PostQueryDR(ctx *gin.Context, in *vo.PostQueryDRInput) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)

	vnpRequestId := now.Format("150405")
	vnpCreateDate := now.Format("20060102150405")
	ipAddr := ip.GetClientIP(ctx)

	vnpOrderInfo := "Truy van GD ma:" + in.OrderId

	// Create data string for signature
	data := strings.Join([]string{
		vnpRequestId,
		"2.1.0",
		"querydr",
		global.Config.Payment.VnpTmnCode,
		in.OrderId,
		in.TransDate,
		vnpCreateDate,
		ipAddr,
		vnpOrderInfo,
	}, "|")

	vnpSecureHash := crypto.CreateHMAC(data, global.Config.Payment.VnpHashSecret)

	dataObj := vo.QueryDataObj{
		VnpRequestId:       vnpRequestId,
		VnpVersion:         "2.1.0",
		VnpCommand:         "querydr",
		VnpTmnCode:         global.Config.Payment.VnpTmnCode,
		VnpTxnRef:          in.OrderId,
		VnpOrderInfo:       vnpOrderInfo,
		VnpTransactionDate: in.TransDate,
		VnpCreateDate:      vnpCreateDate,
		VnpIpAddr:          ipAddr,
		VnpSecureHash:      vnpSecureHash,
	}

	// Make HTTP request to VNPay API
	response, err := payment.MakeAPIRequest(global.Config.Payment.VnpApi, dataObj)
	if err != nil {
		log.Printf("Error querying transaction: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query transaction"})
		return
	}

	log.Printf("Query response: %s", response)
	ctx.JSON(http.StatusOK, gin.H{"message": "Query sent successfully"})
}

// TODO: after user pays, update database
func (p *PaymentImpl) VNPayIPN(ctx *gin.Context) {
	vnpParams := make(vo.VNPayParams)

	// Get all query parameters
	for key, values := range ctx.Request.URL.Query() {
		if len(values) > 0 {
			vnpParams[key] = values[0]
		}
	}

	secureHash := vnpParams["vnp_SecureHash"]
	orderId := vnpParams["vnp_TxnRef"]
	rspCode := vnpParams["vnp_ResponseCode"]

	// Remove hash fields for verification
	delete(vnpParams, "vnp_SecureHash")
	delete(vnpParams, "vnp_SecureHashType")

	// Sort parameters and verify signature
	sortedParams := payment.SortObject(vnpParams)
	signData := payment.BuildQueryString(sortedParams, false)
	signed := crypto.CreateHMAC(signData, global.Config.Payment.VnpHashSecret)

	// Payment status simulation
	paymentStatus := "0" // 0: Initial, 1: Success, 2: Failed
	checkOrderId := true // Check if order exists in database
	checkAmount := true  // Check if amount matches

	if secureHash == signed {
		if checkOrderId {
			if checkAmount {
				if paymentStatus == "0" {
					if rspCode == "00" {
						// Payment successful
						// Update payment status to success in database
						fmt.Printf("Payment successful for order: %s\n", orderId)
						ctx.JSON(http.StatusOK, vo.VNPayResponse{
							RspCode: "00",
							Message: "Success",
						})
					} else {
						// Payment failed
						// Update payment status to failed in database
						fmt.Printf("Payment failed for order: %s\n", orderId)
						ctx.JSON(http.StatusOK, vo.VNPayResponse{
							RspCode: "00",
							Message: "Success",
						})
					}
				} else {
					ctx.JSON(http.StatusOK, vo.VNPayResponse{
						RspCode: "02",
						Message: "This order has been updated to the payment status",
					})
				}
			} else {
				ctx.JSON(http.StatusOK, vo.VNPayResponse{
					RspCode: "04",
					Message: "Amount invalid",
				})
			}
		} else {
			ctx.JSON(http.StatusOK, vo.VNPayResponse{
				RspCode: "01",
				Message: "Order not found",
			})
		}
	} else {
		ctx.JSON(http.StatusOK, vo.VNPayResponse{
			RspCode: "97",
			Message: "Checksum failed",
		})
	}
}

// TODO: after user pays, redirect to frontend to render info to user
func (p *PaymentImpl) VNPayReturn(ctx *gin.Context) {
	vnpParams := make(vo.VNPayParams)

	// Get all query parameters
	for key, values := range ctx.Request.URL.Query() {
		if len(values) > 0 {
			vnpParams[key] = values[0]
		}
	}

	secureHash := vnpParams["vnp_SecureHash"]
	orderID := vnpParams["vnp_TxnRef"]
	responseCode := vnpParams["vnp_ResponseCode"]
	amount := vnpParams["vnp_Amount"]
	bankCode := vnpParams["vnp_BankCode"]
	transactionNo := vnpParams["vnp_TransactionNo"]
	payDate := vnpParams["vnp_PayDate"]

	// Remove hash fields for verification
	delete(vnpParams, "vnp_SecureHash")
	delete(vnpParams, "vnp_SecureHashType")

	// Sort parameters and verify signature
	sortedParams := payment.SortObject(vnpParams)
	signData := payment.BuildQueryString(sortedParams, false)
	signed := crypto.CreateHMAC(signData, global.Config.Payment.VnpHashSecret)

	if secureHash != signed {
		// Signature is valid - check with database and return result
		// code = vnpParams["vnp_ResponseCode"]
		// TODO: redirect to frontend
		p.redirectToReactWithError(ctx, "INVALID SIGNATURE", "Không đúng", orderID)
		return
	}

	// Convert amount to VND
	amountInt, _ := strconv.Atoi(amount)
	amountVND := amountInt / 100

	p.redirectToReactWithResult(ctx, vo.PaymentResultData{
		OrderID:       orderID,
		ResponseCode:  responseCode,
		Amount:        amountVND,
		BankCode:      bankCode,
		TransactionNo: transactionNo,
		PayDate:       payDate,
	})
}

// TODO: create payment url to redirect to VNPay
func (p *PaymentImpl) CreatePaymentURL(ctx *gin.Context, in *vo.CreatePaymentURLInput) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	now := time.Now().In(loc)

	createDate := now.Format("20060102150405")
	orderId := now.Format("02150405")

	ipAddr := ip.GetClientIP(ctx)

	locale := in.Language
	if locale == "" {
		locale = "vn"
	}

	// Create VNPay parameters
	vnpParams := vo.VNPayParams{
		"vnp_Version":    "2.1.0",
		"vnp_Command":    "pay",
		"vnp_TmnCode":    global.Config.Payment.VnpTmnCode,
		"vnp_Locale":     locale,
		"vnp_CurrCode":   "VND",
		"vnp_TxnRef":     orderId,
		"vnp_OrderInfo":  "Thanh toan cho ma GD:" + orderId,
		"vnp_OrderType":  "other",
		"vnp_Amount":     strconv.Itoa(in.Amount * 100),
		"vnp_ReturnUrl":  global.Config.Payment.VnpReturnUrl,
		"vnp_IpAddr":     ipAddr,
		"vnp_CreateDate": createDate,
	}

	// Add bank code if provided
	if in.BankCode != "" {
		vnpParams["vnp_BankCode"] = in.BankCode
	}

	// Sort parameters and create signature
	sortedParams := payment.SortObject(vnpParams)
	signData := payment.BuildQueryString(sortedParams, false)

	// Create HMAC SHA512 signature
	signature := crypto.CreateHMAC(signData, global.Config.Payment.VnpHashSecret)
	vnpParams["vnp_SecureHash"] = signature

	// Build final URL
	finalURL := global.Config.Payment.VnpUrl + "?" + payment.BuildQueryString(vnpParams, true)

	fmt.Printf("CreatePaymentURL success: %s\n", orderId)
	global.Logger.Info("CreatePaymentURL success: ", zap.String("info", orderId))

	ctx.Redirect(response.ErrCodeCreatePaymentURLSuccess, finalURL)
}

func (p *PaymentImpl) redirectToReactWithResult(ctx *gin.Context, data vo.PaymentResultData) {
	var status, message string

	switch data.ResponseCode {
	case "00":
		status = "success"
		message = "Thanh toán thành công"
	case "07":
		status = "cancelled"
		message = "Bạn đã hủy giao dịch"
	case "24":
		status = "expired"
		message = "Giao dịch đã hết hạn"
	case "09":
		status = "failed"
		message = "Thẻ chưa đăng ký Internet Banking"
	case "10":
		status = "failed"
		message = "Thông tin thẻ không chính xác"
	case "11":
		status = "failed"
		message = "Thẻ đã hết hạn"
	case "12":
		status = "failed"
		message = "Thẻ bị khóa"
	case "51":
		status = "failed"
		message = "Tài khoản không đủ số dư"
	case "65":
		status = "failed"
		message = "Vượt quá hạn mức giao dịch"
	default:
		status = "failed"
		message = "Giao dịch không thành công"
	}

	params := url.Values{}
	params.Set("status", status)
	params.Set("message", message)
	params.Set("order_id", data.OrderID)
	params.Set("response_code", data.ResponseCode)
	params.Set("amount", strconv.Itoa(data.Amount))

	if data.BankCode != "" {
		params.Set("bank_code", data.BankCode)
	}
	if data.TransactionNo != "" {
		params.Set("transaction_no", data.TransactionNo)
	}
	if data.PayDate != "" {
		params.Set("pay_date", data.PayDate)
	}

	// Add timestamp for cache busting
	params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	// Build React frontend URL
	frontendURL := fmt.Sprintf("%s/payment/result?%s",
		global.Config.Frontend.Url, params.Encode())

	log.Printf("Redirecting to React frontend: %s", frontendURL)

	// Redirect to React with all data in query params
	ctx.Redirect(http.StatusFound, frontendURL)
}

func (p *PaymentImpl) redirectToReactWithError(ctx *gin.Context, errorCode, message, orderID string) {
	params := url.Values{}
	params.Set("status", "error")
	params.Set("error_code", errorCode)
	params.Set("message", message)

	if orderID != "" {
		params.Set("order_id", orderID)
	}

	params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	frontendURL := fmt.Sprintf("%s/payment/result?%s",
		global.Config.Frontend.Url, params.Encode())

	log.Printf("Redirecting to React with error: %s", frontendURL)
	ctx.Redirect(http.StatusFound, frontendURL)
}

func NewPaymentImpl(sqlc *database.Queries, db *sql.DB) *PaymentImpl {
	return &PaymentImpl{
		sqlc: sqlc,
		db:   db,
	}
}
