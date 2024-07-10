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
