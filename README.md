# Clean Architecture
Desafio Clean Architecture - Go Express

## Índice
1. [Pré-requisitos](#pré-requisitos)
2. [Configuração do ambiente](#configuração-do-ambiente)
3. [Interagindo com a API](#interagindo-com-a-api)
4. [Criando pedidos](#criando-pedidos)
5. [Listando pedidos](#listando-pedidos)

## Pré-requisitos
Assegure-se de ter as seguintes ferramentas instaladas:
- [Golang](https://go.dev/doc/install)
- [Docker + Docker Compose](https://docs.docker.com/compose/install/)
- [evans](https://github.com/ktr0731/evans) 

## Configuração do ambiente
O projeto possui um Makefile com comandos utilitários para execução do ambiente. Alguns dos mais úteis estão listados abaixo:

- Sobe todos os containers necessários para rodar o projeto:
```
$ make docker-up
```

- Derruba todos os containers necessários para rodar o projeto:
```
$ make docker-down
```


- Executa o comando de migrate para aplicar os scripts de migração de banco de dados. Deve ser executado após o docker-up:
```
$ make migrate
```


- Executa o comando de migrate down caso queira destruir o conteudo do banco de dados e recriar a migration:
```
$ make migrate-down
```


## Interagindo com a API
A aplicação possui três tipos de API:
- HTTP REST: responde na porta 8000
- gRPC: responde na porta 50051
- GraphQL: responde na porta 8080

## Criando pedidos 
Após subir o projeto, pode-se criar pedidos de três maneiras:

- **Via HTTP:** Execute o arquivo "./api/create_order.http".
- **Via gRPC:** Execute os comandos:
```
$ make grpc 
$ pb.OrderService@127.0.0.1:50051> call CreateOrder
```


## Listando pedidos
Após subir o projeto, pode-se listar pedidos de três maneiras:

- **Via HTTP:** Execute o arquivo "./api/list_orders.http".
- **Via gRPC:** Execute os comandos:
```
$ make grpc 
$ pb.OrderService@127.0.0.1:50051> call ListOrders
```