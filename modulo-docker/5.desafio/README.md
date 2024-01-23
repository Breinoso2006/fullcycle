# desafio-docker-fullcycle
Desafio para conclusão do módulo sobre Docker do curso FullCycle 3.0


## Desafio 1
Criar uma imagem em Go que quando executarmos `docker run <seu-user>/fullcycle` tenhamos o seguinte resultado: `Full Cycle Rocks!!`

A imagem deverá ser publicada no Docker Hub e ter menos de 2mb.

Link: https://hub.docker.com/r/breinoso2006/fullcycle

## Desafio 2
A idéia principal é que quando um usuário acesse o nginx, o mesmo fará uma chamada em nossa aplicação node.js. Essa aplicação por sua vez adicionará um registro em nosso banco de dados mysql, cadastrando um nome na tabela people.

O retorno da aplicação node.js para o nginx deverá ser: `<h1>Full Cycle Rocks!</h1>` e a lista de nomes cadastrada no banco de dados.

Gere o docker-compose de uma forma que basta apenas rodarmos: `docker-compose up -d` que tudo deverá estar funcionando e disponível na porta: 8080.

A linguagem de programação para este desafio é Node/JavaScript.