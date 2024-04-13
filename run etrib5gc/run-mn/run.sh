#!/bin/bash

./bin/controller -c ./config/controller.json &
./bin/gateway -c ./config/gateway.json &
./bin/ausf -c ./config/ausf.json &
./bin/pcf -c ./config/pcf.json &
./bin/udm -c ./config/udm.json &
./bin/nsm -c ./config/nsm.json &
./bin/amf -c ./config/amf.json &
./bin/pran -c ./config/pran.json &
./bin/damf -c ./config/damf.json &
./bin/smf -c ./config/smf.json &
./bin/upmf -c ./config/upmf.json &

