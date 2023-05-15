FROM golang:alpine3.17 as builder
WORKDIR /src
COPY . .
RUN go build -o runtime-sdk

FROM alpine
WORKDIR /app
COPY --from=builder /src/runtime-sdk /app/runtime-sdk
ENTRYPOINT ["/app/runtime-sdk"]
