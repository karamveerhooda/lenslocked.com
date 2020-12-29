package main

import "fmt"

func main() {

	var x int
	var y string
	var z bool
	//var a int16
	//var b int32
	//var c int64

	x = 42
	y = "James Bond"
	z = true
	/*a = 10
	b = 10
	c = 10*/

	fmt.Println("x: ", x, "\ty: ", y, "\tz: ", z)
	fmt.Printf("\n%T\t%T\t%T", x, y, z)

}
