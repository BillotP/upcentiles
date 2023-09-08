FROM docker.io/golang:1.21-alpine as builder
ARG cmd=server
ARG version_path=upcentile/internal/api.VERSION
RUN apk add git upx

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY internal internal
COPY cmd/$cmd/main.go .

ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X '$version_path=$version'" \ 
    -o /build/app -mod=mod \
    /src/main.go
RUN upx --best --lzma /build/app

FROM scratch
WORKDIR /run
COPY --from=builder /build/app /run/app
# Copy ca certificates for external https calls to work
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/run/app"]
