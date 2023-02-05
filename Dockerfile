FROM golang:1.19-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
# ENV GOARCH=arm
# ENV GOARM=6

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd cmd

RUN go build -a -tags netgo -ldflags '-w' -o nhlstats 

# FROM gcr.io/distroless/static
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chown=0:0 --from=builder /app/nhlstats /bin/

EXPOSE 3000

ENTRYPOINT ["/bin/nhlstats"]