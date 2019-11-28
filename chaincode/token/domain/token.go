package domain

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/lib"
	"strconv"
	"time"
)

type Token struct {
	TokenName     string    `json:"token_name"`
	MaxAmount     int64     `json:"max_amount"`
	CurrentAmount int64     `json:"current_amount"`
	Creator       string    `json:"creator"`
	Issuer        string    `json:"issuer"`
	CreatedTime   time.Time `json:"created_time"`
}

func (token *Token) ToBytes() ([]byte, error) {
	return lib.ToJsonBytes(token)
}

func (token *Token) ToString() string {
	return lib.ToJsonString(token)
}

func (token *Token) ToMap() map[string]string {
	return map[string]string{
		"token":   token.TokenName,
		"max":     strconv.FormatInt(token.MaxAmount, 10),
		"current": strconv.FormatInt(token.CurrentAmount, 10),
		"creator": token.Creator,
		"issuer":  token.Issuer,
	}
}

func (token *Token) GetBaseKey() string {
	return fmt.Sprintf("%s:%s", "creation", token.TokenName)
}

type TokenDataFactory struct {
}

func (factory *TokenDataFactory) ToDataEntity(data []byte) (common.Data, error) {
	result := &Token{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
