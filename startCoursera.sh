CC_RUNTIME_LANGUAGE=golang
CC_SRC_PATH=github.com/coursera

docker exec cli peer chaincode install -n coursera -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n coursera -l "$CC_RUNTIME_LANGUAGE" -v 1.0 -c '{"Args":[]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
sleep 5
docker exec cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n coursera -c '{"function":"CreateSchool","Args":["NTU", "a test school"]}'
