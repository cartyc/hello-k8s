FROM golang:1.9-alpine

WORKDIR /go/src/hello-kubernetes

COPY . .

RUN go install -v ./...

CMD ["hello-kubernetes"]

EXPOSE 80