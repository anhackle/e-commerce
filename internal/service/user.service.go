package service

import (
	"context"

	"github.com/anle/codebase/model"
)

type (
	// interfaces
	IUserAuthen interface {
		Login(ctx context.Context) error
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out *model.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context) error
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin  IUserAdmin
	localUserInfo   IUserInfo
	localUserAuthen IUserAuthen
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement localUserAdmin not found for inteface IUserAdmin")
	}

	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement localUserInfo not found for inteface IUserInfo")
	}

	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserAuthen() IUserAuthen {
	if localUserAuthen == nil {
		panic("implement localUserAuthen not found for inteface IUserAuthen")
	}

	return localUserAuthen
}

func InitUserAuthen(i IUserAuthen) {
	localUserAuthen = i
}
