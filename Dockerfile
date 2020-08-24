FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/nileshkorpad1/simple-blog

# Install our dependencies
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options
RUN go get gopkg.in/go-playground/validator.v9
RUN go get gopkg.in/matryer/respond.v1
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/bson/primitive
RUN go get github.com/stretchr/testify/assert
RUN go get gopkg.in/mgo.v2/dbtest
RUN go get github.com/go-openapi/strfmt

RUN go get github.com/swaggo/swag/cmd/swag
RUN go get github.com/swaggo/http-swagger
RUN go get github.com/alecthomas/template

# Install api binary globally within container
RUN go install github.com/nileshkorpad1/simple-blog

# Set binary as entrypoint
ENTRYPOINT /go/bin/simple-blog

# Expose default port (3000)
EXPOSE 8080
