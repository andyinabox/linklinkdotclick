# Server Setup

These notes are based on a DigitalOcean Ubuntu 23 droplet

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
