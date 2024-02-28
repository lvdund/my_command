#!/bin/bash

gnome-terminal --tab -- ./run/nsm.sh
./bin/ausf -c ./config/ausf.json
exec bash