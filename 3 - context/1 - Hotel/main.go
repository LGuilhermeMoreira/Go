package main

import (
	"context"
	"fmt"
	"time"
)

// contexto controla oq a aplicação tá fazendo

// pode guardar informação dentro do contexto

// olhar documentação: https://pkg.go.dev/context

func main() {
	// iniciando um contexto
	ctx := context.Background()
	/*
		context.WithCancel: cancela depois de executar a função cancel  (ctx)
		context.WithTimeOut: cacela apos a passagem de um time out ou executando a função cancel (ctx,time.Second)
		context.WithDeadLine: cancela apos uma deadline ou executando a função cancel (ctx,time.Now())
		context.WithValue:  Permite associar dados arbitrários (chave-valor) ao contexto (ctx,chave,valor)
	*/
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Utiliza uma declaração select para lidar com múltiplos canais simultaneamente.
	select {
	// Caso o contexto seja cancelado (ex: timeout atingido), entra nesse caso.
	case <-ctx.Done():
		// Imprime uma mensagem indicando que a reserva do hotel foi cancelada devido ao tempo limite atingido.
		fmt.Println("hotel booking cancelled. timeout reached.")

	// Caso o canal retornado por time.After, que representa um temporizador de 5 segundos, envie um sinal.
	case <-time.After(time.Second): // em vez desse time after, poderia ser uma Api ou qualquer outra coisa
		// Imprime uma mensagem indicando que o hotel foi reservado com sucesso após 5 segundos.
		fmt.Println("hotel booked")
	}
}
