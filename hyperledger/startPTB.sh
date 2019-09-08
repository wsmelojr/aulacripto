##########################################################################
# FABPUMP EXPERIMENT - ptb - June/2018
# This script requires a pre configured docker environment, with the fours
# peers and orderer associate the Org ptb already started.
# It uses peer0 for creating a channel and after and joins it. After that,
# it fetches by the channel with the other peers and also joins each peer
# to the channel
# Author: Wilson S. Melo Jr.
##########################################################################
#!/bin/bash
# Define auxiliar vars
CHANNEL=ptb-channel
ORDERER=orderer.ptb.de:7050
DOMAIN=ptb.de

# Exit on first error.
set -e

#Detect architecture
ARCH=`uname -m`

# Grab the current directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Create the channel
docker exec peer0.$DOMAIN peer channel create -o $ORDERER -c $CHANNEL -f /etc/hyperledger/configtx/$CHANNEL.tx

# Join peer0.org1.example.com to the channel.
docker exec peer0.$DOMAIN peer channel join -b $CHANNEL.block

## Fetch the channel in any one of the other peers
docker exec peer1.$DOMAIN peer channel fetch config -o $ORDERER -c $CHANNEL
#docker exec peer2.$DOMAIN peer channel fetch config -o $ORDERER -c $CHANNEL
#docker exec peer3.$DOMAIN peer channel fetch config -o $ORDERER -c $CHANNEL

## Join each one of the peers to the channel.
docker exec peer1.$DOMAIN peer channel join -b ${CHANNEL}_config.block
#docker exec peer2.$DOMAIN peer channel join -b ${CHANNEL}_config.block
#docker exec peer3.$DOMAIN peer channel join -b ${CHANNEL}_config.block

# This command should be used to update the anchor, however the blockchain is working without it
# and it is not clear when such command is indeed necessary.
#docker exec peer0.ptb.de peer channel update -o orderer.ptb.de:7050 -c ptb-channel -f /etc/hyperledger/configtx/ptb-anchors.tx

# Direct commands (just in case you need a fast copy&paste)
# # Create the channel
# docker exec peer0.ptb.de peer channel create -o orderer.ptb.de:7050 -c ptb-channel -f /etc/hyperledger/configtx/ptb-channel.tx

# # Join peer0.org1.example.com to the channel.
# docker exec peer0.ptb.de peer channel join -b ptb-channel.block

# # For each one of the other peers, fetches the created channel
# docker exec peer1.ptb.de peer channel fetch config -o orderer.ptb.de:7050 -c ptb-channel
# docker exec peer2.ptb.de peer channel fetch config -o orderer.ptb.de:7050 -c ptb-channel
# docker exec peer3.ptb.de peer channel fetch config -o orderer.ptb.de:7050 -c ptb-channel

# # Join all the other peers to the channel.
# docker exec peer1.ptb.de peer channel join -b ptb-channel_config.block
# docker exec peer2.ptb.de peer channel join -b ptb-channel_config.block
# docker exec peer3.ptb.de peer channel join -b ptb-channel_config.block
