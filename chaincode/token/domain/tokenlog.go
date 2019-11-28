package domain

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/lib"
	"strconv"
)

type TokenLog struct {
	TokenLogId   string `json:"token_log_id"`
	FromUserName string `json:"from_user_name"`
	ToUserName   string `json:"to_user_name"`
	TokenName    string `json:"token_name"`
	TokenAmount  int64  `json:"token_amount"`
	Comment      string `json:"comment"`
}

func (tokenlog *TokenLog) ToBytes() ([]byte, error) {
	return lib.ToJsonBytes(tokenlog)
}

func (tokenlog *TokenLog) ToString() string {
	return lib.ToJsonString(tokenlog)
}

func (tokenlog *TokenLog) ToMap() map[string]string {
	return map[string]string{
		"fromuser": tokenlog.FromUserName,
		"touser":   tokenlog.ToUserName,
		"token":    tokenlog.TokenName,
		"amount":   strconv.FormatInt(tokenlog.TokenAmount, 10),
		"comment":  tokenlog.Comment,
	}
}

func (tokenlog *TokenLog) GetBaseKey() string {
	return fmt.Sprintf("%s:%s", "log", tokenlog.TokenLogId)
}

type TokenLogDataFactory struct {
}

func (factory *TokenLogDataFactory) ToDataEntity(data []byte) (common.Data, error) {
	result := &TokenLog{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
