FROM golang:1.18.3 AS build
WORKDIR /go/src/boilerplate

COPY . .
COPY ./main.go .

RUN go mod tidy
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest AS release
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /go/src/boilerplate .
RUN rm -rf ./main.go

RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/app

CMD ["./app", "--prod"]
ENTRYPOINT ["/usr/bin/dumb-init", "---"]