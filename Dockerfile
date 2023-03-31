FROM golang:1.20-bullseye as builder
COPY go.mod go.sum /go/src/github.com/kalle/todolist/
WORKDIR /go/src/github.com/kalle/todolist
RUN go mod download
COPY . /go/src/github.com/kalle/todolist
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/todolist .

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/kalle/todolist/build/todolist /usr/bin/todolist
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/todolist"]