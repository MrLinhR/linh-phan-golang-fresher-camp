package userbusiness

import (
	"FirstProject/common"
	"FirstProject/component"
	"FirstProject/component/tokenprovider"
	"FirstProject/modules/user/usermodel"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	passHashed := business.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayLoad{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	refreshToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	account := usermodel.NewAccount(accessToken, refreshToken)
	return account, nil
}
