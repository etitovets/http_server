FROM golang:1.25 as builder
ARG APP_VERSION=0.0.0
ARG DATE=01.01.01
ARG COMMIT=000
WORKDIR /app
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /app/http-server -ldflags  "-X main.buildVersion=${APP_VERSION} -X 'main.buildDate=${DATE}' -X 'main.buildCommit=${COMMIT}'" cmd/*.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/http-server /app/http-server
EXPOSE 8081
ENTRYPOINT ["/app/http-server"]
CMD ["-m=Hey"]
