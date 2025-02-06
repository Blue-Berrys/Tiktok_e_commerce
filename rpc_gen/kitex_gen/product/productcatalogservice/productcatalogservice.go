// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	"errors"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	product "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ListProducts": kitex.NewMethodInfo(
		listProductsHandler,
		newListProductsArgs,
		newListProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetProduct": kitex.NewMethodInfo(
		getProductHandler,
		newGetProductArgs,
		newGetProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SearchProducts": kitex.NewMethodInfo(
		searchProductsHandler,
		newSearchProductsArgs,
		newSearchProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CreateProduct": kitex.NewMethodInfo(
		createProductHandler,
		newCreateProductArgs,
		newCreateProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteProduct": kitex.NewMethodInfo(
		deleteProductHandler,
		newDeleteProductArgs,
		newDeleteProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateProductInfo": kitex.NewMethodInfo(
		updateProductInfoHandler,
		newUpdateProductInfoArgs,
		newUpdateProductInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productCatalogServiceServiceInfo                = NewServiceInfo()
	productCatalogServiceServiceInfoForClient       = NewServiceInfoForClient()
	productCatalogServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ProductCatalogService"
	handlerType := (*product.ProductCatalogService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "product",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func listProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.ListProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).ListProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListProductsArgs:
		success, err := handler.(product.ProductCatalogService).ListProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListProductsArgs() interface{} {
	return &ListProductsArgs{}
}

func newListProductsResult() interface{} {
	return &ListProductsResult{}
}

type ListProductsArgs struct {
	Req *product.ListProductsReq
}

func (p *ListProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.ListProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListProductsArgs) Unmarshal(in []byte) error {
	msg := new(product.ListProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListProductsArgs_Req_DEFAULT *product.ListProductsReq

func (p *ListProductsArgs) GetReq() *product.ListProductsReq {
	if !p.IsSetReq() {
		return ListProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListProductsResult struct {
	Success *product.ListProductsResp
}

var ListProductsResult_Success_DEFAULT *product.ListProductsResp

func (p *ListProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.ListProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListProductsResult) Unmarshal(in []byte) error {
	msg := new(product.ListProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListProductsResult) GetSuccess() *product.ListProductsResp {
	if !p.IsSetSuccess() {
		return ListProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.ListProductsResp)
}

func (p *ListProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListProductsResult) GetResult() interface{} {
	return p.Success
}

func getProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.GetProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).GetProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetProductArgs:
		success, err := handler.(product.ProductCatalogService).GetProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetProductArgs() interface{} {
	return &GetProductArgs{}
}

func newGetProductResult() interface{} {
	return &GetProductResult{}
}

type GetProductArgs struct {
	Req *product.GetProductReq
}

func (p *GetProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.GetProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetProductArgs) Unmarshal(in []byte) error {
	msg := new(product.GetProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetProductArgs_Req_DEFAULT *product.GetProductReq

func (p *GetProductArgs) GetReq() *product.GetProductReq {
	if !p.IsSetReq() {
		return GetProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetProductResult struct {
	Success *product.GetProductResp
}

var GetProductResult_Success_DEFAULT *product.GetProductResp

func (p *GetProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.GetProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetProductResult) Unmarshal(in []byte) error {
	msg := new(product.GetProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProductResult) GetSuccess() *product.GetProductResp {
	if !p.IsSetSuccess() {
		return GetProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.GetProductResp)
}

func (p *GetProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetProductResult) GetResult() interface{} {
	return p.Success
}

func searchProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.SearchProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).SearchProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SearchProductsArgs:
		success, err := handler.(product.ProductCatalogService).SearchProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSearchProductsArgs() interface{} {
	return &SearchProductsArgs{}
}

func newSearchProductsResult() interface{} {
	return &SearchProductsResult{}
}

type SearchProductsArgs struct {
	Req *product.SearchProductsReq
}

func (p *SearchProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.SearchProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchProductsArgs) Unmarshal(in []byte) error {
	msg := new(product.SearchProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchProductsArgs_Req_DEFAULT *product.SearchProductsReq

func (p *SearchProductsArgs) GetReq() *product.SearchProductsReq {
	if !p.IsSetReq() {
		return SearchProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchProductsResult struct {
	Success *product.SearchProductsResp
}

var SearchProductsResult_Success_DEFAULT *product.SearchProductsResp

func (p *SearchProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.SearchProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchProductsResult) Unmarshal(in []byte) error {
	msg := new(product.SearchProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchProductsResult) GetSuccess() *product.SearchProductsResp {
	if !p.IsSetSuccess() {
		return SearchProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.SearchProductsResp)
}

func (p *SearchProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchProductsResult) GetResult() interface{} {
	return p.Success
}

func createProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.CreateProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).CreateProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateProductArgs:
		success, err := handler.(product.ProductCatalogService).CreateProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateProductArgs() interface{} {
	return &CreateProductArgs{}
}

func newCreateProductResult() interface{} {
	return &CreateProductResult{}
}

type CreateProductArgs struct {
	Req *product.CreateProductReq
}

func (p *CreateProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.CreateProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateProductArgs) Unmarshal(in []byte) error {
	msg := new(product.CreateProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateProductArgs_Req_DEFAULT *product.CreateProductReq

func (p *CreateProductArgs) GetReq() *product.CreateProductReq {
	if !p.IsSetReq() {
		return CreateProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateProductResult struct {
	Success *product.CreateProductResp
}

var CreateProductResult_Success_DEFAULT *product.CreateProductResp

func (p *CreateProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.CreateProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateProductResult) Unmarshal(in []byte) error {
	msg := new(product.CreateProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateProductResult) GetSuccess() *product.CreateProductResp {
	if !p.IsSetSuccess() {
		return CreateProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.CreateProductResp)
}

func (p *CreateProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateProductResult) GetResult() interface{} {
	return p.Success
}

func deleteProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.DeleteProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).DeleteProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteProductArgs:
		success, err := handler.(product.ProductCatalogService).DeleteProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteProductArgs() interface{} {
	return &DeleteProductArgs{}
}

func newDeleteProductResult() interface{} {
	return &DeleteProductResult{}
}

type DeleteProductArgs struct {
	Req *product.DeleteProductReq
}

func (p *DeleteProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.DeleteProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteProductArgs) Unmarshal(in []byte) error {
	msg := new(product.DeleteProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteProductArgs_Req_DEFAULT *product.DeleteProductReq

func (p *DeleteProductArgs) GetReq() *product.DeleteProductReq {
	if !p.IsSetReq() {
		return DeleteProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteProductResult struct {
	Success *common.Empty
}

var DeleteProductResult_Success_DEFAULT *common.Empty

func (p *DeleteProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(common.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteProductResult) Unmarshal(in []byte) error {
	msg := new(common.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteProductResult) GetSuccess() *common.Empty {
	if !p.IsSetSuccess() {
		return DeleteProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*common.Empty)
}

func (p *DeleteProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteProductResult) GetResult() interface{} {
	return p.Success
}

func updateProductInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.UpdateProductInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).UpdateProductInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateProductInfoArgs:
		success, err := handler.(product.ProductCatalogService).UpdateProductInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateProductInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateProductInfoArgs() interface{} {
	return &UpdateProductInfoArgs{}
}

func newUpdateProductInfoResult() interface{} {
	return &UpdateProductInfoResult{}
}

type UpdateProductInfoArgs struct {
	Req *product.UpdateProductInfoReq
}

func (p *UpdateProductInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.UpdateProductInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateProductInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateProductInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateProductInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateProductInfoArgs) Unmarshal(in []byte) error {
	msg := new(product.UpdateProductInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateProductInfoArgs_Req_DEFAULT *product.UpdateProductInfoReq

func (p *UpdateProductInfoArgs) GetReq() *product.UpdateProductInfoReq {
	if !p.IsSetReq() {
		return UpdateProductInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateProductInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateProductInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateProductInfoResult struct {
	Success *common.Empty
}

var UpdateProductInfoResult_Success_DEFAULT *common.Empty

func (p *UpdateProductInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(common.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateProductInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateProductInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateProductInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateProductInfoResult) Unmarshal(in []byte) error {
	msg := new(common.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateProductInfoResult) GetSuccess() *common.Empty {
	if !p.IsSetSuccess() {
		return UpdateProductInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateProductInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*common.Empty)
}

func (p *UpdateProductInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateProductInfoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ListProducts(ctx context.Context, Req *product.ListProductsReq) (r *product.ListProductsResp, err error) {
	var _args ListProductsArgs
	_args.Req = Req
	var _result ListProductsResult
	if err = p.c.Call(ctx, "ListProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProduct(ctx context.Context, Req *product.GetProductReq) (r *product.GetProductResp, err error) {
	var _args GetProductArgs
	_args.Req = Req
	var _result GetProductResult
	if err = p.c.Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchProducts(ctx context.Context, Req *product.SearchProductsReq) (r *product.SearchProductsResp, err error) {
	var _args SearchProductsArgs
	_args.Req = Req
	var _result SearchProductsResult
	if err = p.c.Call(ctx, "SearchProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateProduct(ctx context.Context, Req *product.CreateProductReq) (r *product.CreateProductResp, err error) {
	var _args CreateProductArgs
	_args.Req = Req
	var _result CreateProductResult
	if err = p.c.Call(ctx, "CreateProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq) (r *common.Empty, err error) {
	var _args DeleteProductArgs
	_args.Req = Req
	var _result DeleteProductResult
	if err = p.c.Call(ctx, "DeleteProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateProductInfo(ctx context.Context, Req *product.UpdateProductInfoReq) (r *common.Empty, err error) {
	var _args UpdateProductInfoArgs
	_args.Req = Req
	var _result UpdateProductInfoResult
	if err = p.c.Call(ctx, "UpdateProductInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
