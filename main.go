package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Print("elegir un numero del 1 al 5\n")
	fmt.Scan(&num)

	switch num {

	case 1:
		ejer1()
		//break

	case 2:
		ejer2()
		//break

	case 3:
		ejer3()
		//break

	case 4:
		ejer4()
		//break
		//
	case 5:
		ejer5()
		//break

	default:
		fmt.Println("¿sos idiota? del 1 al 5 te dije")
	}

}

func ejer1() {

	fmt.Println("Ejercicio 1.")
	xd := bufio.NewReader(os.Stdin)
	fmt.Println("Escribe una frase.")
	//var frase stri
	//frase = "hola como estas"
	//fmt.Scanln(&frase
	//frasea := bufio.NewScanner((os.Stdin))
	//fmt.Printf("Tu frase fue: %s", frase)
	frase, _ := xd.ReadString('\n')
	fmt.Println("Tu frase fue:", frase)
	//str := bufio.NewScanner(os.Stdin)
	//fmt.Printf("Tus frase fue: %s", str.Text())
}

///////////////////////////////////////////////////////////////////////////////
func ejer2() {

	//fmt.Println("----------------------------------------------------")
	fmt.Println("Ejercicio 2")
	fmt.Println("Escribe un numero.")
	var w int
	fmt.Scanln(&w)
	fmt.Printf("Tu numero es: %d\n ", w)
}

//////////////////////////////////////////////////////////////////////////////////////////
func ejer3() {

	//fmt.Println("----------------------------------------------------")
	fmt.Println("ejercicio 3")
	fmt.Println("Escribe tu primer num")
	var x, y int
	fmt.Scan(&x)
	fmt.Println("Escribe tu segundo num.")
	fmt.Scan(&y)
	fmt.Printf("total %d\n", x+y)
}

func ejer4() {

	//fmt.Println("----------------------------------------------------")
	fmt.Println("ejercicio 4")
	fmt.Println("Escribe tu nombre")
	var (
		nombre, apellido string
		edad             int
	)

	fmt.Scan(&nombre)
	fmt.Println("Escribe tu apellido")
	fmt.Scan(&apellido)
	fmt.Println("Escribe tu edad")
	fmt.Scan(&edad)

	fmt.Printf("Te llamas %s", nombre)
	fmt.Printf(", tu apellido es %s", apellido)
	fmt.Printf(" y tenes %d\n", edad)
}

func ejer5() {

	//fmt.Println("----------------------------------------------------")
	var a, b, c, d, e int
	fmt.Println("Escribe tu primer num")
	fmt.Scan(&a)
	fmt.Println("Escribe tu segundo num.")
	fmt.Scan(&b)
	fmt.Println("Escribe tu 3er num.")
	fmt.Scan(&c)
	fmt.Println("Escribe tu 4to num.")
	fmt.Scan(&d)
	fmt.Println("Escribe tu 5to num.")
	fmt.Scan(&e)
	fmt.Println(a, "+", b, "+", c, "+", d, "+", e, "=", a+b+c+d+e)

	fmt.Println("Tareas realizadas con exito")
	//fmt.Printf("%d %s", a+b+c+d+e)
}
