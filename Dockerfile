FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o weather-api $(go list ./... | xargs go list -f '{{if eq .Name "main"}}{{.ImportPath}}{{end}}' | head -n 1)

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/weather-api .

EXPOSE 8080

CMD ["./weather-api"]