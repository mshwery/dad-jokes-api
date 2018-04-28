FROM golang:1.10.1

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# deps
RUN go get github.com/golang/dep/cmd/dep
#RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/app

# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

# install packages
RUN dep ensure --vendor-only

# app files
ADD *.go *.json ./

# build the app
RUN go build -o /app/main .

ARG PORT=8080
ENV PORT $PORT
EXPOSE $PORT

# check every 30s to ensure this service returns HTTP 200
HEALTHCHECK CMD curl -fs http://localhost:$PORT/health || exit 1

CMD ["/app/main"]