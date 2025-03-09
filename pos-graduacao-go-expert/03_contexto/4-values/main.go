package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

// por convenção, o contexto é o primeiro parâmetro
func bookHotel(ctx context.Context) {
	// usar com sabedoria, pois pode ser que não haja esse valor no contexto
	token := ctx.Value("token")
	fmt.Println(token)
}
