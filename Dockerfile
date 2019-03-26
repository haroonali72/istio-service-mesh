# build stage
FROM golang:1.11.3  AS build-env

# Set our workdir to our current service in the gopath
WORKDIR /go/src/istio-service-mesh/
# Copy the current code into our workdir
COPY . .
ENV GOPATH /go/
RUN go build -o IstioMesh main/main.go

# final stage
FROM ubuntu:bionic

WORKDIR /app
COPY --from=build-env /go/src/istio-service-mesh/IstioMesh /app/

EXPOSE 8654

CMD ./IstioMesh
