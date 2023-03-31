package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {

	_, err := http.Get(servidor)

	if err != nil {
		canal <- servidor + "no está disponible"
		//fmt.Println(servidor, "no está disponible")
	} else {
		//fmt.Println(servidor, "está funcionando")
		canal <- servidor + "está funcionando"
	}

}

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
	}

	for _, servers := range servidores {
		go revisarServidor(servers, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("tiempo de ejecución ", tiempoPaso)
}
