FROM hyperjiang/golang:1.14 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -o /app/main

FROM gcr.io/distroless/base
COPY --from=builder /app/main /
COPY --from=builder /app/configs /configs

ENTRYPOINT ["/main"]
