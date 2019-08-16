CC_RUNTIME_LANGUAGE=golang
CC_SRC_PATH=github.com/booksharing

docker exec cli peer chaincode install -n booksharing -v 1.7 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n booksharing -l "$CC_RUNTIME_LANGUAGE" -v 1.7 -c '{"Args":[]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
# sleep 5
# docker exec cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n booksharing -c '{"function":"CreateSchool","Args":["NTU", "a test school"]}'
