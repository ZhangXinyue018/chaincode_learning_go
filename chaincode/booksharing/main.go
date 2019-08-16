package main

import (
	"fmt"

	"github.com/booksharing/common/domain"
	dp "github.com/booksharing/common/domain_with_plugin"
	"github.com/booksharing/common/plugin"
	"github.com/booksharing/feature"
	"github.com/booksharing/feature/account"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("main")

type Chaincode struct {
	ServicePack feature.ServicePack
	PluginPack  plugin.PluginPack
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	APIstub.PutState("version", []byte("1.0.0"))
	return shim.Success(nil)
}

func (cc *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	function, args := APIstub.GetFunctionAndParameters()

	context := domain.Context{
		Stub:        APIstub,
		Input:       []byte(args[0]),
		AccountName: args[1],
		Sig:         []byte(args[2]),
	}
	contractContext := dp.ContractContext{
		DataContext: &context,
		Plugins:     cc.PluginPack,
	}
	return cc.ServicePack.Invoke(function, contractContext)
}

func main() {
	chaincode := Chaincode{
		ServicePack: feature.ServicePack{
			AccountService: account.AccountService{},
		},
		PluginPack: plugin.PluginPack{
			IdentityVerifier:    plugin.IdentityVerifier{},
			APIDataProcessor:    plugin.JsonProcessor{},
			LedgerDataProcessor: plugin.JsonProcessor{},
		},
	}
	err := shim.Start(&chaincode)
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
