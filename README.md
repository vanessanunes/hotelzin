# hotalzin

## utilizando
Golang
Chi
Docker
~~Go Migrate~~
Swaggo

## configurando

Inicialmente temos dois arquivos para configurações, sendo eles o `.env` e o `config.toml`
Escrevi o `.env` para o docker e o `config.toml` para a api em golang.

Execute o comando `mv env-example .env` para renomear o arquivo

Edite as informações se achar necessário :)

## rodar via docker

Edite as variaveis de ambientes e rode o comando `make docker-local`

## rodar api localmente

Pode ser diretamente do terminal `go run main.go` ou se preferir `make app`

## swagger

Para acessar, basta apenas rodar o app e acesse pelo: `swagger/index.html`

