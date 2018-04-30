FROM golang:1.10.1-alpine as build

WORKDIR /go/src/hello-kubernetes

COPY . .

# Go dep!
RUN go build -o hello-kubernetes

# # Prod
FROM alpine
WORKDIR /app
COPY --from=build /go/src/hello-kubernetes /app/
RUN apk add --update ca-certificates
ENTRYPOINT [ "./hello-kubernetes" ]
EXPOSE 80
