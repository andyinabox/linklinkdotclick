# 🖇 linkydink

This app is contained within a single binary in `bin/linkydink`.

## Configuring

You can see options with `./bin/linkydink -h`.

```
Usage of ./bin/linkydink:
  -dbfile string
        location on sqlite db (default "db/linkydink.db")
  -defaultemail string
        an email for the default user that appears when not logged in (default "linkydink@linkydink.tld")
  -defaultusertitle string
        the default user's site title (default "🖇 my reading list")
  -domain string
        the domain the site is hosted on (linklink.click) (default "linklink.click")
  -mode string
        run mode, use 'release' for production (default "debug")
  -port string
        port to run the webserver on (default "8080")
  -secret string
        secret to use for cookie encryption
  -smtpaddr string
        smtp server (default "127.0.0.1:1025")
```

These options can also be configured as env vars, so `-secret` flag becines `LINKY_SECRET` in the `.env` file.

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