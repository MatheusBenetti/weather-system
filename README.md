# Clima Atual

Digite o CEP desejado para receber a temperatura do momento em graus Celsius, Fahrenheit e Kelvin

## Conteúdo

- [Instalação](#instalação)
- [Ambiente de Desenvolvimento](#developer)
- [Ambiente de Produção](#production)
- [Testes](#testes)
- [Versão de Deploy Google Cloud](#versão-de-deploy)

## Instalação

Para fazer o download das dependências, utilize:
```bash
go mod tidy
```

## Ambiente de Desenvolvimento

Para rodar o projeto em ambiente de desenvolvimento, utilize para criar o container:
```
docker compose up --build
```
ou então utilize docker-compose up --build, dependendo da versão, após gerar a imagem, pode iniciar as próximas vezes com docker compose up.

Depois para acessar o bash, abra outro terminal e digite:
```
docker compose exec web bash
```
E então rode o comando:
```
go run main.go
```
Após isso é só fazer uma requisição no terminal com:
```
curl http://localhost:8080/getTemperature?cep=95670084
```
Ou via Postman/Insomnia

## Ambiente de Produção

Execute o seguinte comando para gerar um build para Linux:
```
docker compose -f docker-compose.yml -f docker-compose.prod.yml up
```
E execute pelo terminal o comando:
```
./temp-system
```
Após isso é só fazer uma requisição no terminal com:
```
curl http://localhost:8080/getTemperature?cep=95670084
```
Ou via Postman/Insomnia

## Testes

Para rodar os testes, utilize:
```
go test .
```

## Versão de Deploy Google Cloud

Para utilizar a versão de deploy, degite a url abaixo no seu navegado, também alterando o cep para o desejado no formato 00000000:
```
https://weather-system-rswqsctlxa-uc.a.run.app/getTemperature?cep=95670084
```
