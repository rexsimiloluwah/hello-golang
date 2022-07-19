FROM golang:1.16-alpine AS builder 
RUN mkdir /app 
COPY . /app 
WORKDIR /app 
# install dependencies
# RUN go mod download 
RUN go mod vendor
# generate the executable 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/main cmd/api/main.go

# for the multi-stage build
FROM alpine  
WORKDIR /app 
COPY --from=builder /app /app/

EXPOSE 8040 

CMD ["/app/bin/main"]
