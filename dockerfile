FROM golang:1.23.3-alpine as builder
WORKDIR /cungsao
COPY . /cungsao

ENV BUILD_TAG 1.0.0
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go mod tidy
RUN go build -o cungsao /cungsao/cmd/server/main.go

# stage2.1: rebuild
FROM alpine
WORKDIR /cungsao
RUN apk add --no-cache \
    ffmpeg \
    build-base \
    git \
    bash
COPY --from=builder /cungsao/cungsao /cungsao/cungsao
COPY --from=builder /cungsao/zalo-audio/ /cungsao/zalo-audio/
COPY --from=builder /cungsao/exported-file/ /cungsao/exported-file/
COPY --from=builder /cungsao/config/production.yaml /cungsao/config/production.yaml
COPY --from=builder /cungsao/config/dev.yaml /cungsao/config/dev.yaml

CMD ["./cungsao"]