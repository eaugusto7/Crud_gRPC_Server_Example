# Crud_gRPC_Server_Example

====================================================================================

Comandos:

Iniciar o banco de dados: docker-compose up
Iniciar a aplicação: go run main.go
Gerar arquivos go a partir do .proto: protoc --go_out=. --go-grpc_out=. proto/*.proto

======================================

Caso o computador dê algum problema de PATH ao executar o comando protoc, executar: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
Reiniciar a IDE após dar o comando
Caso ainda tenha problema, verificar as variaveis de ambiente do sistema operacional

======================================

Utilizar o programa BloomRPC para realizar as requisições, as requisições por padrão estão em localhost:9000 (pode ser alterado a porta)
Para visualizar o banco de dados utilizar o dBeaver
Obter id do container docker: docker container ls | grep postgres
Para obter ip do container docker: docker inspect [id container] | grep "IPAddress"

POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=database   
Porta utilizada no DB: 5432

====================================================================================

![Captura de tela de 2023-01-12 15-38-40](https://user-images.githubusercontent.com/53271581/212151460-4dfbc6f4-9ce8-42cd-85c1-5645dc13277f.png)
