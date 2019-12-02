
go build
GODEBUG=netdns=go CORE_CHAINCODE_LOGGING_LEVEL=debug CORE_CHAINCODE_LOGGING_SHIM=info CORE_PEER_ADDRESS=127.0.0.1:17052 CORE_CHAINCODE_ID_NAME=token:1.0 ./token

#remember to run the following if this is the first time
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode install -n token -v 1.0 -p github.com/chaincode_learning_go/chaincode/token -l golang
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode instantiate -o orderer.ordererorg.example.com:7050 -C mychannel -n token -l golang -v 1.0 -c '{"Args":[]}' -P "OR ('Org1.member','Org2.member')"
