#!/bin/bash

source $PWD/script/deploy/vars.sh

set -e

if [[ "${DEPLOY_ENV}" != "production" ]]; then
  SSH_USER=$PRODUCTION_SSH_USER
  SSH_HOST=$PRODUCTION_SSH_HOST
else
  DEPLOY_ENV=staging
  SSH_USER=$STAGING_SSH_USER
  SSH_HOST=$STAGING_SSH_HOST
fi

TIME=$(date +%s)
REMOTE_ARCHIVE=$REMOTE_TEMP_PATH/linkydink-$TIME.tar.gz

# copy bin to server
echo "copying archive to server $SSH_USER@$SSH_HOST"
scp ./dist/linkydink-linux-amd64.tar.gz $SSH_USER@$SSH_HOST:$REMOTE_ARCHIVE

# update symlink
ssh $SSH_USER@$SSH_HOST /bin/bash << EOF
set -e

echo "unzipping release binary"
tar -xvzf $REMOTE_ARCHIVE
mv $REMOTE_PATH_ROOT/linkydink-linux-amd64 $REMOTE_DEPLOY_PATH/linkydink-$TIME

echo "stopping server"
sudo systemctl stop linkydink

echo "copying linkydink-$TIME into ~/bin dir"
# this avoids error if file doesn't exist
rm -f -- $REMOTE_BIN_PATH/linkydink
cp $REMOTE_DEPLOY_PATH/linkydink-$TIME $REMOTE_BIN_PATH/linkydink

echo "starting server"
sudo systemctl start linkydink

echo "removing archive"
rm -rf $REMOTE_TEMP_PATH/*
EOF
