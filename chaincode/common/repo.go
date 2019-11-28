package common

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	IndexPrefix = "index:"
)

type BaseRepo struct {
	Logger     *shim.ChaincodeLogger
	IndexNames []IndexName
	Factory    DataFactory
}

func (repo *BaseRepo) Create(stub shim.ChaincodeStubInterface, entity Data) error {
	repo.Logger.Debug(fmt.Sprintf("Start creating entity : %s", entity.ToString()))

	baseKey := entity.GetBaseKey()
	err := repo.create(stub, baseKey, entity)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}

	indexKeys := GenIndexKeys(entity, repo.IndexNames)
	repo.Logger.Debug(fmt.Sprintf("Start creating index Keys : %v", indexKeys))
	return repo.addIndexes(stub, indexKeys, baseKey)
}

func (repo *BaseRepo) GetByBaseKey(stub shim.ChaincodeStubInterface, baseKey string) (Data, error) {
	repo.Logger.Debug(fmt.Sprintf("Start getting entity for key: %s\n", baseKey))
	result, err := repo.get(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (repo *BaseRepo) Update(stub shim.ChaincodeStubInterface, baseKey string, entity Data) error {
	repo.Logger.Debug(fmt.Sprintf("Start updating entity for key: %s", baseKey))

	originalEntity, err := repo.get(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if originalEntity == nil {
		repo.Logger.Debug(fmt.Sprintf("Entity not found for key: %s", baseKey))
		return nil
	}

	err = repo.update(stub, baseKey, entity)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}

	originalIndexKeys := GenIndexKeys(originalEntity, repo.IndexNames)
	newIndexKeys := GenIndexKeys(entity, repo.IndexNames)

	return repo.updateIndexes(stub, originalIndexKeys, newIndexKeys, baseKey)
}

func (repo *BaseRepo) Delete(stub shim.ChaincodeStubInterface, baseKey string) (Data, error) {
	repo.Logger.Debug(fmt.Sprintf("Start deleting entity for key: %s", baseKey))

	entity, err := repo.get(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	if entity == nil {
		repo.Logger.Debug(fmt.Sprintf("entity with key: %s does not exist", baseKey))
		return nil, nil
	}

	// delete entity
	err = repo.delete(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}

	// delete index
	indexKeys := GenIndexKeys(entity, repo.IndexNames)
	err = repo.removeIndexes(stub, indexKeys)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	return nil, nil
}
