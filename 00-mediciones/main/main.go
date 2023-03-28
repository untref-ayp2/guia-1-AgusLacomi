package main

import (
	"busquedas"
	"fmt"
	"ordenamiento"
	"sort"
	"time"
	"utiles"
)

func main() {

	arr := []int{1000, 5000, 10000, 15000, 20000, 25000, 50000, 100000, 150000, 200000, 250000, 500000, 1000000}

	for _, tamaño := range arr {

		arreglo := utiles.GenerarArreglo(10, tamaño)
		buscado := -1

		//fmt.Println(arreglo)

		inicio := time.Now()
		// Busqueda Lineal
		fmt.Println(busquedas.BusLineal(arreglo, buscado))
		fmt.Println("Busqueda Lineal: ", time.Since(inicio))

		inicio = time.Now()
		ordenamiento.Burbuja(arreglo)
		fmt.Println(time.Since(inicio))

		inicio = time.Now()
		// Ordenar el arreglo para la busqueda binaria
		sort.Ints(arreglo)
		fmt.Println("Ordenamiento: ", time.Since(inicio))
		//fmt.Println(arreglo)

		inicio = time.Now()
		// Busqueda Binaria
		fmt.Println(busquedas.BusquedaBinaria(arreglo, buscado))
		fmt.Println("Busqueda Binaria: ", time.Since(inicio))
	}
}
