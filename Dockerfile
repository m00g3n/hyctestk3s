FROM golang:1.16-alpine as builder

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o dice ./cmd/main.go

# result container
FROM scratch

COPY --from=builder /app/dice /dice

EXPOSE 8080

USER 1000:1000

CMD ["/dice"]

