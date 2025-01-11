# Utilizando uma imagem leve do Go
FROM golang:1.20-alpine

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando o arquivo go.mod e go.sum para resolver dependências primeiro
COPY go.mod go.sum ./

# Baixar as dependências antes de copiar o código
RUN go mod download

# Copiar o restante do projeto para dentro do container
COPY . .

# Compilar o projeto com o ponto de entrada em 'cmd/'
RUN go build -o carros-pcd-service ./cmd

# Definir a porta para o Railway (PORT é setado automaticamente)
EXPOSE 8080

# Comando de inicialização do container
CMD ["./carros-pcd-service"]
