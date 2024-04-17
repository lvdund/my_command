#!/bin/bash

# Get the upper limit (n) from the user
read -p "Enter the upper limit (n): " n

# Loop through numbers from 1 to n
for ((i = 1; i <= n; i++)); do
    cls; sudo mn -c; sudo python3 test.py
done