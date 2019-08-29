#!/bin/bash

configtxgen -profile GenesisProfile -channelID orderer-system-channel -outputBlock ./config/genesis.block

configtxgen -profile MyChannelProfile -outputCreateChannelTx ./config/channel.tx -channelID mychannel

configtxgen -profile MyChannelProfile -outputAnchorPeersUpdate ./config/org1anchors.tx -channelID mychannel -asOrg Org1

configtxgen -profile MyChannelProfile -outputAnchorPeersUpdate ./config/org2anchors.tx -channelID mychannel -asOrg Org2
