package aluno

import (
	"fmt"
	"programaCLI/data"
	"programaCLI/structs"
)

func CadastroAluno() {
	var novoAluno = &structs.Aluno{}
	novoAluno.Id = len(data.Alunos) + 1

	fmt.Printf("Nome: ")
	fmt.Scanln(&novoAluno.Nome)
	fmt.Printf("Idade: ")
	fmt.Scanln(&novoAluno.Idade)

	data.Alunos = append(data.Alunos, *novoAluno)
}
