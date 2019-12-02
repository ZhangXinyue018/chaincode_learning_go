package common

type Data interface {
	ToStoreBytes() ([]byte, error)

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
	ToDataEntityFromStoredBytes([]byte) (Data, error)
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
