package domain

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/lib"
	"strconv"
)

type UserToken struct {
	UserName    string `json:"user_name"`
	TokenName   string `json:"token_name"`
	TokenAmount int64  `json:"token_amount"`
}

func (userToken *UserToken) ToBytes() ([]byte, error) {
	return lib.ToJsonBytes(userToken)
}

func (userToken *UserToken) ToString() string {
	return lib.ToJsonString(userToken)
}

func (userToken *UserToken) ToMap() map[string]string {
	return map[string]string{
		"user":   userToken.UserName,
		"token":  userToken.TokenName,
		"amount": strconv.FormatInt(userToken.TokenAmount, 10),
	}
}

func (userToken *UserToken) GetBaseKey() string {
	return fmt.Sprintf("%s:%s:%s", "token", userToken.UserName, userToken.TokenName)
}

type UserTokenDataFactory struct {
}

func (factory *UserTokenDataFactory) ToDataEntity(data []byte) (common.Data, error) {
	result := &UserToken{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
