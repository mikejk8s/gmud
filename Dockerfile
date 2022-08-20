FROM golang:alpine AS build

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go get -d -v \
  # This strips out debug information for a smaller binary
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/vulnerability-report

# ======================================================================================================================
FROM scratch AS run

COPY --from=build /app/gmudt /app/gmud
ENTRYPOINT ["/app/gmud"]