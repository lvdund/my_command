# from __future__ import print_function

import os
from mininet.topo import Topo
from mininet.net import Mininet
from mininet.node import Node
from mininet.log import setLogLevel, info
from mininet.cli import CLI
from mininet.link import Intf
from mininet.node import Controller
import time

class NetworkTopo( Topo ):
    # Builds network topology
    def build( self, **_opts ):

        s1 = self.addSwitch ( 's1', failMode='standalone' )

        # Adding hosts
        ctl = self.addHost( 'ctl', ip='192.168.60.1/28' )
        gw = self.addHost( 'gw', ip='192.168.60.2/28' )
        ausf = self.addHost( 'ausf', ip='192.168.60.3/28' )
        pcf = self.addHost( 'pcf', ip='192.168.60.4/28' )
        nsm = self.addHost( 'nsm', ip='192.168.60.5/28' )
        udm = self.addHost( 'udm', ip='192.168.60.6/28' )
        pran = self.addHost( 'pran', ip='192.168.60.7/28' )
        damf = self.addHost( 'damf', ip='192.168.60.8/28' )
        amf = self.addHost( 'amf', ip='192.168.60.9/28' )
        smf = self.addHost( 'smf', ip='192.168.60.10/28' )
        upmf = self.addHost( 'upmf', ip='192.168.60.11/28' )
        upf1 = self.addHost( 'upf1', ip='192.168.60.12/28' )
        upf2 = self.addHost( 'upf2', ip='192.168.60.13/28' )
        upf3 = self.addHost( 'upf3', ip='192.168.60.14/28' )
        gnb = self.addHost( 'gnb', ip='192.168.60.15/28' )
        ue = self.addHost( 'ue', ip='192.168.60.16/28' )
        
        # Connecting hosts to switches
        for d, s in [ (ctl, s1), (gw, s1), (ausf, s1), (pcf, s1), (nsm, s1), (udm, s1), (pran, s1), (damf, s1) ]:
            self.addLink( d, s )
        for d, s in [ (amf, s1), (smf, s1), (upmf, s1), (upf1, s1), (upf2, s1), (upf3, s1), (gnb, s1), (ue, s1) ]:
            self.addLink( d, s )


def run():

    topo = NetworkTopo()
    
    net = Mininet( topo=topo, controller=None )
    net.start()
    # Make Switch act like a normal switch
    #net['s1'].cmd('ovs-ofctl add-flow s1 action=normal')
    # Make Switch act like a hub
    #net['s1'].cmd('ovs-ofctl add-flow s1 action=flood')
    
    net[ 'ctl' ].cmd( './bin/controller -c config/controller.json &' )
    time.sleep(1)
    net[ 'gw' ].cmd( './bin/gateway -c config/gateway.json &' )
    time.sleep(1)
    net[ 'ausf' ].cmd( './bin/ausf -c config/ausf.json &' )
    time.sleep(1)
    net[ 'pcf' ].cmd( './bin/pcf -c config/pcf.json &' )
    time.sleep(1)
    net[ 'nsm' ].cmd( './bin/nsm -c config/nsm.json &' )
    time.sleep(1)
    net[ 'udm' ].cmd( './bin/udm -c config/udm.json &' )
    time.sleep(1)
    net[ 'pran' ].cmd( './bin/pran -c config/pran.json &' )
    time.sleep(1)
    net[ 'damf' ].cmd( './bin/damf -c config/damf.json &' )
    time.sleep(1)
    net[ 'amf' ].cmd( './bin/amf -c config/amf.json &' )
    time.sleep(1)
    net[ 'smf' ].cmd( './bin/smf -c config/smf.json &' )
    time.sleep(1)
    net[ 'upmf' ].cmd( './bin/upmf -c config/upmf.json &' )
    time.sleep(1)
    net[ 'upf1' ].cmd( './bin/upf -c config/upf1.json &' )
    # time.sleep(1)
    # net[ 'uf2' ].cmd( './bin/uf2 -c config/uf2.json &' )
    # time.sleep(1)
    # net[ 'upf3' ].cmd( './bin/upf3 -c config/upf3.json &' )
    # time.sleep(1)
    # net[ 'gnb' ].cmd( './config/ueransim/nr-gnb -c ./config/free5gc-gnb.yaml &' )
    # time.sleep(1)
    # net[ 'ue' ].cmd( './config/ueransim/nr-ue -c ./config/free5gc-ue.yaml &' )
    
    CLI( net )
    net.stop()

if __name__ == '__main__':
    setLogLevel( 'info' )
    run()
