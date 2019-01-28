# build stage
FROM golang:1.11.3  AS build-env

# Set our workdir to our current service in the gopath
WORKDIR /go/src/IstioMesh/
# Copy the current code into our workdir
COPY . .
ENV GOPATH /go/
RUN go build -o IstioMesh src/Istio/main/main.go

# final stage
FROM ubuntu:bionic

WORKDIR /app
COPY --from=build-env /go/src/IstioMesh/IstioMesh /app/

EXPOSE 8654

CMD ./IstioMesh
