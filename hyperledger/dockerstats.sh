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
if [ "$#" -ne 3 ]; then
  echo "Usage: $0 <chaincode container ID> <endorser container ID> <couchdb container ID>" >&2
  exit 1
fi

#remove any previous cpu.stats file
rm -f cpu.stats

#user awk and tr to process the output of the command docker stats. The command stays
#in a continuous loop until receives a SIGKILL (or CTRL+C)
while true; do docker stats $1 $2 $3 --no-stream | awk '{print $3}' | tr '\n' ' ' | tr -d 'NAME' >> cpu.stats; echo >> cpu.stats ; done