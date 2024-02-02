#!/bin/bash
set -e

HOST=linklink.click
TIME=$(date +%s)

# copy bin to server
scp ./dist/linkydink-linux-amd64 andy@$HOST:/home/andy/deploy/linkydink-$TIME

# stop application
ssh -t andy@$HOST sudo systemctl stop linkydink

# update symlink
ssh andy@$HOST /bin/bash << EOF
rm /home/andy/bin/linkydink
ln -s /home/andy/deploy/linkydink-$TIME /home/andy/bin/linkydink
EOF

# start application
ssh -t andy@$HOST sudo systemctl start linkydink
