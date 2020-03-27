[![Build Status](https://dev.azure.com/kaknight47/kaknight47/_apis/build/status/keithaknight.go-rest-api-docker-gin-arangodb?branchName=master)](https://dev.azure.com/kaknight47/kaknight47/_build/latest?definitionId=1&branchName=master)

# go-rest-api-docker-gin-arangodb
A simple Go (Golang) REST API starter project running in Docker with Gin and ArangoDB

Note: this will return HTTP status code 500 (Internal Server Error) until you
have configured the ArangoDB settings.

As coded, this has routes for users and companies.  Each route expects an
ArangoDB collection with the same name.

The Dockerfile uses a straged approach.  First, it creates a build container.
Then, the output is copied from the build container to the final execution
container.  This drastically reduces the size of the production container (from
988MB to 18MB).

## Configuration
Configuration can be controlled through environment variables
or through a configuration object in arangoConfig.go

#### Web Server Environment Variables:
- HOST (Default: 127.0.0.1 local; 0.0.0.0 in Docker)
- PORT (Default: 8080)

#### ArangoDB Environment Variables:
- ARANGODB_URLS (comma separated urls of database servers, including protocol and port)
- ARANGODB_USER (database user name)
- ARANGODB_PASSWORD (database user password)
- ARANGODB_DATABASE (database name)

## Build
#### Local:
`$ go build .`

#### Docker
`$ docker build .`

## Run Tests
`$ go test`

## Running
#### Local:
`$ ./api`

#### Docker:
After building the docker image, it will say: 'Successfully built {IMAGE HASH}'.
Use the image hash to start the container.  For legibility, you can
assign a tag to the image and reference it by tag name instead.

`$ docker run -p 8080:8080 {IMAGE HASH}`

## Calling:
`$ curl http://localhost:8080/users/{user id}`
