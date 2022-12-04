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
COPY process.sh process.sh

RUN chmod +x process.sh

ENV DB_HOST=172.17.0.3
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASS=root
ENV DB_NAME=carrefour

EXPOSE 80

CMD ./process.sh