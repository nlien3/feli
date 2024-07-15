# Etapa de construcción
FROM golang:1.18 AS builder
LABEL authors="Arnon"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Etapa de ejecución
FROM ubuntu:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Instalar ca-certificates para las conexiones seguras
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean
EXPOSE 8080
# Esperar a que MySQL esté disponible antes de iniciar la aplicación
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

CMD ["wait-for-it.sh", "db:3306", "--", "./main"]