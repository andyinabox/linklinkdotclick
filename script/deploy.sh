#!/bin/bash

source $PWD/.scriptenv

set -e


if [[ "${DEPLOY_ENV}" == "production" ]]; then
  SSH_USER=$PRODUCTION_SSH_USER
  SSH_HOST=$PRODUCTION_SSH_HOST
else
  DEPLOY_ENV="staging"
  SSH_USER=$STAGING_SSH_USER
  SSH_HOST=$STAGING_SSH_HOST
fi

REMOTE_PATH_ROOT=/home/$SSH_USER
REMOTE_TEMP_PATH=$REMOTE_PATH_ROOT/tmp
REMOTE_DEPLOY_PATH=$REMOTE_PATH_ROOT/deploy
REMOTE_BIN_PATH=$REMOTE_PATH_ROOT/bin

TIME=$(date +%s)
REMOTE_ARCHIVE=$REMOTE_TEMP_PATH/linkydink-$TIME.tar.gz

echo " -> DEPLOYING to $DEPLOY_ENV"

# copy bin to server
echo " -> copying archive to server $SSH_USER@$SSH_HOST"
scp ./dist/linkydink-linux-amd64.tar.gz $SSH_USER@$SSH_HOST:$REMOTE_ARCHIVE

# update symlink
ssh $SSH_USER@$SSH_HOST /bin/bash << EOF
set -e

echo " -> unzipping release binary"
tar -xvzf $REMOTE_ARCHIVE
mv $REMOTE_PATH_ROOT/linkydink-linux-amd64 $REMOTE_DEPLOY_PATH/linkydink-$TIME

echo " -> stopping server"
sudo systemctl stop linkydink

echo " -> making a backup of the db"
mkdir -p db/backups
cp db/linkydink.db db/backups/linkydink-$TIME.db

echo "copying linkydink-$TIME into ~/bin dir"
# this avoids error if file doesn't exist
rm -f -- $REMOTE_BIN_PATH/linkydink
cp $REMOTE_DEPLOY_PATH/linkydink-$TIME $REMOTE_BIN_PATH/linkydink

echo " -> starting server"
sudo systemctl start linkydink

echo " -> logging deploy"
touch deploy.log
echo "$TIME" >> deploy.log

echo " -> removing archive"
rm -rf $REMOTE_TEMP_PATH/*
EOF
