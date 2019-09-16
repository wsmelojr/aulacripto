#!/bin/bash
############################################################################
# The Paillier Experiment - PTB
#
# This script resets and restarts a fresh instance of our Fabric network.
# It is practical to get the environment with a white ledger and run the
# performance tests.
# @author: Wilson S. Melo - Inmetro
# @date Mar/2019
############################################################################

#stops any running instance and remove all the container images
./teardown.sh

#starts fabric containers, selecting only orderer, peer0 and cli0
docker-compose -f docker-compose-ptb.yaml up -d orderer.ptb.de peer0.ptb.de cli0

#creates the ptb-channel and starts the ledger with the genesis block
./startPTB.sh

#install fabmorph chaindode in peer0
docker exec cli0 peer chaincode install -n fabmorph -p github.com/hyperledger/fabric/peer/channel-artifacts/fabmorph -v 1.0

#instantiate fabmorph chaincode in the ptb-channel
docker exec cli0 peer chaincode instantiate -o orderer.ptb.de:7050 -C ptb-channel -n fabmorph -v 1.0 -c '{"Args":[]}'
