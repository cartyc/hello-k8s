FROM golang:1.12 as build

ENV GO111MODULE=on

WORKDIR /go/src/hello-kubernetes

COPY . .

RUN go get -u -v ./... && \
    go install -v ./...

# Get binary for runtime
FROM gcr.io/distroless/base

COPY --from=build /go/bin/hello-kubernetes .
COPY ./templates ./templates

CMD ["./hello-kubernetes"]

EXPOSE 80