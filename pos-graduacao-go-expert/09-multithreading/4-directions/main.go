package main 

func main(){
	channel := make(chan string)
	x := "Hello, world!"
	go recebe(x, channel)
	le(channel)
}

// A seta indica que o canal apenas recebe valores (receive only)
func recebe(nome string, ch chan<- string) {
	// Preenche o canal com uma string
	ch <- nome
}

// A seta indica que o canal apenas "esvazia" valores (send only)
func le(ch <-chan string) {
	// Le a string do canal
	println(<-ch)
}
