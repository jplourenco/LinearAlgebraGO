
package linalgo

type Matrix [][]int

func Eye(x int) Matrix{
	
	mat := make([][]int,x)

	for i := 0 ; i < x ;i++ {
		mat[i] = make([]int,x)
		mat[i][i] = 1
	}
	return mat
}

func Dot(a,b []int ) int{

	l := len(a)
	res := 0
	for i := 0; i < l; i++{
		res += a[i]*b[i]
	}
	return res
}

func Transpose(a Matrix) Matrix{
	
	mat := make([][]int,len(a[0]))

	for i:= 0 ; i < len(a[0]); i++{
		mat[i] = make([]int,len(a))
		for j := 0; j < len(a);j++{
			mat[i][j] = a[j][i]
		}
	}
	return mat
}

func Mult(a,b Matrix) Matrix{

	if len(a[0]) != len(b) {
		return nil
	}

	tb := Transpose(b)

	res := make([][]int,len(a))

	for i := 0; i < len(a); i++ {
		res[i] = make([]int,len(b[0]))
		for j := 0; j < len(b[0]); j++{
			res[i][j] = Dot(a[i],tb[j])
		}
	}
	return res
}

func ScalarMult(a Matrix, n int) Matrix{

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

func Add(a,b Matrix) Matrix{

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

func Sub(a,b Matrix) Matrix{

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


