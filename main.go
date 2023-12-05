package main

import (
	"fmt"
	"go-golang/variables"
)

func  main()  {
		estado, texto := variables.ConviertoaTexto(1588)
		fmt.Println(estado)
		fmt.Println(texto)
}
