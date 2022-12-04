FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o backend *.go

FROM nginx:stable-alpine

COPY --from=builder /app/backend backend
COPY nginx.conf /etc/nginx/conf.d/default.conf

RUN chmod +x backend
RUN nohup ./backend > backend.log 2>&1 &

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]