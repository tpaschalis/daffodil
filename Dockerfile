FROM golang:alpine as builder
WORKDIR /go/daffodil
COPY . ./
ARG version=dev

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=$version" -o daffodil-exec ./cmd/daffodil/main.go

FROM scratch
COPY --from=builder /go/daffodil/daffodil-exec .

CMD ["./daffodil-exec"]
