# 🖇 linkydink

This app is contained within a single binary in `bin/linkydink`. You can see options with `./bin/linkydink -h`

```bash
Usage of ./bin/linkydink:
  -dbfile string
    	location on sqlite db (default "db/linkydink.db")
  -defaultemail string
    	an email for the default user that appears when not logged in (default "linkydink@linkydink.tld")
  -domain string
    	the domain the site is hosted on (linklink.click) (default "linklink.click")
  -mode string
    	run mode, use 'release' for production (default "debug")
  -port string
    	port to run the webserver on (default "8080")
  -smtpaddr string
    	smtp server (default "127.0.0.1:1025")
```

## Build commands

A few usefull commands. See the rest in the [`Makefile`](./Makefile)

```bash
# build single binary into dist/
make

# build and run server
make run

# run server, restart when files change
make watch
```

## Additional Documentation

 - [Server Setup](./docs/server-setup.md)