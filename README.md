# Overview
This is a learning repo for hyperledger fabric chaincode

# How to use
Taking token chaincode as an example, two modes are provided. Developers are recommended to use dev mode.

## Normal Mode
### [Step 1] Deploy basic network
Run ./startNetwork.sh

### [Step 2] Deploy chaincode
1. Go to `chaincode/token`
2. Run script using `installChaincode.sh`

### [Step 3] Run application
1. Go to `app/token`
2. Run app using `node index.js`

### [Step 4] Check swagger
1. Go to browser, type in localhost:3000/api-docs
2. Try to create token, issue token and transfer token

## Dev Mode
Dev mode enable developer to edit chaincode and restart chaincode again.
### [Step 1] Deploy dev network
Run ./startNetworkDev.sh

### [Step 2] Deploy chaincode
1. Go to `chaincode/token`
2. Run script using `installChaincodeDev.sh`
3. Install and instantiate chaincode, run cmd 
```
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode install -n token -v 1.0 -p github.com/chaincode_learning_go/chaincode/token -l golang
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode instantiate -o orderer.ordererorg.example.com:7050 -C mychannel -n token -l golang -v 1.0 -c '{"Args":[]}' -P "OR ('Org1.member','Org2.member')"
```

### [Step 3] Run application
1. Go to `app/token`
2. Run app using `node index.js`

### [Step 4] Check swagger
1. Go to browser, type in localhost:3000/api-docs
2. Try to create token, issue token and transfer token

### [Step 5] Develop
1. Feel free to edit chaincode and restart chaincode again
** Notes: Make sure the chaincode is fully built and started before invoking any chaincode.
