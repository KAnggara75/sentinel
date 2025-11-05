# ---------- STAGE 1: Builder ----------
FROM golang:1.25.3-trixie AS builder

# set module cache & build dir
WORKDIR /app

# copy dan download dependency dulu biar cache efisien
COPY backend/go.mod backend/go.sum ./
RUN go mod download && go mod verify

# copy seluruh source code backend
COPY backend/ .

# build binary tanpa simbol debug & dengan path yang rapih
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /sentinel ./cmd/sentinel

# ---------- STAGE 2: Runtime ----------
FROM gcr.io/distroless/static-debian12:nonroot

# copy binary dari builder
COPY --from=builder /sentinel /app/sentinel

# set working dir & port
WORKDIR /app
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/app/sentinel"]
