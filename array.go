package main

import "fmt"

func array() {

var listaCompras [5]string 

listaCompras[0] = "Arroz\n"
listaCompras[1] = "Feijão\n"
listaCompras[2] = "Macarrão\n"
listaCompras[3] = "Leite\n"
listaCompras[4] = "Café\n"

//fmt.Println("Lista de Compras:", listaCompras)
//fmt.Println("Só o segundo item:", listaCompras[1])

for i := 0; i < len(listaCompras); i++ {
	fmt.Println("Item", i, ":", listaCompras[i])
}

}