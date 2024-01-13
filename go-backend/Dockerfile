ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY internal/data ./internal/data

RUN CGO_ENABLED=0 GOOS=linux go build -o /crease-coach

ARG PORT=4000
EXPOSE ${PORT}

ENTRYPOINT ["/crease-coach"]
