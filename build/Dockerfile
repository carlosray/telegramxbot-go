FROM golang:1.19-alpine as build_base
WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./xbot-app ./cmd/xbot/main.go

FROM alpine:3.17.3
COPY --from=build_base /tmp/app/xbot-app /app/xbot-app
COPY configs ./configs
CMD ["/app/xbot-app"]
