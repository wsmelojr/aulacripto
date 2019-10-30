#!/bin/bash
############################################################################
# Hyperledger Fabric 1.4 LTS Installation Script
# @author: Andre Vieira - Federal University of Rio de Janeiro
# @author: Wilson S. Melo - Inmetro
# @date Feb/2019
#
# IMPORTANT: this script was tested with a Ubuntu 18.04 LTS system
############################################################################

DIR=`dirname $0`
VARFILE="/apps-bin-path.sh"
USER=`whoami`

#install curl
sudo apt-get -y install curl

#install docker
curl -fsSl https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
echo "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -c | awk '{print $2}') stable" | sudo tee --append /etc/apt/sources.list.d/docker.list 1> /dev/null
sudo apt-get update
sudo apt-get -y install docker-ce

#install docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod a+x /usr/local/bin/docker-compose

#install golang
sudo snap install go --classic

#install nodejs
sudo apt-get -y install nodejs

#install npm
sudo apt-get -y install npm
sudo npm install npm@5.6.0 -g

sudo cp $DIR$VARFILE /etc/profile.d/

#install Hyperledger Fabric 1.4 LTS
#third-party images (couchdb, kafka, zookeeper) must be the version 0.4.15
cd ~/
curl -sSL "http://bit.ly/2ysbOFE" | sudo bash -s -- 1.4.0 1.4.0 0.4.15
cd - 1> /dev/null

#customizes the $PATH variable to include the Fabric tools 
PATH="$PATH:$HOME/fabric-samples/bin"

#also modifies $PATH in the .profile, making this change persistent
echo '' >> $HOME/.profile
echo '#adding Fabric tools to the $PATH' >> $HOME/.profile
echo 'PATH="$PATH:$HOME/fabric-samples/bin"' >> $HOME/.profile
#does the same with bashrc file, to be sure that $PATH changes also will happen in desktop mode
echo '' >> $HOME/.bashrc
echo '#adding Fabric tools to the $PATH' >> $HOME/.bashrc
echo 'PATH="$PATH:$HOME/fabric-samples/bin"' >> $HOME/.bashrc

#add current user to the docker group, so she can now to manage containers
sudo usermod -a -G docker $USER

#log the results
echo "Prerequirements was successful installed."
echo "Your system will be reboot now!"
read -p "Press ENTER to continue..."
reboot
