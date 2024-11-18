package application

import (
	"bufio"
	"log"
	"os"
	"strings"
	"github.com/w0ikid/rpn/pkg/rpn"
)

type Application struct {

}

func New() *Application {
	return &Application{}
}

func (a* Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)

		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}

		result, err := rpn.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}
