from mininet.cli import CLI
from mininet.log import lg, info
from mininet.net import Mininet
from mininet.node import Node
from mininet.topo import Topo
from mininet.link import TCLink, TCIntf
from mininet.util import custom

class CustomTopo(Topo):
    "Simple topology example."

    def build(self):
        h1 = self.addHost('h1')
        h2 = self.addHost('h2')
        s1 = self.addSwitch('s1')
        self.addLink(h1, s1, delay='10ms')
        self.addLink(h2, s1)

if __name__ == '__main__':
    lg.setLogLevel('info')
    info("*** Setup\n")
    topo = CustomTopo()
    intf = custom(TCIntf, bw=10)
    net = Mininet(topo=topo, intf=intf, link=TCLink)
    
    # Add NAT connectivity
    nat = net.addNAT().configDefault()

    net.start()
    info("*** Hosts are running and should have internet connectivity\n")
    # net['h1'].cmd('echo "lvdund" > text.txt')
    CLI(net)
    # Shut down NAT
    net.stop()
