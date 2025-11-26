FROM golang:1.25-alpine AS go-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM node:18-alpine AS react-builder

WORKDIR /app/star-catalog
COPY star-catalog/package*.json ./
RUN npm ci
COPY star-catalog/ ./
RUN npm run build 

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root

COPY --from=go-builder /app/main .
COPY --from=react-builder /app/star-catalog/build ./web

ENV STATIC_DIR=/root/web

EXPOSE 1323
CMD ["./main"]