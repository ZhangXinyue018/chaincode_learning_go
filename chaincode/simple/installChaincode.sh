CC_RUNTIME_LANGUAGE=golang
CC_SRC_PATH=github.com/simple

docker exec peer0.org1.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec peer1.org1.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec peer0.org2.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
docker exec peer1.org2.example.com peer chaincode install -n simple -v 1.0 -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"