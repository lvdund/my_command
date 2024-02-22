- enp0s3: NAT interface

```
$ sudo mn --topo=single,2  (topology of consisting 1 switch and 2 hosts)
$ sudo ifconfig s1 up
$ sudo ovs-vsctl add-port s1 enp0s3(or your NAT interface of virtual machine)
$ ifconfig enp0s3 0
$ dhclient s1 (To get the IP address for s1. Till here VM will get the internet connectivity through OVS.)
mininet> xterm h1
h1> ifconfig h1-eth0 0
h1> dhclient h1-eth0 (This will give the IP address to host h1)
h1> exit
h2> ifconfig h2-eth0 0
h2> dhclient h2-eth0 (This will give the IP address to host h2. h2 will be in the same network as h1).
mininet> h1 ping -c 10 8.8.8.8
mininet> h1 ping h2
mininet> h2 ping h1 (This will update the rules in OVS.)
```