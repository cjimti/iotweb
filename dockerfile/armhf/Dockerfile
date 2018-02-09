FROM arm32v6/alpine
RUN apk --no-cache add ca-certificates
WORKDIR /
RUN mkdir www
COPY . .
CMD ["/iotweb"]