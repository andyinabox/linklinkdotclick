#!/bin/bash
set -e

HOST=linklink.click
TIME=$(date +%s)

# copy bin to server
scp ./bin/linkydink-linux-amd64 andy@$HOST:/home/andy/deploy/linkydink-$TIME

# stop application
# ssh -t andy@$HOST sudo systemctl stop linkydink

# update symlink
ssh andy@$HOST /bin/bash << EOF
echo "stopping server"
sudo systemctl stop linkydink

echo "replacing symlink for linkydink-$TIME"
rm /home/andy/bin/linkydink
ln -s /home/andy/deploy/linkydink-$TIME /home/andy/bin/linkydink

echo "starting server"
sudo systemctl start linkydink
EOF

# start application
# ssh -t andy@$HOST sudo systemctl start linkydink
