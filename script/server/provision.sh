#!/bin/bash

set -e

if [[ "$(whoami)" != root ]]; then
  echo "ERROR: Only user root can run this script."
  exit 1
fi


if [[ -z "${APP_NAME}" ]]; then
  echo "ERROR: APP_NAME is not set"
  exit 1
fi

if [[ -z "${APP_DOMAIN}" ]]; then
  echo "ERROR: APP_DOMAIN is not set"
  exit 1
fi

if [[ -z "${USER_NAME}" ]]; then
  USER_NAME=$APP_NAME
fi

# https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04

# 
# firewall setup
#
echo " -> setting up firewall"
ufw allow OpenSSH
ufw allow https
ufw enable
ufw status


# 
# user setup
#
echo " -> creating user $USER_NAME"
adduser $USER_NAME
usermod -aG sudo $USER_NAME
rsync --archive --chown=$USER_NAME:$USER_NAME ~/.ssh /home/$APP_NUSER_NAMEAME

# https://linuxize.com/post/how-to-add-user-to-sudoers-in-ubuntu/
echo " -> giving $USER_NAME non-password sudo access to service control"
echo << EOF
$USER_NAME ALL=(ALL) NOPASSWD:/usr/bin/systemctl start $APP_NAME
$USER_NAME ALL=(ALL) NOPASSWD:/usr/bin/systemctl stop $APP_NAME
$USER_NAME ALL=(ALL) NOPASSWD:/usr/bin/systemctl restart $APP_NAME
EOF > /etc/sudoers.d/$USER)NAME

echo " -> creating application deploy/run file structure"
mkdir /home/$USER_NAME/bin
chown $USER_NAME:$USER_NAME home/$USER_NAME/bin
mkdir /home/$USER_NAME/deploy
chown $USER_NAME:$USER_NAME home/$USER_NAME/deploy
mkdir /home/$USER_NAME/db
chown $USER_NAME:$USER_NAME home/$USER_NAME/db
mkdir /home/$USER_NAME/tmp
chown $USER_NAME:$USER_NAME home/$USER_NAME/tmp
mkdir /home/$USER_NAME/Maildir
chown $USER_NAME:$USER_NAME home/$USER_NAME/Maildir

#
# postfix
#
echo " -> setting up postfix"
# https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-postfix-on-ubuntu-20-04
apt update

# this will open a configuration GUI. you'll have to use the article above for instructions
echo "You'll find instructions for how to configure postfix here:"
echo "https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-postfix-on-ubuntu-20-04"
echo "Press any key to continue"
read -n 1 -p "Continue?"

# install with configuration gui
DEBIAN_PRIORITY=low apt install postfix
# set additional config options
postconf -e "home_mailbox= Maildir/"
postconf -e "virtual_alias_maps= hash:/etc/postfix/virtual"
echo "@$APP_DOMAIN $USER_NAME" > /etc/postfix/virtual
postmap /etc/postfix/virtual
systemctl restart postfix
ufw allow Postfix

#
# port access
#
# https://www.baeldung.com/linux/bind-process-privileged-port
echo " -> giving application access to priveleged ports"
touch home/$USER_NAME/bin/$APP_NAME
chmod +x home/$USER_NAME/bin/$APP_NAME
sudo setcap 'CAP_NET_BIND_SERVICE+ep' home/$USER_NAME/bin/$APP_NAME
rm home/$USER_NAME/bin/$APP_NAME


#
# systemctl setup
#
echo " -> creating service for $APP_NAME"
echo <<EOF
Description=$APP_NAME application server

[Service]
Type=simple
ExecStart=/home/$USER_NAME/bin/$APP_NAME\
 -domain=$APP_DOMAIN\
 -mode=release\
 -dbfile=/home/$USER_NAME/db/$APP_NAME.db\
 -smtpaddr=127.0.0.1:25\
 -secret=$(openssl rand -hex 20)
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF > /etc/systemd/system/$APP_NAME.service
systemctl daemon-reload

echo " -> done with initial provisioning. you should be ready to deploy the app now"