package controler

import (
	"fmt"
	"time"


)

//PegarDataAtual serve para converter uma data para string com formato
func PegarDataAtualString() string {
	timeString := time.Now().Format("02/01/2006 03:04:05")
	fmt.Println("Teste impressão", timeString)
	return timeString
}
