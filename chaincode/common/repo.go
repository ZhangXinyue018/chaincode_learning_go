package common

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type BaseRepo struct {
	Logger        *shim.ChaincodeLogger
	IndexNames    []IndexName
	Factory       DataFactory
	BaseKeyPrefix string
}

func (repo *BaseRepo) Create(stub shim.ChaincodeStubInterface, entity Data) error {
	repo.Logger.Debugf("Start creating entity : %s", entity.ToString())

	baseKey := repo.getBaseKeyByEntity(entity)
	err := repo.create(stub, baseKey, entity)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	indexKeys := GenIndexKeys(entity, repo.IndexNames)
	repo.Logger.Debugf("Start creating index Keys : %v", indexKeys)
	return repo.addIndexes(stub, indexKeys, baseKey)
}

func (repo *BaseRepo) GetByPrimaryKey(stub shim.ChaincodeStubInterface, primaryKey string) (Data, error) {
	baseKey := repo.getBaseKeyByPrimaryKey(primaryKey)
	repo.Logger.Debugf("Start getting entity for key: %s\n", baseKey)
	result, err := repo.get(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	return result, nil
}

func (repo *BaseRepo) UpdateWithEntity(stub shim.ChaincodeStubInterface, originalEntity Data, entity Data) error {
	repo.Logger.Debugf("Update entity: %v", entity)
	baseKey := repo.getBaseKeyByEntity(entity)
	err := repo.update(stub, baseKey, entity)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}

	originalIndexKeys := GenIndexKeys(originalEntity, repo.IndexNames)
	newIndexKeys := GenIndexKeys(entity, repo.IndexNames)

	return repo.updateIndexes(stub, originalIndexKeys, newIndexKeys, baseKey)
}

func (repo *BaseRepo) UpdateWithPrimaryKey(stub shim.ChaincodeStubInterface, primaryKey string, entity Data) error {
	baseKey := repo.getBaseKeyByPrimaryKey(primaryKey)
	repo.Logger.Debugf("Start updating entity for key: %s", baseKey)

	originalEntity, err := repo.get(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	if originalEntity == nil {
		repo.Logger.Debugf("Entity not found for key: %s", baseKey)
		return nil
	}

	return repo.UpdateWithEntity(stub, originalEntity, entity)
}

func (repo *BaseRepo) Delete(stub shim.ChaincodeStubInterface, primaryKey string) (Data, error) {
	baseKey := repo.getBaseKeyByPrimaryKey(primaryKey)
	repo.Logger.Debugf("Start deleting entity for key: %s", baseKey)

	entity, err := repo.get(stub, baseKey)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}
	if entity == nil {
		repo.Logger.Debugf("entity with key: %s does not exist", baseKey)
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
