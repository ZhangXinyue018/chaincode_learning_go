package repo

import (
	"errors"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var TokenCreationRepository *_TokenCreationRepo

func init() {
	TokenCreationRepository = genTokenCreationRepo()
}

type _TokenCreationRepo struct {
	common.BaseRepo
}

func genTokenCreationRepo() *_TokenCreationRepo {
	logger := shim.NewLogger("token-creation-repo")
	return &_TokenCreationRepo{common.BaseRepo{
		Logger: logger,
		IndexNamePackage: common.IndexNamePackage{
			{Indicator: "ByCreator", Names: []string{"creator", "token_name"}},
		},
		Factory:       &domain.TokenCreationDataFactory{},
		BaseKeyPrefix: "tokencreation",
	}}
}

func (repo *_TokenCreationRepo) UpdateTokenIssueAmount(stub shim.ChaincodeStubInterface, userName, tokenName string,
	tokenAmount int64) error {
	tokenCreationPrimaryKey := domain.GetTokenCreationPrimaryKey(tokenName)
	token, err := repo.GetByPrimaryKey(stub, tokenCreationPrimaryKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if token == nil {
		return errors.New("token not exist")
	}

	tokenEntity := token.(*domain.TokenCreation)
	if tokenEntity.Issuer != userName {
		return errors.New("token issue is not allowed")
	}
	if tokenEntity.CurrentAmount+tokenAmount > tokenEntity.MaxAmount {
		return errors.New("amount exceed")
	}

	tokenEntity.CurrentAmount = tokenEntity.CurrentAmount + tokenAmount
	err = repo.UpdateWithEntity(stub, tokenEntity, tokenEntity)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	return nil
}
