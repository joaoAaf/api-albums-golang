# API de Gerenciamento de Albums

## Sobre a Aplicação

A API de Gerenciamento de Albums é uma aplicação web para estudo, que gerencia informações de álbuns de música. O projeto foi baseado no [tutorial](https://go.dev/doc/tutorial/web-service-gin) para criação de uma API REST com Go da propria documentação da linguagem, com a adição da conexão com um banco de dados MongoDB.

## Recursos Utilizados

*   [Go](https://go.dev/) - Linguagem de programação utilizada para construir a API.
*   [Gin](https://gin-gonic.com/) - Framework web para Go.
*   [MongoDB](https://www.mongodb.com/) - Banco de dados NoSQL utilizado para armazenar e recuperar os dados.
*   [Docker](https://docs.docker.com/get-docker/) - Utilizado para construir e executar o servidor.

## Objetivos

*   [x] Criar um servidor web com Go e Gin.
*   [x] Conecta-lo ao banco de dados MongoDB.
*   [x] Realizar CRUD (Create, Read, Update, Delete) de álbuns.

### Arquivos

*   `main.go`: O arquivo principal da aplicação, responsável por iniciar o servidor.
*   `albunsRoutes.go`: O arquivo de rotas para a API, definindo as rotas para gerenciar álbuns.
*   `albunsController.go`: O arquivo de controladores para a API, responsável por fornecer uma interface para manipular os dados dos álbuns.
*   `mongodb.go`: O arquivo de repositório MongoDB, responsável por interagir com o banco de dados.
*   `album.go`: O arquivo de modelo de álbuns, definindo a estrutura dos dados.
*   `env.go`: O arquivo que define as variáveis de ambiente da aplicação.

## Como Usar

Para usar a API de Gerenciamento de Albums, você precisará seguir os passos abaixo:

1.  Clone o repositório: `git clone https://github.com/joaoAaf/api-doc-go-mongo.git`;
1.  Entre na pasta raiz do projeto;
1.  Execute o comando para construir e subir os containers: `docker-compose up -d`;
1.  Os recursos da API estarão disponíveis em `/albums`;
1.  Use o comando `curl` ou um cliente HTTP como [Postman](https://www.postman.com/) para realizar requisições HTTP para a API.

## Exemplos de requisições:

*   `curl -X POST 'http://localhost:8080/albums' --data '{"title": "Um álbum qualquer","artist": "Fulano de Tal","price": 2.99}'`
*   `curl -X POST 'http://localhost:8080/albums' --data '{"title": "Esse é bom","artist": "Beltrano da Silva","price": 1.89}'`
*   `curl -X GET 'http://localhost:8080/albums'`
*   `curl -X GET 'http://localhost:8080/albums/{id}'`
*   `curl -X PUT 'http://localhost:8080/albums/{id}' --data '{"title": "Um álbum qualquer","artist": "Fulano de Castro","price": 10.55}'`
*   `curl -X DELETE 'http://localhost:8080/albums/{id}'`

## Dicas e Precauções

*   Certifique-se de que o **Docker** e **Docker Compose** estão instalados em sua máquina antes de executar o comando para construir e subir os containers.
*   A porta padrão utilizada pela API é a **8080**, mas você pode alterá-la no arquivo `docker-compose.yml`.
*   Você pode personalizar as variáveis de ambiente da aplicação criando um arquivo `.env` e substituindo os valores das variáveis no `docker-compose.yml` por `${VARIAVEL}`.
*   A única variavel de ambiente necessária para executar a aplicação é a `MONGO_CONNECTION`, que deve conter a string de conexão com o banco de dados, seu valor padrão é `mongodb://localhost:27017`.


### Licença

A API de Gerenciamento de Albums é distribuída sob a licença MIT. Você pode usar, modificar e distribuir o código livremente, desde que você inclua a licença na sua distribuição.