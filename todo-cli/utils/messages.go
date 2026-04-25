package utils

import(
	"fmt"
	"text/tabwriter"
	"os"
)

func PrintHelpMessage(){
	w := tabwriter.NewWriter(os.Stdout,6,4,4,' ',0)
	fmt.Println("Uso da ferramenta - todo [args]... [props]...")
	fmt.Println("Possíveis Argumentos:")
	fmt.Fprintln(w,"\th, help\tMostra essa tela de ajuda.\t")
	fmt.Fprintln(w,"\tc, create\tAdiciona uma tarefa a lista To Do.\t")
	fmt.Fprintln(w,"\tr, read\tMostra todas as tarefas da lista To Do.\t")
	fmt.Fprintln(w,"\tu, update\tAtualiza uma das tarefa da lista To Do.\t")
	fmt.Fprintln(w,"\td, delete\tExclui uma tarefa a lista To Do.\t")
	w.Flush()
	fmt.Println("Ordem das Propriedades:")
	fmt.Fprintln(w,"\tOpção Usada\tParâmetros\t")
	fmt.Fprintln(w,"\tc, create\t'TÍTULO' [Está feita? s, sim | n, nao] \t")
	fmt.Fprintln(w,"\tr, read\t[NÃO HÁ NECESSIDADE DE PARÂMETROS]\t")
	fmt.Fprintln(w,"\tu, update\t[ID] [PROP=titulo, t | estado, e] [NOVA PROP]\t")
	fmt.Fprintln(w,"\td, delete\t[ID]\t")
	w.Flush()
	fmt.Println("Exemplos:")
	fmt.Println("\t$ tasks c 'Comprar bolo' n")
	fmt.Println("\t$ tasks r")
	fmt.Println("\t$ tasks u 1 t 'Comprar Bolo e Coca'")
	fmt.Println("\t$ tasks d 1")
}

func PrintWrongCommand(cmd []string){
	fmt.Printf("\033[31mCOMANDO INVÁLIDO! %s\nVEJA O MENU DE AJUDA 'tasks help'\033[0m\n",cmd)
}