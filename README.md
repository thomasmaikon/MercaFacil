# MercaFacil
Api de cadastro de numeros de celular

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
