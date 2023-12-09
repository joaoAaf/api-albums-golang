#imagem base estagio de build
FROM golang:alpine AS build
#local onde serão copiados os arquivos da aplicação
WORKDIR /app
#copia os arquivos de dependência para a pasta api
COPY api/go.mod api/go.sum ./
#baixa e verifica as dependências do arquivo de configuração
RUN go mod download && go mod verify
#copia os arquivos da aplicação para a pasta api
COPY api .
#constroi a imagem da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o ./api
#imagem base para stagio de realease
FROM gcr.io/distroless/static-debian12
#copia o binario da aplicação do estagio de build
COPY --from=build /app/api /usr/local/bin/api
#informa ao docker que o container ouve a porta especificada
EXPOSE 8080
#define o usuário e o grupo para container para nonroot
USER nonroot:nonroot
#executa a aplicação
CMD [ "api" ]