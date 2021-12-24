FROM golang:1.15


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

FROM xushikuan/alpine-build:2.0 AS builder

ENV WORK_DIR=$GOPATH/src/github.com/jolinGalal/jumia/cmd/customer

WORKDIR $WORK_DIR

COPY . .

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main *.go

# copy the executable file from the las
COPY --from=0 /customer/app .

# run the service
CMD [./app]
