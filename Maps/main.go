package main

import "fmt"

func main() {

	var n int
	fmt.Println("Ingrese el numero de estudiantes a evaluar: ")
	fmt.Scanln(&n)

	estudiantes := make(map[string]float64)
	var nombre string
	var nota float64

	for i := 0; i < n; i++ {
		fmt.Println("Estudiantes: ", i+1)
		fmt.Scanln(&nombre)
		fmt.Println("Nota: ")
		fmt.Scanln(&nota)
		estudiantes[nombre] = nota
	}

	fmt.Println("\nDatos ingresados: ")
	for name, grade := range estudiantes {
		fmt.Println(name, ":", grade)
	}

	max := 0.0
	min := 11.0
	for _, nota := range estudiantes {
		if nota > max {
			max = nota
		}
		if nota < min {
			min = nota
		}
	}
	fmt.Println("\nLa nota maxima es:", max)
	fmt.Println("\nLa nota minima es:", min)
}
