# IoT Web Server

A very small web server written in Go for static sites served on devices
such as the Raspberry Pi. Uses bunyan logging.

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