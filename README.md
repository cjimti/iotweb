[![GoDoc](https://godoc.org/github.com/cjimti/iotweb?status.svg)](https://godoc.org/github.com/cjimti/iotweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/cjimti/iotweb)](https://goreportcard.com/report/github.com/cjimti/iotweb)


# IoT Web Server

A very small web server written in Go for static sites served on devices
such as the Raspberry Pi. Uses bunyan logging.

## Env

Uses environment variables for configuration. The following
are default values, override where needed.

- `export IOTWEB_BASEPATH=/`
- `export IOTWEB_STATICPATH=www`
- `export IOTWEB_PORT=8080`
- `export IOTWEB_FSAPIPATH=yes`
- `export IOTWEB_FSAPIPATH=fsapi/`

## Try

`docker run -it --rm -p 8080:8080 cjimti/iotweb:1.0.0`

for arm base devices use:

`docker run -it --rm -p 8080:8080 cjimti/iotweb:armhf-1.0.0`


### Development

Uses [goreleaser](https://goreleaser.com):

Install goreleaser with brew (mac):
`brew install goreleaser/tap/goreleaser`

Build without releasing:
`goreleaser --skip-publish --rm-dist --skip-validate`