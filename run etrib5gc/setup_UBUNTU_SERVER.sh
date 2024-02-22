#!/bin/bash

cd ~
mkdir -p $HOME/Environment/golang/gopath
mkdir -p $HOME/Environment/golang/goroot

# sudo apt update -y
# sudo apt upgrade -y

sudo apt -y install git gcc g++ make autoconf libtool pkg-config libmnl-dev libyaml-dev libsctp-dev lksctp-tools iproute2 gnupg curl net-tools
sudo snap install cmake --classic

wget https://dl.google.com/go/go1.18.10.linux-amd64.tar.gz
sudo tar -C $HOME/Environment/golang/goroot -zxvf go1.18.10.linux-amd64.tar.gz

echo '' >> ~/.bashrc
echo '' >> ~/.bashrc
echo 'export GOPATH=$HOME/Environment/golang/gopath' >> ~/.bashrc
echo 'export GOROOT=$HOME/Environment/golang/goroot/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.bashrc
echo 'export GO111MODULE=auto' >> ~/.bashrc
echo '' >> ~/.bashrc

rm go1.18.10.linux-amd64.tar.gz

cd ~
git clone -b v0.8.2 https://github.com/free5gc/gtp5g.git
cd gtp5g
make
sudo make install

# cd ~
# git clone --recursive -b v3.3.0 -j `nproc` https://github.com/free5gc/free5gc.git
# cd free5gc
# make

# cd ~
# sudo apt remove cmdtest yarn
# curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
# echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
# sudo apt update
# sudo apt install -y yarn
# curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
# source ~/.bashrc
# nvm install v12.22.12

# cd ~/free5gc
# make webconsole

# curl -fsSL https://pgp.mongodb.com/server-4.4.asc | sudo gpg -o /usr/share/keyrings/mongodb-server-4.4.gpg --dearmor
# echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-4.4.gpg ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list
# sudo apt-get update
# sudo apt-get install -y mongodb-org=4.4.25 mongodb-org-server=4.4.25 mongodb-org-shell=4.4.25 mongodb-org-mongos=4.4.25 mongodb-org-tools=4.4.25
# echo "mongodb-org hold" | sudo dpkg --set-selections
# echo "mongodb-org-server hold" | sudo dpkg --set-selections
# echo "mongodb-org-shell hold" | sudo dpkg --set-selections
# echo "mongodb-org-mongos hold" | sudo dpkg --set-selections
# echo "mongodb-org-tools hold" | sudo dpkg --set-selections
# sudo systemctl start mongod
# sudo systemctl enable mongod


cd ~
git clone https://github.com/aligungr/UERANSIM
cd UERANSIM
git checkout 3a96298
# make

cd ~

source .bashrc
