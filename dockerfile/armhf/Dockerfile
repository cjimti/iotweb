FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY iotweb /iotweb
CMD ["/iotweb"]