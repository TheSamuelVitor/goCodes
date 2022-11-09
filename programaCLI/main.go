package main

import (
	"fmt"
	"programaCLI/handlers/aluno"
	"programaCLI/handlers/opcoes"
)

func main() {
	var opcao int = 1
	for opcao != 0 {
		opcoes.GetOpcoes()
		fmt.Scanln(&opcao)
		switch opcao {
		case 1:
			aluno.GetAlunos()
		case 2:
			aluno.CadastroAluno()
		case 3:
			aluno.AtualizaCadastro()
		case 0:
			fmt.Println("tchau")
		default:
			fmt.Println("funcao nao cadastrada")
		}
	}
}
