FROM    golang:1.15-alpine as builder
RUN     mkdir /build
ADD     . /build
WORKDIR /build/app
RUN     CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .
FROM    scratch
COPY    --from=builder /build/app/app /
WORKDIR /
CMD     ["./app"]