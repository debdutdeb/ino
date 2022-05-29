FROM golang:1.18.2-alpine3.16

COPY src /mnt

RUN go build -o /mnt/ino /mnt/src/main.go

FROM alpine:3.16.0

COPY --from=0 /mnt/ino /bin/ino

ENTRYPOINT [ "ino" ]