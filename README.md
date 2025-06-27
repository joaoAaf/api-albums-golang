# API de Gerenciamento de Albums

## Sobre a Aplicação

A API de Gerenciamento de Albums é uma aplicação web para estudo, que gerencia informações de álbuns de música. O projeto foi baseado no [tutorial](https://go.dev/doc/tutorial/web-service-gin) para criação de uma API REST com Go da própria documentação da linguagem, com a adição da conexão com um banco de dados MongoDB e recursos de monitoramento com Prometheus e Grafana.

## Recursos Utilizados

*   [Go](https://go.dev/) - Linguagem de programação utilizada para construir a API.
*   [Gin](https://gin-gonic.com/) - Framework web para Go.
*   [MongoDB](https://www.mongodb.com/) - Banco de dados NoSQL utilizado para armazenar e recuperar os dados.
*   [Docker](https://docs.docker.com/get-docker/) - Utilizado para construir e executar o servidor.
*   [Prometheus](https://prometheus.io/) - Monitoramento e coleta de métricas da aplicação.
*   [Grafana](https://grafana.com/) - Visualização de métricas e dashboards.
*   [Kubernetes](https://kubernetes.io/) - Orquestração de containers e deploy da aplicação (opcional).

## Objetivos

*   [x] Criar um servidor web com Go e Gin.
*   [x] Conectá-lo ao banco de dados MongoDB.
*   [x] Realizar CRUD (Create, Read, Update, Delete) de álbuns.
*   [x] Expor métricas customizadas para Prometheus.
*   [x] Visualizar métricas no Grafana.
*   [x] Deploy automatizado via GitHub Actions e Kubernetes.

### Arquivos

*   `main.go`: O arquivo principal da aplicação, responsável por iniciar o servidor.
*   `albumsRoutes.go`: O arquivo de rotas para a API, definindo as rotas para gerenciar álbuns.
*   `albunsController.go`: O arquivo de controladores para a API, responsável por fornecer uma interface para manipular os dados dos álbuns.
*   `mongodb.go`: O arquivo de repositório MongoDB, responsável por interagir com o banco de dados.
*   `album.go`: O arquivo de modelo de álbuns, definindo a estrutura dos dados.
*   `env.go`: O arquivo que define as variáveis de ambiente da aplicação.
*   `metrics.go` e `prometheus.go`: Arquivos de configuração e middleware para métricas Prometheus.
*   `docker-compose.yml`: Sobe API, MongoDB, Prometheus e Grafana em containers.
*   Diretório `k8s/`: Manifests para deploy no Kubernetes (namespaces, volumes, deployments, services, configmaps, etc).
*   Diretório `monitoring/`: Configurações do Prometheus e dashboards do Grafana.

## Como Usar

### Com Docker Compose

1.  Clone o repositório: `git clone https://github.com/joaoAaf/api-doc-go-mongo.git`
2.  Entre na pasta raiz do projeto.
3.  Execute o comando para construir e subir os containers: `docker-compose up -d`
4.  Os recursos da API estarão disponíveis em `http://localhost:8080/albums`
5.  O Grafana estará disponível em `http://localhost:3000` (login padrão: admin/admin).
6.  Use o comando `curl` ou um cliente HTTP como [Postman](https://www.postman.com/) para realizar requisições HTTP para a API.

### Com Kubernetes

1.  Certifique-se de ter um cluster Kubernetes configurado.
2.  Aplique os manifests do diretório `k8s/`:
    ```sh
    kubectl apply -f ./k8s/init/ -R
    kubectl apply -f ./k8s/albums-app/pvc/ -R
    kubectl apply -f ./k8s/albums-app/deploy_svc/ -R
    ./k8s/albums-monitoring/pvc_configmap/grafana-dashboards.sh ./monitoring/grafana/provisioning/dashboards
    kubectl apply -f ./k8s/albums-monitoring/pvc_configmap/ -R
    kubectl apply -f ./k8s/albums-monitoring/deploy_svc/ -R
    ```
3.  Os serviços estarão disponíveis conforme configurado nos arquivos de Service (NodePort/ClusterIP).

## Exemplos de requisições:

*   `curl -X POST 'http://localhost:8080/albums' --data '{"title": "Um álbum qualquer","artist": "Fulano de Tal","price": 2.99}'`
*   `curl -X POST 'http://localhost:8080/albums' --data '{"title": "Esse é bom","artist": "Beltrano da Silva","price": 1.89}'`
*   `curl -X GET 'http://localhost:8080/albums'`
*   `curl -X GET 'http://localhost:8080/albums/{id}'`
*   `curl -X PUT 'http://localhost:8080/albums/{id}' --data '{"title": "Um álbum qualquer","artist": "Fulano de Castro","price": 10.55}'`
*   `curl -X DELETE 'http://localhost:8080/albums/{id}'`

## Monitoramento

- As métricas Prometheus estão disponíveis em `/metrics` na API.
- Dashboards customizados para Grafana estão em `monitoring/grafana/provisioning/dashboards/`.
- O Prometheus coleta métricas da API automaticamente via configuração e anotações nos serviços do Kubernetes.

## CI/CD

- O projeto possui workflows no GitHub Actions para testes, build e deploy automático no Kubernetes.
- Foi utilizado uma maquina self-hosted para realizar o deploy, consulte este [link](https://docs.github.com/en/actions/how-tos/hosting-your-own-runners/managing-self-hosted-runners/adding-self-hosted-runners) para mais detalhes.
- Veja os arquivos em [`.github/workflows/`](.github/workflows/).

## Dicas e Precauções

*   Certifique-se de que o **Docker** e **Docker Compose** estão instalados em sua máquina antes de executar o comando para construir e subir os containers.
*   A porta padrão utilizada pela API é a **8080**, mas você pode alterá-la no arquivo `docker-compose.yml` ou nos manifests do Kubernetes.
*   Você pode personalizar as variáveis de ambiente da aplicação criando um arquivo `.env` e substituindo os valores das variáveis no `docker-compose.yml` por `${VARIAVEL}`.
*   A única variável de ambiente necessária para executar a aplicação é a `MONGO_CONNECTION`, que deve conter a string de conexão com o banco de dados. Seu valor padrão é `mongodb://localhost:27017`.

### Licença

A API de Gerenciamento de Albums é distribuída sob a licença MIT. Você pode usar, modificar e distribuir o código livremente, desde que você inclua a licença na sua