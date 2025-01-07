// Package fundamentos_da_linguagem demonstrates the basic data types in Go.
// It includes examples of integer, unsigned integer, floating-point, complex,
// boolean, string, and byte types, showcasing their usage and formatting.
package fundamentals_of_language

import "fmt"

// TiposDeDados demonstrates the declaration and usage of various data types in Go.
// It includes examples of integer types (int64, int32), unsigned integer types (uint64, uint32),
// floating-point types (float64, float32), complex number types (complex64, complex128),
// boolean type (bool), string type (string), and byte type (byte).
// The function returns a formatted string that shows the value and type of each variable.
func TypeOfData() string {
	var IntNumber64 int64 = 1000000000000000000
	var IntNumber32 int32 = 1000000000

	var numberUintero64 uint64 = 1000000000000000000
	var numberUintero32 uint32 = 1000000000

	var numberReal64 float64 = 1000000000000000000.0
	var numberReal32 float32 = 1000000000.0

	var numberComplex complex64 = 1 + 1i
	var numberComplex2 complex128 = 1 + 1i

	var booleano bool = true

	var texto string = "Texto"

	var caracter byte = 'A'

	variavel := fmt.Sprintf(
		"\nIntNumber64: %d :: %T\nIntNumber32: %d :: %T\nnumeroUintero64: %d :: %T\nnumeroUintero32: %d :: %T\nnumeroReal64: %f :: %T\nnumeroReal32: %f :: %T\nnumeroComplexo: %f :: %T\nnumeroComplexo2: %f :: %T\nbooleano: %t :: %T\ntexto: %s :: %T\ncaracter: %c :: %T",
		IntNumber64, IntNumber64, IntNumber32, IntNumber32, numberUintero64, numberUintero64, numberUintero32, numberUintero32, numberReal64, numberReal64, numberReal32, numberReal32, numberComplex, numberComplex, numberComplex2, numberComplex2, booleano, booleano, texto, texto, caracter, caracter,
	)

	return variavel
}
