package main

import (

	//"bufio"
	"fmt"
	//"os"
)


func main() {

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

	fmt.Println("----------------------------------------------------")
	fmt.Println("Ejercicio 2")
	fmt.Println("Escribe un numero.")
	var w int
	fmt.Scanln(&w)
	fmt.Printf("Tu numero es: %d\n ", w)

	fmt.Println("----------------------------------------------------")
	fmt.Println("ejercicio 3")
	fmt.Println("Escribe tu primer num")
	var x, y int
	fmt.Scan(&x)
	fmt.Println("Escribe tu segundo num.")
	fmt.Scan(&y)
	fmt.Printf("total %d\n", x+y)

	fmt.Println("----------------------------------------------------")
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

	fmt.Println("----------------------------------------------------")
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
}*/

func main() {
	var usuario, contra string
	fmt.Println("Ingrese su usuario:")
	fmt.Scanln(&usuario)
	fmt.Println("tu usuario es:", usuario)
	fmt.Println("Ingrese contraseña")
	fmt.Scanln(&contra)
	/*/if contra == "hola" {
			fmt.Println("bienvenido", usuario)
		} else {
			for x := contra ;
			x == contra;
			fmt.scan(&contra)

			}
			fmt.Println("Contraseña incorrecta")
		}


	}*/
	if contra == "hola" {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
	for x := 0; x == 1; 
		if contra == "hola" {
		fmt.Println("bienvenido", usuario)
	} else {
		fmt.Println("Contraseña incorrecta") {
		
		}
	}
}
