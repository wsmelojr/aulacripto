#!/bin/bash
############################################################################
# The Paillier Experiment - PTB
#
# This script collects statistics related to CPU usage of until 3 containers.
# They usually are the chaincode container (that is created in runtime and 
# usually has a composed name with the chaincode name), the endorser container
# (where the chaincode was installed) and the couchdb container associated with
# the endorser.
#
# The script output is a table with 3 collumns showing the CPU usage curve.
#
# @author: Wilson S. Melo - Inmetro
# @date Mar/2019
############################################################################

#test if we have the paramaters
#if [ "$#" -ne 3 ]; then
#  echo "Usage: $0 <chaincode container ID> <endorser container ID> <couchdb container ID>" >&2
#  exit 1
#fi

#user awk and tr to process the log of the peer0.ptb.de
echo -n "Total of blocks: "
docker logs peer0.ptb.de 2>&1 | grep Committed | awk '{print $12}' | tail -n1

echo -n "Total of transactions: "
docker logs peer0.ptb.de 2>&1 | grep Committed | awk '{print $12 " " $14}' | awk '{ sum += $2 } END { print sum }'

echo -n "Total of MVCC: "
docker logs peer0.ptb.de 2>&1 | grep MVCC_READ_CONFLICT | wc -l
