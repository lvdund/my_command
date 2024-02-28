#!/bin/bash

gnome-terminal --tab -- ./run/controller.sh
./bin/gateway -c ./config/gateway.json
exec bash