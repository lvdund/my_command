#!/bin/bash

sudo sysctl -w net.ipv4.ip_forward=1
sudo iptables -t nat -A POSTROUTING -o enp0s31f6 -j MASQUERADE
sudo iptables -A FORWARD -p tcp -m tcp --tcp-flags SYN,RST SYN -j TCPMSS --set-mss 1400
sudo systemctl stop ufw
sudo systemctl disable ufw # prevents the firewall to wake up after a OS reboot
sudo iptables -I FORWARD 1 -j ACCEPT

sudo route add default gw 192.168.1.1 enp0s3
