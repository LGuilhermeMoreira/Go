package main

import (
	"bufio"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
)

func return3numbers() (int, int, int) {
	return 1, 2, 3
}

func aceppt3numbers(a, b, c int) int {
	return a + b + c
}

func intanyway(a *int) {
	println(a)
}

func intanyway2(a int) {
	println(a)
}

//func foo(a, b, c, d int)

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	// a := aceppt3numbers(return3numbers())
	// println(a)

	// i := 12

	// i2 := &i

	// intanyway(&i)

	// intanyway2(*i2)

	// foo(
	// 	1,
	// 	2,
	// 	3,
	// 	4,
	// )

	// time.Sleep(3 * time.Second)

	// clearTerminal()

	// mapa := map[string]string{"cor": "azul", "cor2": "verder"}

	// println(mapa["cor"])
	// println(mapa["cor2"])

	println("resposta: ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	value := scanner.Text()

	scanner.Scan()

	value2 := scanner.Text()
	print(value, value2)

	clearTerminal()

	i, err := strconv.ParseInt(value, 0, 64)

	if err != nil {
		panic(err)
	}

	print(i, reflect.TypeOf(i))

}
