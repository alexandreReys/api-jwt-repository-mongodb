# API Golang

### Extensão Go VsCode

### Instal :

curl -OL https://golang.org/dl/go1.18.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.18.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

### Pega a url do repositorio do github com o https

❯ go mod init github.com/alexandreReys/api-jwt-repository-mongodb
go: creating new go.mod: module github.com/alexandreReys/api-jwt-repository-mongodb

### Executar

go run main.go ou
go run .

### dotenv

*https://github.com/joho/godotenv*

❯ go get github.com/joho/godotenv
go: added github.com/joho/godotenv v1.5.1

### Gin Gonic Web

*https://github.com/gin-gonic/gin*

> go get -u github.com/gin-gonic/gin

### Go Package validator

*https://github.com/go-playground/validator/blob/master/README.md*

> go get github.com/go-playground/validator/v10

### Go Package zap

*https://github.com/uber-go/zap*

> go get -u go.uber.org/zap
