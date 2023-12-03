#!/bin/bash

gnome-terminal --tab -- ./run/damf.sh
./bin/amf -c ./config/amf.json
exec bash