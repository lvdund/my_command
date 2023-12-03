#!/bin/bash

gnome-terminal --tab -- ./run/smf.sh
./bin/upmf -c ./config/upmf.json
exec bash