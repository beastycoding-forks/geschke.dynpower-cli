FROM golang:alpine as builder

LABEL maintainer="Ralf Geschke <ralf@kuerbis.org>"
LABEL last_changed="2020-09-24"

RUN apk update && apk add --no-cache git
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o dynpower-cli .

# Build minimal image with dynpower-cli binary
FROM scratch
COPY --from=builder /build/dynpower-cli /app/
WORKDIR /app
CMD ["./dynpower-cli"]