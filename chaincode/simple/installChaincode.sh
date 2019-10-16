CC_RUNTIME_LANGUAGE=golang
CC_SRC_PATH=github.com/simple

# The reason why we don't use peer container to install chaincode is because peer node has no go env installed
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec -e "CORE_PEER_ADDRESS=peer1.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec -e "CORE_PEER_ADDRESS=peer0.org2.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org2" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org2.example.com/users/Admin@org2.example.com/msp" chaincode.helper.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec -e "CORE_PEER_ADDRESS=peer1.org2.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org2" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org2.example.com/users/Admin@org2.example.com/msp" chaincode.helper.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"

# TODO: not yet success
docker exec -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" -e "CORE_PEER_LOCALMSPID=Org1" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/org1.example.com/users/Admin@org1.example.com/msp" chaincode.helper.example.com peer chaincode instantiate -o orderer.ordererorg.example.com:7050 -C mychannel -n simple -l "$CC_RUNTIME_LANGUAGE" -v 1.0 -c '{"Args":[]}' -P "AND ('Org1.member','Org2.member')"

# Invoke example
# peer chaincode invoke -o orderer.ordererorg.example.com:7050 -C mychannel -n simple -c '{"function": "Ping", "Args":[]}'
# peer chaincode invoke -o orderer.ordererorg.example.com:7050 -C mychannel -n simple -c '{"function": "Create", "Args":["test"]}'