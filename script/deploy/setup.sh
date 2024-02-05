#!/bin/bash

source $PWD/script/deploy/vars.sh

echo "setting up server for deploy"
ssh $SSH_USER@$SSH_HOST /bin/bash << EOF
set -e
mkdir -p $REMOTE_TEMP_PATH
mkdir -p $REMOTE_DEPLOY_PATH
mkdir -p $REMOTE_BIN_PATH
EOF
