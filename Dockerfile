#imagem base
FROM golang:alpine
#local onde serão copiados os arquivos da aplicação
WORKDIR /api
#copia os arquivos de dependência
COPY api/go.mod api/go.sum ./
#baixa as dependências do arquivo de configuração
RUN go mod download && go mod verify
#copia os arquivos da aplicação para a raiz
COPY api .
#builda a imagem
RUN go build -o /usr/local/bin/api
#define o comando para executar a aplicação
CMD [ "api" ]