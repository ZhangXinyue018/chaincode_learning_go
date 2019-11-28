set -ev


# This script will bring up the network, create a channel called mychannel and join all peers to the channel

docker-compose -f orderer-ca.yml -f peer-dev.yml down

docker-compose -f orderer-ca.yml -f peer-dev.yml up -d ca.example.com orderer.ordererorg.example.com peer0.org1.couchdb peer0.org1.example.com chaincode.helper.example.com
docker ps -a

export FABRIC_START_TIMEOUT=10
sleep ${FABRIC_START_TIMEOUT}

# ==================== create and join channel ===================
# Create the channel -- using a client msp config
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer.ordererorg.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/channel.tx

# Join peer0.org1.example.com to the channel. -- using a client msp config
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel fetch config mychannel.block -c mychannel -o orderer.ordererorg.example.com:7050
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b mychannel.block

# Join peer1.org1.example.com to the channel. -- using a client msp config
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer1.org1.example.com peer channel fetch config mychannel.block -c mychannel -o orderer.ordererorg.example.com:7050
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer1.org1.example.com peer channel join -b mychannel.block

# Join peer0.org2.example.com to the channel. -- using a client msp config
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel fetch config mychannel.block -c mychannel -o orderer.ordererorg.example.com:7050
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel join -b mychannel.block

# Join peer1.org2.example.com to the channel. -- using a client msp config
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.example.com/msp" peer1.org2.example.com peer channel fetch config mychannel.block -c mychannel -o orderer.ordererorg.example.com:7050
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.example.com/msp" peer1.org2.example.com peer channel join -b mychannel.block


# ==================== update anchor peer ===================
# TODO: explore on anchor peer stuff
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel update -o orderer.ordererorg.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/org1anchors.tx
#docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel update -o orderer.ordererorg.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/org2anchors.tx