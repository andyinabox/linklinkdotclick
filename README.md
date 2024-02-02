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
sudo systemctl start
# get info
sudo systemctl status
# stop service
sudo systemctl stop
# view logs
sudo journalctl -u linkydink.service
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
 - [ ] Add SSL
 - [ ] Authentication
 - [ ] Ability to edit styles in browser
