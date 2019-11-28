package domain

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/lib"
	"strconv"
	"time"
)

type TokenCreation struct {
	TokenName     string    `json:"token_name"`
	MaxAmount     int64     `json:"max_amount"`
	CurrentAmount int64     `json:"current_amount"`
	Creator       string    `json:"creator"`
	Issuer        string    `json:"issuer"`
	CreatedTime   time.Time `json:"created_time"`
}

func (token *TokenCreation) ToBytes() ([]byte, error) {
	return lib.ToJsonBytes(token)
}

func (token *TokenCreation) ToString() string {
	return lib.ToJsonString(token)
}

func (token *TokenCreation) ToMap() map[string]string {
	return map[string]string{
		"token":   token.TokenName,
		"max":     strconv.FormatInt(token.MaxAmount, 10),
		"current": strconv.FormatInt(token.CurrentAmount, 10),
		"creator": token.Creator,
		"issuer":  token.Issuer,
	}
}

func (token *TokenCreation) GetPrimaryKey() string {
	return GetTokenCreationPrimaryKey(token.TokenName)
}

func GetTokenCreationPrimaryKey(tokenName string) string {
	return fmt.Sprintf("%s", tokenName)
}

type TokenCreationDataFactory struct {
}

func (factory *TokenCreationDataFactory) ToDataEntity(data []byte) (common.Data, error) {
	result := &TokenCreation{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
