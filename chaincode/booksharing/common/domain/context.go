package domain

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Context struct {
	Stub        shim.ChaincodeStubInterface
	Input       []byte
	AccountName string
	Sig         []byte
	Data        map[string]interface{}
}
