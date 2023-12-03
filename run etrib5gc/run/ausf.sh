#!/bin/bash

gnome-terminal --tab -- ./run/controller.sh
./bin/ausf -c ./config/ausf.json
exec bash