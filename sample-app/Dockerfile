FROM golang:alpine AS base
ADD . /go/src/sample-app
WORKDIR /go/src/sample-app
RUN apk update && apk upgrade && \
    apk add --no-cache bash git
RUN go get gopkg.in/mgo.v2
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sample

FROM alpine
WORKDIR /app
COPY --from=base /go/src/sample-app/sample /app
ENTRYPOINT ./sample
EXPOSE 80
