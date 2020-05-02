FROM golang:alpine as builder
WORKDIR /go/daffodil
COPY . ./
ARG version=dev

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags "-X main.version=$version" -o daffodil-exec ./cmd/daffodil/main.go

FROM scratch
COPY --from=builder /go/daffodil/daffodil-exec .
# COPY --from=builder /go/daffodil/.env.dist .env
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./daffodil-exec"]
