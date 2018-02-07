# IoT Web Server

A very small web server written in Go for static sites served on devices
such as the Raspberry Pi. Uses bunyan logging.

## Env

Uses environment variables for configuration with the following
default values.

- export **IOTWEB_BASEPATH**=`/`
- export **IOTWEB_STATICPATH**=`www`
- export **IOTWEB_PORT**=`8080`

## Try

`docker run -it --rm -p 8080:8080 cjimti/iotweb:0.1.3`

for arm base devices use:

`docker run -it --rm -p 8080:8080 cjimti/iotweb:armhf-0.1.3`


### Development

Uses [goreleaser](https://goreleaser.com):

Install goreleaser with brew (mac):
`brew install goreleaser/tap/goreleaser`

Build without releasing:
`goreleaser --skip-publish --rm-dist --skip-validate`