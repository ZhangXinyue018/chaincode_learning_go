package common

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	IndexPrefix = "index"
)

func (repo *BaseRepo) create(stub shim.ChaincodeStubInterface, key string, entity Data) error {
	entityBytes, err := entity.ToBytes()
	if err != nil {
		return err
	}
	return stub.PutState(key, entityBytes)
}

func (repo *BaseRepo) update(stub shim.ChaincodeStubInterface, key string, entity Data) error {
	return repo.create(stub, key, entity)
}

func (repo *BaseRepo) updateIndexes(stub shim.ChaincodeStubInterface, originalIndexKeys, newIndexKeys []IndexKey, baseKey string) error {
	originalIndexMap := map[string]bool{}
	for _, originalIndexKey := range originalIndexKeys {
		compositeKey, err := stub.CreateCompositeKey(repo.getIndexKeyPrefix(), originalIndexKey)
		if err != nil {
			return err
		}
		originalIndexMap[compositeKey] = true
	}

	newIndexMap := map[string]bool{}
	for _, newIndexKey := range newIndexKeys {
		compositeKey, err := stub.CreateCompositeKey(repo.getIndexKeyPrefix(), newIndexKey)
		if err != nil {
			return err
		}
		newIndexMap[compositeKey] = true
	}

	for key := range newIndexMap {
		if _, ok := originalIndexMap[key]; ok {
			delete(originalIndexMap, key)
		} else {
			err := stub.PutState(key, []byte(baseKey))
			if err != nil {
				return err
			}
		}
	}

	for key := range originalIndexMap {
		err := stub.DelState(key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *BaseRepo) addIndexes(stub shim.ChaincodeStubInterface, indexKeys []IndexKey, baseKey string) error {
	for _, indexKey := range indexKeys {
		compositeKey, err := stub.CreateCompositeKey(repo.getIndexKeyPrefix(), indexKey)
		if err != nil {
			return err
		}
		repo.Logger.Debugf("Add index key: %s, value: %s", compositeKey, baseKey)
		err = stub.PutState(compositeKey, []byte(baseKey))
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *BaseRepo) removeIndexes(stub shim.ChaincodeStubInterface, indexKeys []IndexKey) error {
	for _, indexKey := range indexKeys {
		compositeKey, err := stub.CreateCompositeKey(repo.getIndexKeyPrefix(), indexKey)
		if err != nil {
			return err
		}

		err = stub.DelState(compositeKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *BaseRepo) get(stub shim.ChaincodeStubInterface, key string) (Data, error) {
	entityBytes, err := stub.GetState(key)
	if err != nil {
		repo.Logger.Error(err.Error())
		return nil, err
	}

	return repo.Factory.ToDataEntity(entityBytes)
}

func (repo *BaseRepo) delete(stub shim.ChaincodeStubInterface, key string) error {
	repo.Logger.Debugf("Start deleting entity for key: %s", key)
	return stub.DelState(key)
}

func (repo *BaseRepo) getIndexKeyPrefix() string {
	return fmt.Sprintf("%s:%s:", IndexPrefix, repo.BaseKeyPrefix)
}

func (repo *BaseRepo) getBaseKeyByPrimaryKey(key string) string {
	return fmt.Sprintf("%s:%s", repo.BaseKeyPrefix, key)
}

func (repo *BaseRepo) getBaseKeyByEntity(entity Data) string {
	return repo.getBaseKeyByPrimaryKey(entity.GetPrimaryKey())
}
