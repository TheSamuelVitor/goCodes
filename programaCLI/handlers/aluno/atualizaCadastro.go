package aluno

import (
	"fmt"
	"programaCLI/data"
	"programaCLI/structs"
)

func AtualizaCadastro() {

	var id int
	fmt.Printf("Id: ")
	fmt.Scanln(&id)

	novoAluno := structs.Aluno{}
	fmt.Printf(" --- ATUALIZAR CADASTRO ---\n")
	fmt.Printf("Nome: ")
	fmt.Scanln(&novoAluno.Nome)
	fmt.Printf("Idade: ")
	fmt.Scanln(&novoAluno.Idade)
	novoAluno.Id = id

	data.Alunos[id] = novoAluno
}
