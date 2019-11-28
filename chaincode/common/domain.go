package common

type Data interface {
	ToBytes() ([]byte, error)

	ToString() string

	ToMap() map[string]string

	GetPrimaryKey() string
}

func GenIndexKeys(data Data, indexNames []IndexName) []IndexKey {
	entityMap := data.ToMap()
	indexKeys := make([]IndexKey, 0)
	for _, indexName := range indexNames {
		singleKey := make([]string, 0)
		for _, singleName := range indexName {
			singleKey = append(singleKey, entityMap[singleName])
		}
		indexKeys = append(indexKeys, singleKey)
	}
	return indexKeys
}

type DataFactory interface {
	ToDataEntity([]byte) (Data, error)
}

type IndexName []string
type IndexKey []string
