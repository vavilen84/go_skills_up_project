FROM golang:1.18 as modules

COPY app/go.mod app/go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM modules as dev
ARG WORKDIR
WORKDIR $WORKDIR
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
CMD [ "air", "-c", ".air.toml" ]
