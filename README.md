# Example Go service

[![Build Status](https://travis-ci.org/powerman/go-service-narada-example.svg?branch=master)](https://travis-ci.org/powerman/go-service-narada-example)
[![Build Status](https://circleci.com/gh/powerman/go-service-narada-example.svg?style=shield)](https://circleci.com/gh/powerman/go-service-narada-example)

## INSTALL

To install this project **for development**:

- install external dependencies
  - setup (in home dir or system-wide) Perl-module Narada (for ex.: `cpanm
    Narada`)
  - install OS packages `runit` и `socklog`
- clone the repo into `$(go env GOPATH)/src/github.com/powerman/go-service-narada-example/`
- build and deploy project (these commands should be run in project
  directory)
  - install dependencies needed to build project `go get -t ./...`
  - run `./release && ./deploy -n`
    - this will install current version of project into subdirectory
      `_live/`
    - user's cron will be updated with command to start project's services
      (they'll run in about a minute)
- setup project (these commands should be run in `_live/` project's
  subdirectory)
  - setup some config files:
    - `config/listen` (should contain value suitable for
      `http.ListenAndServe()`, for ex. ":8080")
  - restart main service, so it'll notice updated config files: `sv t
    ./service/main/`
- setup nginx as per example in `doc/nginx-example/`
  - web-service's port in nginx config for website should match
    `config/listen`
  - pay attention to `ssl_dhparam` and `resolver` settings in nginx.conf -
    file should exists (command to generate it shows in nginx.conf
    comments), IP should be same as in `/etc/resolv.conf`
  - setup your local IP addresses in
    `conf/directive-only/protect-local-sites.conf` (this config should be
    disabled on server)
  - symlink `/var/www/api.example/main` to
    `_live/public/` project's subdirectory (path
    `/var/www/api.example` must match `root /var/www/$host;`
    setting in nginx config)
    - NOTE: at the moment this isn't used, so you can delay this setup
      until we'll need to serve static files by HTTP

See more documentation about project's development flow at
[Narada framework](https://metacpan.org/pod/Narada) and in
`doc/narada-base.adoc`.

Installation on server differs from described above (there is no repo with
project's sources, server has only contents of `_live/`). To deploy
some (current by default) version on server run:

```sh
./deploy_server [<version>]
```

## Server

Project is installed on server.

HTTP location: https://api.example/main/


## Development

After making some changes you have to restart main service: `sv t
./service/main/` (this will be automatically done by `./deploy` if it
makes any changes in `_live/`).
