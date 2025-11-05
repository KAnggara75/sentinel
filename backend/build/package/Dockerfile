FROM golang:1.25.3-trixie AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -x -ldflags="-w -s" -trimpath -o sentinel ./cmd/sentinel

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /build/sentinel /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]
