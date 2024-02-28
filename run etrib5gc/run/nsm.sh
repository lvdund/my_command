#!/bin/bash

gnome-terminal --tab -- ./run/gateway.sh
./bin/nsm -c ./config/nsm.json
exec bash