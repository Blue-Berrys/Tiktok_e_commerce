// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newRegisterArgs,
		newRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteUser": kitex.NewMethodInfo(
		deleteUserHandler,
		newDeleteUserArgs,
		newDeleteUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserInfo": kitex.NewMethodInfo(
		updateUserInfoHandler,
		newUpdateUserInfoArgs,
		newUpdateUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
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
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.RegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RegisterArgs:
		success, err := handler.(user.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.RegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.RegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.RegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.RegisterReq

func (p *RegisterArgs) GetReq() *user.RegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *user.RegisterResp
}

var RegisterResult_Success_DEFAULT *user.RegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.RegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.RegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.RegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.RegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.LoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(user.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.LoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.LoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.LoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.LoginReq

func (p *LoginArgs) GetReq() *user.LoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *user.LoginResp
}

var LoginResult_Success_DEFAULT *user.LoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.LoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.LoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.LoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DeleteUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).DeleteUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteUserArgs:
		success, err := handler.(user.UserService).DeleteUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteUserArgs() interface{} {
	return &DeleteUserArgs{}
}

func newDeleteUserResult() interface{} {
	return &DeleteUserResult{}
}

type DeleteUserArgs struct {
	Req *user.DeleteUserReq
}

func (p *DeleteUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DeleteUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteUserArgs) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteUserArgs_Req_DEFAULT *user.DeleteUserReq

func (p *DeleteUserArgs) GetReq() *user.DeleteUserReq {
	if !p.IsSetReq() {
		return DeleteUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteUserResult struct {
	Success *common.Empty
}

var DeleteUserResult_Success_DEFAULT *common.Empty

func (p *DeleteUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(common.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteUserResult) Unmarshal(in []byte) error {
	msg := new(common.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserResult) GetSuccess() *common.Empty {
	if !p.IsSetSuccess() {
		return DeleteUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*common.Empty)
}

func (p *DeleteUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUserResult) GetResult() interface{} {
	return p.Success
}

func updateUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserInfoArgs:
		success, err := handler.(user.UserService).UpdateUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserInfoArgs() interface{} {
	return &UpdateUserInfoArgs{}
}

func newUpdateUserInfoResult() interface{} {
	return &UpdateUserInfoResult{}
}

type UpdateUserInfoArgs struct {
	Req *user.UpdateUserInfoReq
}

func (p *UpdateUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserInfoArgs_Req_DEFAULT *user.UpdateUserInfoReq

func (p *UpdateUserInfoArgs) GetReq() *user.UpdateUserInfoReq {
	if !p.IsSetReq() {
		return UpdateUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserInfoResult struct {
	Success *common.Empty
}

var UpdateUserInfoResult_Success_DEFAULT *common.Empty

func (p *UpdateUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(common.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserInfoResult) Unmarshal(in []byte) error {
	msg := new(common.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserInfoResult) GetSuccess() *common.Empty {
	if !p.IsSetSuccess() {
		return UpdateUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*common.Empty)
}

func (p *UpdateUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserInfoResult) GetResult() interface{} {
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

func (p *kClient) Register(ctx context.Context, Req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, Req *user.DeleteUserReq) (r *common.Empty, err error) {
	var _args DeleteUserArgs
	_args.Req = Req
	var _result DeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq) (r *common.Empty, err error) {
	var _args UpdateUserInfoArgs
	_args.Req = Req
	var _result UpdateUserInfoResult
	if err = p.c.Call(ctx, "UpdateUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
