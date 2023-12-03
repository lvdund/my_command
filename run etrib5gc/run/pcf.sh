#!/bin/bash

gnome-terminal --tab -- ./run/udm.sh
./bin/pcf -c ./config/pcf.json
exec bash