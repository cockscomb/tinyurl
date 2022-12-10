# syntax=docker/dockerfile:1.4
FROM golang:1.19-bullseye AS builder

WORKDIR /go/src/github.com/cockscomb/tinyurl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/tinyurl


FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder --chown=nonroot:nonroot /go/bin/tinyurl /tinyurl
CMD ["/tinyurl"]
