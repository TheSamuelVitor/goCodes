package aluno

import (
	"fmt"
	"programaCLI/data"
)

func GetAlunos() {
	for _, aluno := range data.Alunos {
		fmt.Printf("\t\n Id: %d\t\n Nome: %s\t\n Idade: %d\n", aluno.Id, aluno.Nome, aluno.Idade)
	}
}
