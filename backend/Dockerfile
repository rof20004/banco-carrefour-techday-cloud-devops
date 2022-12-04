FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o backend *.go

FROM nginx:stable-alpine

COPY --from=builder /app/backend backend
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY init.sh init.sh

RUN chmod +x init.sh

EXPOSE 80

CMD ./init.sh