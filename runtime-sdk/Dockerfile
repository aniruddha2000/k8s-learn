######### Build #########
FROM golang:alpine3.17 AS builder
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o manager

######### Run #########
FROM alpine
WORKDIR /
COPY --from=builder /src/manager .
USER 65532
ENTRYPOINT ["/manager"]
