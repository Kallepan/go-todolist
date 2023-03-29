FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/kalle/todolist/
WORKDIR /go/src/github.com/kalle/todolist
RUN go mod download
COPY . /go/src/github.com/kalle/todolist
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/todolist github.com/kalle/todolist

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/kalle/todolist/build/todolist /usr/bin/todolist
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/todolist"]