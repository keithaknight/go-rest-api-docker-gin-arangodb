#The standard golang image contains all of the resources to build
#But is very large.  So build on it, then copy the output to the
#final runtime container
FROM golang:latest AS buildContainer
WORKDIR /go/src/app

COPY . .

#flags: -s -w to remove symbol table and debug info
#CGO_ENALBED=0 is required for the code to run properly when copied alpine
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod mod -ldflags "-s -w" -o restapi .

#Now build the runtime container, just a stripped down linux and copy the
#binary to it.
FROM alpine:latest
WORKDIR /app
COPY --from=buildContainer /go/src/app/restapi .

ENV GIN_MODE release

ENV HOST 0.0.0.0
ENV PORT 8080
EXPOSE 8080

CMD ["./restapi"]
