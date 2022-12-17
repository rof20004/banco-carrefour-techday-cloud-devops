FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o backend *.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/backend backend

EXPOSE 8080

CMD ["./backend"]

# Build antigo com o uso do Nginx para expôr a api em Go

#FROM nginx:stable-alpine

#COPY --from=builder /app/backend backend
#COPY nginx.conf /etc/nginx/conf.d/default.conf
#COPY init.sh init.sh

#RUN chmod +x init.sh

#EXPOSE 80

#CMD ./init.sh