# Builder image
FROM golang:1.16 as builder

ARG WORKDIR
WORKDIR $WORKDIR

RUN git clone https://github.com/golang-migrate/migrate && \
    cd migrate && \
    make build

# Runner image
FROM alpine:3.13

USER 1000

ARG WORKDIR
WORKDIR $WORKDIR

COPY --from=builder $WORKDIR/migrate/migrate  ./
