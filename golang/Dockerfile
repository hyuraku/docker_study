FROM golang:1.14.6-alpine3.12 as builder
WORKDIR /golang/
ADD . ./
RUN go build -o main ./src/form.go

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /golang/main /app/
COPY --from=builder /golang/src/login.gtpl /app/

RUN chmod +x main
# EXPOSE 2224
CMD ["./main"]
