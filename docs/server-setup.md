# Provisioning

> **Note** this script has only been tested on Ubuntu 23 running on a Digital Ocean Droplet

You should be able to provision the server using `script/server/provision.sh`. This script is meant to be run as `root` on the server after you first spin it up.

First copy the script onto the server, then SSH in as root:

```
localhost:~$ scp ./script/server/provision.sh <user>@<host>:provision.sh
localhost:~$ ssh <user>@<host>
```

Next, as the `root` user, you can run the provisioning script:

```
root@server:~# chmod +x provision.sh
root@server:~# APP_NAME=linkydink APP_DOMAIN=<domain> ./provision.sh
```

You'll be asked for some input, the main ones being:
1. To enter a password and information for the user
2. To configure postfix. For this you'll pretty much use all the defaults, except you'll want to set the account username (in the above example "linkydink" ) as the root and postmaster mail recipient. [Read more about postfix configuration here](https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-postfix-on-ubuntu-20-04)

Finally, you can deploy the app to the new infrastructure:

```
localhost:~$ make deploy
```

**In order for postfix to work correctly, you'll need to make sure you've set up your DNS mail records.** See the bottom of this doc for more info on that.


# Server Setup Notes

These are some notes about what the provisioning script is doing, in case you need to troubleshoot.

## Allow binding to port 80

- [Bind process to a priveleged port](https://www.baeldung.com/linux/bind-process-privileged-port)

```bash
sudo setcap 'CAP_NET_BIND_SERVICE+ep' /path/to/linkydink
```

## Running as a daemon

- [Setting up a custom service](https://www.slingacademy.com/article/ubuntu-how-to-create-a-custom-systemd-service/)


`/etc/systemd/system/linkydink.service`
```
[Unit]
Description=linkydink

[Service]
Type=simple
ExecStart=/home/andy/bin/linkydink\
 --domain=linklink.click\
 --mode=release\
 --dbfile=/home/andy/db/linkydink.db\
 --smtpaddr=127.0.0.1:25
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

Commands

```bash
# reload after changing settings
sudo systemctl daemon-reload
# start the service
sudo systemctl start linkydink
# get info
sudo systemctl status linkydink
# stop service
sudo systemctl stop linkydink
# view logs
sudo journalctl -u linkydink.service
```

Remove sudo password requirement:

 - https://www.digitalocean.com/community/tutorials/how-to-edit-the-sudoers-file
 - https://unix.stackexchange.com/questions/192706/how-could-we-allow-non-root-users-to-control-a-systemd-service
 - https://linuxize.com/post/how-to-add-user-to-sudoers-in-ubuntu/
 - https://help.ubuntu.com/community/Sudoers

Add this to `/etc/sudoers.d/andy`

```
andy ALL=(ALL) NOPASSWD:/usr/bin/systemctl start linkydink
andy ALL=(ALL) NOPASSWD:/usr/bin/systemctl stop linkydink
andy ALL=(ALL) NOPASSWD:/usr/bin/systemctl restart linkydink
```

## Setting up postfix

> **Note:** it's probably going to work better to just use Gmail SMTP if I can do it for free

 - Golang settings for postfix https://gist.github.com/jniltinho/d90034994f29d7d25e59c9e0fe5548d2
 - https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-postfix-on-ubuntu-20-04


Also had to set up some DNS records:

| Record | Domain                | Value             |
|--------|-----------------------|-------------------|
| MX     | linklink.click        | linklink.click 10 |
| TXT    | _dmarc.linklink.click | v=DMARC1; p=none; rua=mailto:dmarc@linklink.click  |
| TXT    | linklink.click        | v=spf1 ip4:161.35.108.49 include:linklink.click -all |


## Avoiding getting sent to spam

- https://blog.codinghorror.com/so-youd-like-to-send-some-email-through-code/
