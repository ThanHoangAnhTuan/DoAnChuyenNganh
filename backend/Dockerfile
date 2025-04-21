FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ecommerce.golang.com ./cmd/server

FROM scratch

COPY ./configs ./configs

COPY --from=builder /build/ecommerce.golang.com /

ENTRYPOINT [ "/ecommerce.golang.com", "configs/local.yaml" ]