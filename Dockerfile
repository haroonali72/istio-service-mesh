# build stage
FROM golang:1.12.17  AS build-env

# Set our workdir to our current service in the gopath
WORKDIR /go/src/bitbucket.org/cloudplex-devs/istio-service-mesh/
# Copy the current code into our workdir
COPY . .
ENV GOPATH /go/
RUN CGO_ENABLED=0 go build -o IstioMesh main.go

# final stage
FROM ubuntu:bionic

WORKDIR /app
COPY --from=build-env /go/src/bitbucket.org/cloudplex-devs/istio-service-mesh/IstioMesh /app/

EXPOSE 8654

CMD ./IstioMesh
