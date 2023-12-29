package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "token", "valor")

	bookHotel(ctx, "Pedra dos ventos")
}

func bookHotel(ctx context.Context, nameHotel string) {
	token := ctx.Value("token")

	fmt.Printf("%s tem o token: %v\n", nameHotel, token)
}
