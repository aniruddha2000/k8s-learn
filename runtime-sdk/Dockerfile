FROM golang:alpine3.17 as builder
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o runtime-sdk

FROM alpine
WORKDIR /app
COPY --from=builder /src/runtime-sdk /app/runtime-sdk
ENTRYPOINT ["/app/runtime-sdk"]
