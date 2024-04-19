import time
from mininet.cli import CLI
from mininet.log import lg, info
from mininet.topolib import TreeNet
from mininet.link import TCIntf
from mininet.util import custom

if __name__ == '__main__':
    lg.setLogLevel( 'info')
    info( "*** Settup\n" )
    intf = custom( TCIntf, bw=10, delay='1ms')
    net = TreeNet( depth=1, fanout=14, intf=intf )
    net.addNAT().configDefault()
    net.start()
    info( "*** Hosts are running and should have internet connectivity\n" )
    # net[ 'h1' ].cmd( 'echo "lvdund" > text.txt' )
    
    net[ 'h1' ].cmd( './bin/controller -c config/controller.json &' )
    net[ 'h2' ].cmd( './bin/gateway -c config/gateway.json &' )
    net[ 'h3' ].cmd( './bin/ausf -c config/ausf.json &' )
    net[ 'h4' ].cmd( './bin/pcf -c config/pcf.json &' )
    net[ 'h5' ].cmd( './bin/nsm -c config/nsm.json &' )
    net[ 'h6' ].cmd( './bin/udm -c config/udm.json &' )
    net[ 'h7' ].cmd( './bin/pran -c config/pran.json &' )
    net[ 'h8' ].cmd( './bin/damf -c config/damf.json &' )
    net[ 'h9' ].cmd( './bin/amf -c config/amf.json &' )
    net[ 'h10' ].cmd( './bin/smf -c config/smf.json &' )
    net[ 'h11' ].cmd( './bin/upmf -c config/upmf.json &' )
    time.sleep(2)
    net[ 'h12' ].cmd( './bin/upf -c config/upf1.json &' )
    net[ 'h13' ].cmd( './run/gnb.sh &' )
    time.sleep(2)
    net[ 'h13' ].cmd( './run/ue.sh &' )
    time.sleep(2)
    
    # CLI( net )
    net.stop()