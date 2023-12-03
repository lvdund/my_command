#!/bin/bash

gnome-terminal --tab -- ./run/pran.sh
./bin/smf -c ./config/smf.json
exec bash