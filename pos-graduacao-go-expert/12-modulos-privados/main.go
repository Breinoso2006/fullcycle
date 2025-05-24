package main

// repositórios privados normalmente não são possíveis de serem importados
// utilizar export GOPRIVATE=github.com/breinoso2006/...
// necessário ter autenticação no github (token ou ssh)
// bitbucket também tem repositórios privados
import (
	"fmt"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/11-manipulando-eventos/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}


// O go possui um proxy para garantir a entrega de pacotes, pra caso o github esteja 
// fora do ar e também para devolver de forma mais rápida esses dados, centralizando 
// pacotes mais baixados por exemplo. ([go proxy](https://proxy.golang.org/))
// A forma ideal de ter todas as suas dependências na raiz do seu projeto e 
// mitigar possíveis problemas com pacotes não cacheados é usando go mod vendor