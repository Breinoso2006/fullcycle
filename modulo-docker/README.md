# Docker

## Comandos utilizados ao longo das aulas e suas respectivas funções

### Padrão
- `docker ps` => mostrar os containers ativos.
  - -a, mostra os containers ativos e inativos.
  - -q, mostra apenas os id's dos containers.
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
  - `docker rm $(docker ps -a -q) -f` => derruba e apaga todos os containers de uma vez.
- `docker exec {container} {comando}` => para executar o comando no container que está de pé.
- `docker volume {nome para o volume}` => cria um volume dentro do meu pc, que poderá ser utilizado pelos containers para armazenar dados.
- `docker volume prune` => limpa volumes não utilizados.
- `docker pull {imagem}` => quando quisermos apenas baixar a imagem.
- `docker images` => ver as imagens que estão baixadas no computador.
- `docker rmi {imagem}` => apagar imagem específica. 
- `docker build -t {nome da tag ou usuário}/{nome da imagem}:{versão} {diretório onde está o arquivo}` => criar sua própria imagem a partir de um Dockerfile.
- `docker push {imagem}` => para subir uma imagem proprietária no dockerhub.
### Network
- `docker network` => mostra os comandos relacionados a redes.
- `docker network inspect {tipo}` => inspecionar rede.
- `docker network create --driver {tipo de rede} {nome da rede}` => para criar uma rede nova.
- `docker network connect {rede} {container}` => para conectar um container a determinada rede.
## Estrutura no docker ps

- CONTAINER_ID: literalmente um id que representa o container, um hash
- IMAGE: o nome da imagem utilizada pelo container
- COMMAND: o comando que foi executado pelo container
- CREATED: quando o container foi criado
- STATUS: status atual do container
- PORTS: as portas que esse container utilizou
- NAMES: nome dado ao container (caso não dê na hora do run, será criado um aleatório)

## Arquivo Dockerfile

- `FROM {imagem}:{versão}` => imagem base a ser utilizada.
- `RUN {comando}` => comando que será executado ao longo do build.
- `WORKDIR` => para criar uma pasta dentro do container e nos deixar ali dentroquando entrarmos.
- `COPY {arquivo ou pasta a ser copiada} {para onde ela irá no container}` => copiar arquivo que está dentro do computador para o container.
- `ENTRYPOINT [comando, parâmetro]` => executará o comando em questão, mas não pode ser substituído pelo usuário.
- `CMD [comando, parâmetro]` => executará o comando em questão, podendo ser substituído. Entra como parâmetro do entrypoint e é considerado como "padrão" caso o usuário não passe outro parâmetro.
- `ENV {variável}` => para criar variáveis de ambiente.
- `EXPOSE` => para expôr uma porta.

## Observações
- `exec "$@"` => no final de um arquivo .sh significa que irá executar o comando que vier depois dele.
- `http://host.docker.internal:{porta}` => para acessar algum recurso da sua máquina estando no container docker.
- `docker build -t {tag}/{imagem} {diretório} -f {caminho até o dockerfile desejado}` =>  caso estejamos trabalhando com 2 dockerfile's na mesma pasta (local/prod).