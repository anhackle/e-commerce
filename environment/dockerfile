# stage1: build
FROM golang:1.23.3-alpine as builder
WORKDIR /e-commerce
COPY . /e-commerce

ENV BUILD_TAG 1.0.0
ENV GO111MODULE on
# This disables CGO, which is Go’s support for calling C code.
ENV CGO_ENABLED=0
# Set the target operating system for Go’s compiler 
ENV GOOS=linux
RUN go mod tidy
# go build with -ldflags to reduce binary size (strip debug info)
RUN go build -ldflags="-s -w" -o e-commerce /e-commerce/cmd/server/main.go

# stage2.1: run
FROM scratch
WORKDIR /e-commerce
COPY --from=builder /e-commerce/e-commerce /e-commerce/e-commerce
COPY --from=builder /e-commerce/config/production.yaml /e-commerce/config/production.yaml

EXPOSE 8082

CMD ["./e-commerce"]