FROM golang:alpine AS intermediate

RUN apk update && \
    apk add --no-cache git make

RUN adduser -D -g '' gonverter

WORKDIR $GOPATH/src/

COPY . .

RUN go mod download
RUN go mod verify
RUN make build

FROM scratch

COPY --from=intermediate /go/src/bin/gonverter /go/bin/gonverter
COPY --from=intermediate /etc/passwd /etc/passwd

USER gonverter

WORKDIR /go/bin

CMD ["/go/bin/gonverter"]