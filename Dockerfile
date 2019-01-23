# build stage
FROM golang:1.11.3  AS build-env

# Set our workdir to our current service in the gopath
WORKDIR /go/src/IstioMesh/
# Copy the current code into our workdir
COPY . .
ENV GOPATH /go/
RUN go build -o IstioMesh main/main.go

# final stage
FROM ubuntu:bionic

WORKDIR /app
COPY --from=build-env /go/src/IstioMesh/service-engine /app/

EXPOSE 8090

CMD ./IstioMesh
