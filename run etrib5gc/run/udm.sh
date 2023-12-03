#!/bin/bash

gnome-terminal --tab -- ./run/ausf.sh
./bin/udm -c ./config/udm.json
exec bash