FROM golang:alpine AS builder 
WORKDIR /app 
ADD . /app 
# build cmd 
RUN go build -ldflags="-w -s" -o release/test_temp main.go 

# Compress the binary
RUN apk add upx && upx -o  release/test release/test_temp 
RUN rm release/test_temp

## Put the binary to new image
FROM alpine 
RUN set -eux; \
    apk update \
    && apk upgrade \
    && apk add --no-cache \
    dumb-init 
WORKDIR /app 
COPY --from=builder /app/release /app/ 

## Run the service

ENTRYPOINT [ "dumb-init", "--", "/app/test" ]
EXPOSE 9512