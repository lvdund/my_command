#!/bin/bash

gnome-terminal --tab -- ./run/pcf.sh
./bin/damf -c ./config/damf.json
exec bash