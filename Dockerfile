# Usar uma imagem base do Go
FROM golang:1.18.2-alpine AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar apenas os arquivos de go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o restante do código-fonte para o container
COPY . .

# Copiar o arquivo de configuração
COPY config.yaml .  

# Compilar o aplicativo
RUN go build -o go-auth main.go

# Usar uma imagem menor para rodar o aplicativo
FROM alpine:latest  

WORKDIR /root/

# Copiar o binário compilado da imagem de builder
COPY --from=builder /app/go-auth .
COPY --from=builder /app/config.yaml .  

# Expor a porta que o aplicativo vai usar
EXPOSE 8000

# Comando para rodar o aplicativo
CMD ["./go-auth"]
