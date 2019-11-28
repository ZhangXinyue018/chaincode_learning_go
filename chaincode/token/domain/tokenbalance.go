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
		"user_name":   userToken.UserName,
		"token_name":  userToken.TokenName,
		"token_amount": strconv.FormatInt(userToken.TokenAmount, 10),
	}
}

func (userToken *TokenBalance) GetPrimaryKey() string {
	return GetTokenBalancePrimaryKey(userToken.UserName, userToken.TokenName)
}

func (userToken *TokenBalance) AddBalance(amount int64) {
	userToken.TokenAmount = userToken.TokenAmount + amount
}

func GetTokenBalancePrimaryKey(userName, tokenName string) string {
	return fmt.Sprintf("%s:%s", userName, tokenName)
}

func GetDefaultTokenBalance(userName, tokenName string) *TokenBalance {
	return &TokenBalance{
		UserName:    userName,
		TokenName:   tokenName,
		TokenAmount: 0,
	}
}

type TokenBalanceDataFactory struct {
}

func (factory *TokenBalanceDataFactory) ToDataEntity(data []byte) (common.Data, error) {
	if data == nil {
		return nil, nil
	}
	result := &TokenBalance{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
