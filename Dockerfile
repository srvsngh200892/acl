FROM golang:1.10

MAINTAINER Saurav Singh "srvsngh200892@gmail.com"

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
COPY . /go/src/github.com/srvsngh200892/acl

WORKDIR /go/src/github.com/srvsngh200892/acl

RUN dep ensure --vendor-only

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o acl acl.go

RUN chmod a+x ./run.sh

EXPOSE  8080

CMD ["./run.sh"]