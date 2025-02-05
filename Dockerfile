# Usa a versão mais recente do Go compatível com seu projeto
FROM golang:1.23

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para dentro do container
COPY . .

# Baixa as dependências
RUN go mod tidy

# Compila o binário do programa
RUN go build -o main .

# Expõe a porta que o serviço vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/app/main"]
