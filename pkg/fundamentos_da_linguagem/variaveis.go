package fundamentos_da_linguagem

import "fmt"

func Variaveis() string {
	var var1 string = "Variável 1"
	var2 := "Variável 2"

	var (
		var3 string = "Variável 3"
		var4 string = "Variável 4"
	)

	var5, var6 := "Variável 5", "Variável 6"

	const constante1 string = "Constante 1"

	variavel := fmt.Sprintf(
		"\nvar1: %s :: %T\nvar2: %s :: %T\nvar3: %s :: %T\nvar4: %s :: %T\nvar5: %s :: %T\nvar6: %s :: %T\nconstante1: %s :: %T\n",
		var1, var1, var2, var2, var3, var3, var4, var4, var5, var5, var6, var6, constante1, constante1,
	)

	return variavel
}
