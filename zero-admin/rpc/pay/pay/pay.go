// Code generated by goctl. DO NOT EDIT!
// Source: pay.proto

package pay

import (
	"context"

	"zero-admin/rpc/pay/payclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	H5UnifiedOrderResp = payclient.H5UnifiedOrderResp
	OrderQueryReq      = payclient.OrderQueryReq
	OrderQueryResp     = payclient.OrderQueryResp
	UnifiedOrderReq    = payclient.UnifiedOrderReq
	UnifiedOrderResp   = payclient.UnifiedOrderResp

	Pay interface {
		AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error)
		H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*H5UnifiedOrderResp, error)
		JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error)
		OrderQuery(ctx context.Context, in *OrderQueryReq, opts ...grpc.CallOption) (*OrderQueryResp, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

func (m *defaultPay) AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error) {
	client := payclient.NewPayClient(m.cli.Conn())
	return client.AppUnifiedOrder(ctx, in, opts...)
}

func (m *defaultPay) H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*H5UnifiedOrderResp, error) {
	client := payclient.NewPayClient(m.cli.Conn())
	return client.H5UnifiedOrder(ctx, in, opts...)
}

func (m *defaultPay) JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error) {
	client := payclient.NewPayClient(m.cli.Conn())
	return client.JsUnifiedOrder(ctx, in, opts...)
}

func (m *defaultPay) OrderQuery(ctx context.Context, in *OrderQueryReq, opts ...grpc.CallOption) (*OrderQueryResp, error) {
	client := payclient.NewPayClient(m.cli.Conn())
	return client.OrderQuery(ctx, in, opts...)
}
