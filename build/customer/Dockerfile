FROM golang:1.15

# Set the Current Working Directory inside the container
WORKDIR /customer

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# build the image
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o app cmd/customer/*.go

#-------------------------------------------
# build the output image based on alpine
FROM xushikuan/alpine-build:2.0 AS builder

# set the working directory
WORKDIR /customer

# copy the executable file from the las
COPY --from=0 /customer/app .

# run the service
CMD ["./app"]  

