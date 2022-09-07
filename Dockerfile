FROM golang:1.19.1-bullseye as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./coordinator cmd/coordinator/main.go

RUN go build -o ./consumer cmd/consumer/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=builder /build/coordinator /coordinator

COPY --from=builder /build/consumer /consumer