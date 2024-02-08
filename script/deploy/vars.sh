#!/bin/bash

source $PWD/.scriptenv

export REMOTE_PATH_ROOT=/home/$SSH_USER
export REMOTE_TEMP_PATH=$REMOTE_PATH_ROOT/tmp
export REMOTE_DEPLOY_PATH=$REMOTE_PATH_ROOT/deploy
export REMOTE_BIN_PATH=$REMOTE_PATH_ROOT/bin