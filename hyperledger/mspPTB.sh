#/bin/bash
################################################################
# This script creates MSP and channel stuffs to implement 
# a basic PTB blockchain network.
################################################################
# Please check if the following environment var is properly 
# defined (e.g., /etc/profile file). Otherwise, uncomment the 
# following line.
#export FABRIC_CFG_PATH=$PWD

cd "$(dirname "$0")"

# generates MSP artefacts (keys, digital certificates, etc) in the folder ./crypto-config
cryptogen extend --config=./crypto-config-ptb.yaml

# creates the genesis block
configtxgen -profile PTBGenesis -outputBlock ./ptb-genesis.block

# creates the channel config file 
configtxgen -profile PTBChannel -outputCreateChannelTx ./ptb-channel.tx -channelID ptb-channel

# creates anchors config for each organization (in practice, we have only PTB)
configtxgen -profile PTBChannel -outputAnchorPeersUpdate ptb-anchors.tx -channelID ptb-channel -asOrg PTB