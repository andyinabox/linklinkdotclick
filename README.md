# 🖇 linkydink

This app is contained within a single binary in `dist/linkydink`

```bash
./dist/linkydink --port=8000 --host=127.0.0.1 --dbfile=db/linkydink.db
```

## Build commands

```bash
# build single binary into dist/
make

# build and run server
make run

# run server, restart when files change
make watch
```

## Server Setup

### Allow binding to port 80

-> [Bind process to a priveleged port](https://www.baeldung.com/linux/bind-process-privileged-port)

```bash
sudo setcap 'CAP_NET_BIND_SERVICE+ep' /path/to/linkydink
```

### Running as a daemon

-> [Setting up a custom service](https://www.slingacademy.com/article/ubuntu-how-to-create-a-custom-systemd-service/)


`/etc/systemd/system/linkydink.service`
```
[Unit]
Description=linkydink

[Service]
Type=simple
ExecStart=/home/andy/bin/linkydink --port=80 --mode=release --dbfile=/home/andy/db/linkydink>
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

Commands

```bash
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

Add this to `sudoers`

```
andy ALL=(ALL) NOPASSWD:/usr/bin/systemctl start linkydink
andy ALL=(ALL) NOPASSWD:/usr/bin/systemctl stop linkydink
andy ALL=(ALL) NOPASSWD:/usr/bin/systemctl restart linkydink
```


## Todo

 - [x] Implement API endpoints with ~~test data~~ sqlite
   -  [x] `GET /api/links`
   -  [x] `POST /api/links`
   -  [x] `GET /api/links/{id}`
   -  [x] `DELETE /api/links/{id}`
   -  [x] `PUT /api/links/{id}` 
   -  [ ] ~~`PATCH /api/links/{id}`~~
 - [x] Add persistence (sqlite?)
   - [x] Create `linkrepository`
 - [ ] Functionality
   - [x] Fetching updated data for links
   - [x] Adding links
   - [x] Update `LastClicked` after click
   - [x] Re-ordering links on frontend
   - [x] Deleting links
   - [x] Editing link title
   - [ ] Ability to have non-RSS links
   - [ ] Avoid duplicates
 - [ ] Containerize
 - [x] Deploy
 - [x] Add SSL
 - [x] Setup CORS
 - [ ] Multi-user
 - [ ] OPML import
 - [ ] Authentication
 - [ ] Ability to edit styles in browser
