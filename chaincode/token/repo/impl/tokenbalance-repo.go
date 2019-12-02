package impl

import (
	"errors"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain/obj"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type _TokenBalanceRepo struct {
	common.BaseRepo
}

func GenTokenBalanceRepo() *_TokenBalanceRepo {
	logger := shim.NewLogger("token-balance-repo")
	return &_TokenBalanceRepo{common.BaseRepo{
		Logger: logger,
		IndexNamePackage: common.IndexNamePackage{
			{Indicator: "ByToken", Names: []string{"token_name", "user_name"}},
			{Indicator: "ByUser", Names: []string{"user_name", "token_name"}},
		},
		Factory:       &obj.TokenBalanceDataFactory{},
		BaseKeyPrefix: "tokenbalance",
	}}
}

func (repo *_TokenBalanceRepo) GetBalance(stub shim.ChaincodeStubInterface, userName, tokenName string) (*obj.TokenBalance, error) {
	userBalancePrimaryKey := obj.GetTokenBalancePrimaryKey(userName, tokenName)
	userBalance, err := repo.GetByPrimaryKey(stub, userBalancePrimaryKey)
	if err != nil {
		return nil, err
	}
	if userBalance == nil {
		userBalance = obj.GetDefaultTokenBalance(userName, tokenName)
	}
	return userBalance.(*obj.TokenBalance), nil
}

func (repo *_TokenBalanceRepo) ListBalanceByUserName(stub shim.ChaincodeStubInterface,
	query []string) ([]*obj.TokenBalance, error) {
	resultList, err := repo.ListByIndexKey(stub, common.IndexSearchKey{Indicator: "ByUser", Keys: query})
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	tokenBalanceList := make([]*obj.TokenBalance, 0)
	for _, result := range resultList {
		tokenBalanceList = append(tokenBalanceList, result.(*obj.TokenBalance))
	}
	return tokenBalanceList, nil
}

func (repo *_TokenBalanceRepo) PaginateBalanceByUserName(stub shim.ChaincodeStubInterface,
	query []string, pageSize int32, bookMark string) ([]*obj.TokenBalance, string, error) {
	indexSearchKey := common.IndexSearchKey{Indicator: "ByUser", Keys: query}
	resultList, newBookMark, err := repo.PaginateByIndexKey(stub, indexSearchKey, pageSize, bookMark)
	if err != nil {
		return nil, "", err
	}
	finalResultList := make([]*obj.TokenBalance, 0)
	for _, result := range resultList {
		finalResultList = append(finalResultList, result.(*obj.TokenBalance))
	}
	return finalResultList, newBookMark, nil
}

func (repo *_TokenBalanceRepo) ListBalanceByTokenName(stub shim.ChaincodeStubInterface,
	query []string) ([]*obj.TokenBalance, error) {
	resultList, err := repo.ListByIndexKey(
		stub, common.IndexSearchKey{Indicator: "ByToken", Keys: query})
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	tokenBalanceList := make([]*obj.TokenBalance, 0)
	for _, result := range resultList {
		tokenBalanceList = append(tokenBalanceList, result.(*obj.TokenBalance))
	}
	return tokenBalanceList, nil
}

func (repo *_TokenBalanceRepo) PaginateBalanceByTokenName(stub shim.ChaincodeStubInterface,
	query []string, pageSize int32, bookMark string) ([]*obj.TokenBalance, string, error) {
	indexSearchKey := common.IndexSearchKey{Indicator: "ByToken", Keys: query}
	resultList, newBookMark, err := repo.PaginateByIndexKey(stub, indexSearchKey, pageSize, bookMark)
	if err != nil {
		return nil, "", err
	}
	finalResultList := make([]*obj.TokenBalance, 0)
	for _, result := range resultList {
		finalResultList = append(finalResultList, result.(*obj.TokenBalance))
	}
	return finalResultList, newBookMark, nil
}

func (repo *_TokenBalanceRepo) AddBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error {
	tokenBalancePrimaryKey := obj.GetTokenBalancePrimaryKey(userName, tokenName)
	oldTokenBalance, err := repo.GetByPrimaryKey(stub, tokenBalancePrimaryKey)

	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if oldTokenBalance == nil {
		oldTokenBalance = obj.GetDefaultTokenBalance(userName, tokenName)
		oldTokenBalance.(*obj.TokenBalance).AddBalance(amount)
		return repo.Create(stub, oldTokenBalance)
	} else {
		oldTokenBalance.(*obj.TokenBalance).AddBalance(amount)
		return repo.UpdateWithEntity(stub, oldTokenBalance, oldTokenBalance)
	}
}

func (repo *_TokenBalanceRepo) DeductBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error {
	tokenBalancePrimaryKey := obj.GetTokenBalancePrimaryKey(userName, tokenName)
	oldTokenBalance, err := repo.GetByPrimaryKey(stub, tokenBalancePrimaryKey)

	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if oldTokenBalance == nil {
		return errors.New("user has no balance")
	} else {
		entity := oldTokenBalance.(*obj.TokenBalance)
		entity.AddBalance(-amount)
		if entity.TokenAmount < 0 {
			return errors.New("token not enough")
		}
		return repo.UpdateWithEntity(stub, entity, entity)
	}
}
