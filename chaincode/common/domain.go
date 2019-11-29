package common

import (
	"github.com/chaincode_learning_go/chaincode/lib"
)

type Data interface {
	ToBytes() ([]byte, error)

	ToString() string

	ToMap() map[string]string

	GetPrimaryKey() string
}

func GenIndexKeys(data Data, indexNamePackage IndexNamePackage) IndexKeyPackage {
	entityMap := data.ToMap()
	indexKeys := IndexKeyPackage{}
	for _, indexName := range indexNamePackage {
		singleKey := IndexKey{string(indexName.Indicator + ":")}
		for _, singleName := range indexName.Names {
			singleKey = append(singleKey, entityMap[singleName])
		}
		indexKeys = append(indexKeys, singleKey)
	}
	return indexKeys
}

type DataFactory interface {
	ToDataEntity([]byte) (Data, error)
}

type IndexIndicator string

type IndexName struct {
	Indicator IndexIndicator
	Names     []string
}

type IndexNamePackage []IndexName

type IndexKey []string

type IndexKeyPackage []IndexKey

type IndexSearchKey struct {
	Indicator IndexIndicator
	Keys      []string
}

func (indexSearchKey *IndexSearchKey) ToSearchKeyList() IndexKey {
	result := IndexKey{string(indexSearchKey.Indicator + ":")}
	result = append(result, indexSearchKey.Keys...)
	return result
}

type PaginationResponse struct {
	PageSize int32  `json:"page_size"`
	BookMark string `json:"book_mark"`
	Results  []Data `json:"results"`
}

// TODO: test on this
func (paginationResponse *PaginationResponse) ToBytes() []byte {
	result, err := lib.ToJsonBytes(paginationResponse)
	if err != nil {
		return []byte("{}")
	}
	return result
}
