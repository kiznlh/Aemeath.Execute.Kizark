package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/muesli/termenv"

	"Aemeath.Execute.Me/utils"
)

func main() {
	// Play music
	musicChan, err := utils.PlayMusic()
	if err != nil {
		log.Fatal(err)
	}
	////////

	output := termenv.NewOutput(os.Stdout)
	s := output.String("hehehe")

	utils.SlowPrint(s, 100*time.Millisecond)
	fmt.Println()
	output.CursorUp(1)

	s = s.CrossOut().Reverse()

	utils.SlowPrint(s, 10*time.Millisecond)

	// Play till end of music
	<-musicChan
}
