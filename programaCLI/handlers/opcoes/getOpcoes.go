package opcoes

import (
	"fmt"
	"programaCLI/data"
)

func GetOpcoes() {
	fmt.Printf("--- MENU ---\n")
	for _, opcao := range data.Opcoes {
		fmt.Printf("%d. %s\n", opcao.Id, opcao.Descricao)
	}
}
