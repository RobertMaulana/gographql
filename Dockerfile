FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/graphql-go
COPY . .
RUN go get
RUN go build -o graphql-go
ENTRYPOINT ./graphql-go
EXPOSE 8080