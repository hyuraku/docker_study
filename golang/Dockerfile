FROM golang:1.14.6-alpine3.12 as builder
WORKDIR /golang/
ADD . ./
RUN go build -o main ./src/

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /golang/main /app/

RUN chmod +x main
# EXPOSE 2224
CMD ["./main"]
