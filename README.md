# Clima Atual

Digite o CEP desejado para receber a temperatura do momento em graus Celsius, Fahrenheit e Kelvin

## Conteúdo

- [Instalação](#instalação)
- [Utilização](#utilização)
- [Testes](#testes)
- [Docker](#docker)
- [API Requests](#api-requests)
- [Versão de Deploy](#versão-de-deploy)

## Instalação

Para fazer o download das dependências, utilize:
```bash
go mod tidy
```

## Utilização

Para rodar o projeto localmente, utilize:
```
go run main.go
```

## Testes

Para rodar os testes, utilize:
```
go test .
```

## Docker
Para rodar com o docker, primeiro faça o build com:
```
docker compose up --build
```
ou então utilize docker-compose up --build, dependendo da versão, após gerar a imagem, pode iniciar as próximas vezes com docker compose up.

## API Requests
Faça uma requisição http GET no Postman/Insomnia na url, podendo alterar o CEP para o desejado com o formato 00000000:
```
http://localhost:8080/getTemperature?cep=95670084
```

## Versão de Deploy

Para utilizar a versão de deploy, degite a url abaixo no seu navegado, também alterando o cep para o desejado no formato 00000000:
```
https://weather-system-rswqsctlxa-uc.a.run.app/getTemperature?cep=95670084
```
