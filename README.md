# MercaFacil
Api de cadastro de numeros de celular

---

## Utilização

Para utilizar essa aplicação é necessario que siga, pelo menos, os seguintes passos:

:warning: OBS(Utilizar o ```shell sudo```, antes dos comandos que utilizam o dokcer, caso seu usuario nao tenha permisao)

1. Instalar o Docker na sua maquina
2. Clonar o repositório
```shell git clone https://github.com/thomasmaikon/MercaFacil.git ```
3. Construir a imagem referente a aplicação. 
```shell docker build --tag aplication:1.0.0 . ```
4. Criar uma rede local para que o container da aplicacao, criado anterior mente possa conectar-se com os bancos que criaremos no proximo passo.
```shell docker network create --drive bridge backend```
1. Rodams executamos nosso docker compose para rodar o banco PostgreSQL e MySQL
```shell docker compose up -d```
6. Por fim executamos a imagem, que foi criada no topico 3, com alguns parametros
```shell docker run -p 8000:8000 --network=backend --name=app```

---

### Sobre

* Este trabalho consiste em implementar uma API em que temos 2 tipos de clientes
* Macapa e Varejao
* A api serve para cadastrar numero de celulares juntamente ao usuario daquele numero
* Para o cliente Macapa ele pretende cadastrar os dados na seguinte sequencia {nome, numero} -> [NOME, +XX (XX) XXXXX-XXXX]
* Para o cliente Varejao é cadastrado no seguinte formato {nome, numero} -> [nome, XXXXXXXXXXXXX]
* Cada cliente possue sua propria conexao com o banco, no caso, os utilizados foram MySql e PostreSQL para Macapa e Varejao respectivamente
* Os clientes so podem cadastrar dados referentes a eles, ou seja, apenas clientes Macapa é que conseguem cadastrar, caso eles tentem cadastrar na API do Varejao o acesso é recusado
* Para realizarem o cadastro os clientes precisam estar autenticados.
* Testes os foram implementados para verificar a conexao com o banco

## Tecnologias

Este projeto foi criado utilizando:

[Go(1.18.3)](https://go.dev/doc/)
[Gorm(1.23.5)](https://github.com/gin-gonic/gin)
[Gin(1.8.1)](https://gorm.io/index.html)
[jwt-go](https://github.com/dgrijalva/jwt-go)
[Docker(20.10.10)](https://docs.docker.com/)
