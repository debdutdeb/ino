FROM golang:1.18.2-alpine3.16

COPY src /mnt
COPY go.mod /mnt
COPY go.sum /mnt

RUN cd /mnt && go build -o ino

FROM alpine:3.16.0

COPY --from=0 /mnt/ino /bin/ino
COPY entrypoint.sh /mnt/entrypoint.sh

ENTRYPOINT ["/bin/sh",  "/mnt/entrypoint.sh" ]