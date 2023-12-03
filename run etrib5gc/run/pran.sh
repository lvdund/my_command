#!/bin/bash

gnome-terminal --tab -- ./run/amf.sh
./bin/pran -c ./config/pran.json
exec bash