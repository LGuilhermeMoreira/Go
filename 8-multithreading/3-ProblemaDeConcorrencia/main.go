package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var visitas uint64

// func main() {
// 	mux := http.NewServeMux()

// 	/*
// 		sse é um simples caso de concorrência:

// 		Como muitas requisições podem ser feitas ao mesmo tempo
// 		essa caracteristica sincrona não vai funcionar.

// 		caso eu faça 1000 requisições ao mesmo tempo,
// 		todas essa requisições não vão ser contadas.
// 	*/
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		visitas++
// 		w.Write([]byte(fmt.Sprint(visitas, "\n")))
// 	})

// 	http.ListenAndServe(":8080", mux)
// }

// func main() {
// 	mux := http.NewServeMux()

// 	mutex := sync.Mutex{}

// 	/*
// 		Podemos utilizar o mutex para travar e destravar variaveis
// 		de forma manual, a fim de evitar condições de corrida
// 	*/

// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		mutex.Lock()
// 		visitas++
// 		mutex.Unlock()
// 		w.Write([]byte(fmt.Sprint(visitas, "\n")))
// 	})

// 	http.ListenAndServe(":8080", mux)
// }

func main() {
	mux := http.NewServeMux()

	/*
		podemos utilizar operações atomica para impedir race conditions

		OBS: Operações atomicas são oprerações que
			impedem codições de corrida

		utilizando o pacote atomic, podemos fazer opreações atomicas
	*/

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&visitas, 1)
		w.Write([]byte(fmt.Sprint(visitas, "\n")))
	})

	http.ListenAndServe(":8080", mux)
}
