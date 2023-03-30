FROM golang:1.20-alpine as BuildStage

WORKDIR /app

COPY . .
RUN go mod download
EXPOSE 8080
RUN go build -o /test main.go

FROM alpine:latest
WORKDIR /
COPY --from=BuildStage /test /test
EXPOSE 8080
ENTRYPOINT [ "/test" ]
