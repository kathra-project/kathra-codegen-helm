FROM golang:1-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/kathra-code-generator-manager-server/main.go

FROM alpine/helm:3.0.0
RUN apk --no-cache add ca-certificates bash sed grep gawk
WORKDIR /root/
COPY --from=builder /app/main .
ENV PORT=8080
ENV HOST=0.0.0.0
EXPOSE 8080
ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "./main"] 