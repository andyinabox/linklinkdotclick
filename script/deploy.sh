#!/bin/bash
set -e

IP=`doctl compute droplet get linklink.click --template {{.PublicIPv4}}`
ssh -t andy@$IP sudo systemctl stop linkydink
scp ./dist/linkydink-linux-amd64 andy@$IP:/home/andy/bin/linkydink
ssh -t andy@$IP sudo systemctl start linkydink
