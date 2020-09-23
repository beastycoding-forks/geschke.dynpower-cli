FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o dynpower-cli .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./dynpower-cli"]