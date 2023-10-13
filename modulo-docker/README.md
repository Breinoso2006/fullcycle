# Docker

## Comandos utilizados ao longo das aulas e suas respectivas funções

- `docker ps` => mostrar os containers ativos.
  - -a, mostra os containers ativos e inativos.
- `docker run {imagem}:{versão} {comando}` => executa a imagem em um container juntamente com um comando específico, caso não ponha a versão a chamada será a `latest`.
  - --name {nome para o container} {imagem}, para subir o container com um nome específico
  - --rm, irá apagar o container assim que ele for parado.
  - -it, é o mesmo que -i -t: -i é o modo interativo e vai manter o stdin ativo; -t representa tty, para podermos digitar no terminal.
  - -d, para desatachar o terminal do processo que iremos rodar.
  - -p {porta da minha máquina}:{porta do container}, redireciona a minha porta para a porta do container, publica a porta.
  - -v {caminho na sua máquina}:{caminho dentro do container}, para montar um volume do seu computador dentro do container e não perder os arquivos alterados. Pode criar a pasta se não existir. Podemos apontar um volume neste caso também.
  - --mount type=bind,source={caminho na sua máquina},target={caminho dentro do container}, mais esplícito que o -v e não cria a pasta se ela não existir.
  - --mount type=volume,source={nome do volume},target={caminho no container}, utiliza um volume criado no pc que pode ser compartilhado entre outros containers.
- `docker start {container}` => sobe um container anteriormente criado.
- `docker rm {container}` => para remover um container.
- `docker exec {container} {comando}` => para executar o comando no container que está de pé.
- `docker volume {nome para o volume}` => cria um volume dentro do meu pc, que poderá ser utilizado pelos containers para armazenar dados.
- `docker volume prune` => limpa volumes não utilizados.
- `docker pull {imagem}` => quando quisermos apenas baixar a imagem.
- `docker images` => ver as imagens que estão baixadas no computador.
- `docker rmi {imagem}` => apagar imagem específica. 

## Estrutura no docker ps

- CONTAINER_ID: literalmente um id que representa o container, um hash
- IMAGE: o nome da imagem utilizada pelo container
- COMMAND: o comando que foi executado pelo container
- CREATED: quando o container foi criado
- STATUS: status atual do container
- PORTS: as portas que esse container utilizou
- NAMES: nome dado ao container (caso não dê na hora do run, será criado um aleatório)
