# Usar a imagem oficial do Golang
FROM golang:1.22-alpine

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código da aplicação
COPY . .

# Definir o diretório de trabalho para o local do main.go
WORKDIR /app/cmd/server

# Compilar a aplicação
RUN go build -o /app/myapp

# Definir o diretório de trabalho para a raiz da aplicação
WORKDIR /app

# Expor a porta em que a aplicação irá rodar
EXPOSE 8080

# Comando de inicialização da aplicação
CMD ["./myapp"]
