package repo

import (
	"errors"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var TokenBalanceRepository *_TokenBalanceRepo

func init() {
	TokenBalanceRepository = genTokenBalanceRepo()
}

type _TokenBalanceRepo struct {
	common.BaseRepo
}

func genTokenBalanceRepo() *_TokenBalanceRepo {
	logger := shim.NewLogger("token-balance-repo")
	logger.SetLevel(shim.LogDebug)
	return &_TokenBalanceRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"token_name", "user_name"},
		},
		Factory:       &domain.TokenBalanceDataFactory{},
		BaseKeyPrefix: "tokenbalance",
	}}
}

func (repo *_TokenBalanceRepo) AddBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error {
	tokenBalancePrimaryKey := domain.GetTokenBalancePrimaryKey(userName, tokenName)
	oldTokenBalance, err := repo.GetByPrimaryKey(stub, tokenBalancePrimaryKey)

	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if oldTokenBalance == nil {
		oldTokenBalance = domain.GetDefaultTokenBalance(userName, tokenName)
		oldTokenBalance.(*domain.TokenBalance).AddBalance(amount)
		return repo.Create(stub, oldTokenBalance)
	} else {
		oldTokenBalance.(*domain.TokenBalance).AddBalance(amount)
		return repo.UpdateWithEntity(stub, oldTokenBalance, oldTokenBalance)
	}
}

func (repo *_TokenBalanceRepo) DeductBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error{
	tokenBalancePrimaryKey := domain.GetTokenBalancePrimaryKey(userName, tokenName)
	oldTokenBalance, err := repo.GetByPrimaryKey(stub, tokenBalancePrimaryKey)

	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if oldTokenBalance == nil {
		return errors.New("user has no balance")
	} else {
		entity := oldTokenBalance.(*domain.TokenBalance)
		entity.AddBalance(-amount)
		if entity.TokenAmount < 0{
			return errors.New("token not enough")
		}
		return repo.UpdateWithEntity(stub, entity, entity)
	}
}
