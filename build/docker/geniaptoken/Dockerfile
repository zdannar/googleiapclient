#----------------------------
# GOLANG BUILDER CONTAINER
#----------------------------
FROM golang:alpine as builder
WORKDIR /go/src/app
COPY cmd/geniaptoken /go/src/app/src
#COPY internal /go/src/app/src/internal
COPY pkg /go/src/app/src/pkg
COPY go.mod /go/src/app/src
WORKDIR /go/src/app/src
ENV GO111MODULE=on
RUN go build -o main main.go

#----------------------------
# EXEC CONTAINER
#----------------------------
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/src/main /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/main" ]
CMD [ "--help" ]
