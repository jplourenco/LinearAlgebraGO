package main

import "fmt"

type Matrix [][]int

func eye(x int) Matrix{
	
	mat := make([][]int,x)

	for i := 0 ; i < x ;i++ {
		mat[i] = make([]int,x)
		mat[i][i] = 1
	}
	return mat
}

func dot(a,b []int ) int{

	l := len(a)
	res := 0
	for i := 0; i < l; i++{
		res += a[i]*b[i]
	}
	return res
}

func transpose(a Matrix) Matrix{
	
	mat := make([][]int,len(a[0]))

	for i:= 0 ; i < len(a[0]); i++{
		mat[i] = make([]int,len(a))
		for j := 0; j < len(a);j++{
			mat[i][j] = a[j][i]
		}
	}
	return mat
}

func mult(a,b Matrix) Matrix{

	if len(a[0]) != len(b) {
		return nil
	}

	tb := transpose(b)

	res := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		res[i] = make([]int,len(b[0]))
		for j := 0; j < len(b[0]); j++{
			res[i][j] = dot(a[i],tb[j])
		}
	}
	return res
}

func scalarMult(a Matrix, n int) Matrix{

	mat := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		mat[i] = make([]int,len(a[0]))
		if n != 0 {
			for j := 0; j < len(a[0]);j++{
				mat[i][j] = n*a[i][j]
			}
		}
	}
	return mat
}

func add(a,b Matrix) Matrix{

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil
	}

	mat := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		mat[i] = make([]int,len(a[0]))
		for j := 0; j < len(a[0]); j++{
			mat[i][j] = a[i][j] + b[i][j]
		}
	}

	return mat
}

func sub(a,b Matrix) Matrix{

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil
	}

	mat := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		mat[i] = make([]int,len(a[0]))
		for j := 0; j < len(a[0]); j++{
			mat[i][j] = a[i][j] - b[i][j]
		}
	}

	return mat
}

func eMult(a,b Matrix) Matrix{
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil
	}

	mat := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		mat[i] = make([]int,len(a[0]))
		for j := 0; j < len(a[0]); j++{
			mat[i][j] = a[i][j] * b[i][j]
		}
	}

	return mat
}

func eDiv(a,b Matrix) Matrix{
	
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return nil
	}

	mat := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		mat[i] = make([]int,len(a[0]))
		for j := 0; j < len(a[0]); j++{
			mat[i][j] = a[i][j] / b[i][j]
		}
	}
	return mat
}

func main() {
	
	a := []int{1,2}
	b := []int{4,5}
	
	c:= []int{4,5,3}	

	z := [][]int{a,b}
	y := [][]int{c,c}

	//fmt.Println(z)
	fmt.Println(scalarMult(eye(4),10))
	fmt.Println(eDiv(z,y))
	fmt.Println(eDiv(z,z))



}

