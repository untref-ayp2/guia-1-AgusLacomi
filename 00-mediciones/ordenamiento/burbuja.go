package ordenamiento

// By AgusLacomi
func Burbuja(arreglo []int) {

	n := len(arreglo)
	// iteramos n-1 veces (en cada iteración dejamos fijo el elemento más grande)
	for i := 0; i < n-1; i++ {
		// en cada iteración comparamos elementos adyacentes y los intercambiamos si están en el orden incorrecto
		for j := 0; j < n-i-1; j++ {
			if arreglo[j] > arreglo[j+1] {
				arreglo[j], arreglo[j+1] = arreglo[j+1], arreglo[j]
			}
		}
	}
}
