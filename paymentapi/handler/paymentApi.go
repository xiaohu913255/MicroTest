package handler

import (
	"context"
	"errors"
	_ "github.com/plutov/paypal/v4"
	"github.com/xiaohu913255/MicroTest/common"
	go_micro_service_payment "github.com/xiaohu913255/MicroTest/payment/proto/payment"
	paymentApi "github.com/xiaohu913255/MicroTest/paymentApi/proto/paymentApi"

	"strconv"
)

type PaymentApi struct {
	PaymentService go_micro_service_payment.PaymentService
}

var (
	ClientID string = "AUn6IUU-wdZIAuR_nNh1kRy38X7h1-y0MuNC6Tip_S80u9wbGLcpVbbT4rYMYxKMBYKQ9Nt5U_VqogcB"
)

// PaymentApi.PayPalRefund 通过API向外暴露为/paymentApi/payPalRefund，接收http请求
// 即：/paymentApi/payPalRefund 请求会调用 go.micro.api.paymentApi 服务的PaymentApi.PayPalRefund
func (e *PaymentApi) PayPalRefund(ctx context.Context, req *paymentApi.Request, rsp *paymentApi.Response) error {
	//验证payment 支付通道是否赋值
	if err := isOK("payment_id", req); err != nil {
		rsp.StatusCode = 500
		return err
	}
	//验证 退款号
	if err := isOK("refund_id", req); err != nil {
		rsp.StatusCode = 500
		return err
	}
	//验证 退款金额
	if err := isOK("money", req); err != nil {
		rsp.StatusCode = 500
		return err
	}

	//获取paymentID
	payID, err := strconv.ParseInt(req.Get["payment_id"].Values[0], 10, 64)
	if err != nil {
		common.Error(err)
		return err
	}
	//获取支付通道信息
	paymentInfo, err := e.PaymentService.FindPaymentByID(ctx, &go_micro_service_payment.PaymentID{PaymentId: payID})
	if err != nil {
		common.Error(err)
		return err
	}
	//SID 获取 paymentInfo.PaymentSid
	//支付模式
	status := paypal.APIBaseSandBox
	if paymentInfo.PaymentStatus {
		status = paypal.APIBaseLive
	}
	//退款例子
	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			EmailSubject: req.Get["refund_id"].Values[0] + " imooc 提醒你收款！",
			EmailMessage: req.Get["refund_id"].Values[0] + " 您有一个收款信息！",
			//每笔转账都要唯一
			SenderBatchID: req.Get["refund_id"].Values[0],
		},
		Items: []paypal.PayoutItem{
			{
				RecipientType: "EMAIL",
				//RecipientWallet: "",
				Receiver: "sb-vvhq82259765@personal.example.com",
				Amount: &paypal.AmountPayout{
					//币种
					Currency: "USD",
					Value:    req.Get["money"].Values[0],
				},
				Note:         req.Get["refund_id"].Values[0],
				SenderItemID: req.Get["refund_id"].Values[0],
			},
		},
	}
	//创建支付客户端
	payPalClient, err := paypal.NewClient(ClientID, paymentInfo.PaymentSid, status)
	if err != nil {
		common.Error(err)
	}
	// 获取 token
	_, err = payPalClient.GetAccessToken()
	if err != nil {
		common.Error(err)
	}
	paymentResult, err := payPalClient.CreateSinglePayout(payout)
	if err != nil {
		common.Error(err)
	}
	common.Info(paymentResult)
	rsp.Body = req.Get["refund_id"].Values[0] + "支付成功！"
	return err
}

func isOK(key string, req *paymentApi.Request) error {
	if _, ok := req.Get[key]; !ok {
		err := errors.New(key + " 参数异常")
		common.Error(err)
		return err
	}
	return nil
}
