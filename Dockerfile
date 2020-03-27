#The standard golang image contains all of the resources to build
#But is very large.  So build on it, then copy the output to the
#final runtime container
FROM golang:latest AS buildContainer
WORKDIR /go/src/app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/arangodb/go-driver

COPY *.go ./

#flags: -s -w to remove debug info and symbol table
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o api .

#Now build the runtime container, just a stripped down linux and copy the
#binary to it.
FROM alpine:latest
WORKDIR /app
COPY --from=buildContainer /go/src/app .

ENV GIN_MODE release

ENV HOST 0.0.0.0
ENV PORT 8080
EXPOSE 8080

CMD ["./api"]
