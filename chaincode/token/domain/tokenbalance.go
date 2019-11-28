package domain

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/lib"
	"strconv"
)

type TokenBalance struct {
	UserName    string `json:"user_name"`
	TokenName   string `json:"token_name"`
	TokenAmount int64  `json:"token_amount"`
}

func (userToken *TokenBalance) ToBytes() ([]byte, error) {
	return lib.ToJsonBytes(userToken)
}

func (userToken *TokenBalance) ToString() string {
	return lib.ToJsonString(userToken)
}

func (userToken *TokenBalance) ToMap() map[string]string {
	return map[string]string{
		"user":   userToken.UserName,
		"token":  userToken.TokenName,
		"amount": strconv.FormatInt(userToken.TokenAmount, 10),
	}
}

func (userToken *TokenBalance) GetPrimaryKey() string {
	return GetTokenBalancePrimaryKey(userToken.UserName, userToken.TokenName)
}

func GetTokenBalancePrimaryKey(userName, tokenName string) string {
	return fmt.Sprintf("%s:%s", userName, tokenName)
}

type TokenBalanceDataFactory struct {
}

func (factory *TokenBalanceDataFactory) ToDataEntity(data []byte) (common.Data, error) {
	result := &TokenBalance{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
