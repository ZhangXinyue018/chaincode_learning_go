package book

import (
	"encoding/json"

	"github.com/booksharing/common/domain"
	"github.com/booksharing/common/plugin"
)

type BookLedger struct {
	Prefix string
}

func (al *BookLedger) Create(context domain.Context, name string, book Book) error {
	dataBytes, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return context.Stub.PutState(al.genStateKey(context, name), dataBytes)
}

// todo: change value based on key name
func (al *BookLedger) Update(context domain.Context, name, key string, value interface{}) error {
	// 	m := make(map[string]string)
	// err := json.Unmarshal(input, &m)

	return nil
}

func (al *BookLedger) Delete(context domain.Context, name string) error {
	return context.Stub.DelState(al.genStateKey(context, name))
}

func (al *BookLedger) GetBytes(context domain.Context, name string) ([]byte, error) {
	return context.Stub.GetState(al.genStateKey(context, name))
}

func (al *BookLedger) genStateKey(context domain.Context, name string) string {
	return plugin.GenLedgerKey(al.Prefix, name)
}
