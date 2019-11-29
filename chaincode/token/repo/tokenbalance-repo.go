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
	return &_TokenBalanceRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: common.IndexNamePackage{
			{Indicator: "ByToken", Names: []string{"token_name", "user_name"}},
			{Indicator: "ByUser", Names: []string{"user_name", "token_name"}},
		},
		Factory:       &domain.TokenBalanceDataFactory{},
		BaseKeyPrefix: "tokenbalance",
	}}
}

func (repo *_TokenBalanceRepo) GetBalance(stub shim.ChaincodeStubInterface, userName, tokenName string) (*domain.TokenBalance, error) {
	userBalancePrimaryKey := domain.GetTokenBalancePrimaryKey(userName, tokenName)
	userBalance, err := repo.GetByPrimaryKey(stub, userBalancePrimaryKey)
	if err != nil {
		return nil, err
	}
	if userBalance == nil {
		userBalance = domain.GetDefaultTokenBalance(userName, tokenName)
	}
	return userBalance.(*domain.TokenBalance), nil
}

func (repo *_TokenBalanceRepo) ListBalanceByUserName(stub shim.ChaincodeStubInterface,
	query []string) ([]*domain.TokenBalance, error) {
	resultList, err := repo.ListByIndexKey(stub, common.IndexSearchKey{Indicator: "ByUser", Keys: query})
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	tokenBalanceList := make([]*domain.TokenBalance, 0)
	for _, result := range resultList {
		tokenBalanceList = append(tokenBalanceList, result.(*domain.TokenBalance))
	}
	return tokenBalanceList, nil
}

func (repo *_TokenBalanceRepo) PaginateBalanceByUserName(stub shim.ChaincodeStubInterface,
	query []string, pageSize int32, bookMark string) (*common.PaginationResponse, error) {
	return repo.PaginateByIndexKey(
		stub, common.IndexSearchKey{Indicator: "ByUser", Keys: query},
		pageSize, bookMark)
}

func (repo *_TokenBalanceRepo) ListBalanceByTokenName(stub shim.ChaincodeStubInterface,
	query []string) ([]*domain.TokenBalance, error) {
	resultList, err := repo.ListByIndexKey(
		stub, common.IndexSearchKey{Indicator: "ByToken", Keys: query})
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	tokenBalanceList := make([]*domain.TokenBalance, 0)
	for _, result := range resultList {
		tokenBalanceList = append(tokenBalanceList, result.(*domain.TokenBalance))
	}
	return tokenBalanceList, nil
}

func (repo *_TokenBalanceRepo) PaginateBalanceByTokenName(stub shim.ChaincodeStubInterface,
	query []string, pageSize int32, bookMark string) (*common.PaginationResponse, error) {
	return repo.PaginateByIndexKey(
		stub, common.IndexSearchKey{Indicator: "ByToken", Keys: query}, pageSize, bookMark)
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

func (repo *_TokenBalanceRepo) DeductBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error {
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
		if entity.TokenAmount < 0 {
			return errors.New("token not enough")
		}
		return repo.UpdateWithEntity(stub, entity, entity)
	}
}
