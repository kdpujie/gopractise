package library

import (
	"fmt"
	"time"
)
type WAVPlayer struct {
	stat int
	progress int
}
func (w *WAVPlayer)Play(source string){
	fmt.Println("\tPlaying Wav music",source)
	w.progress = 0
	for w.progress < 100 {
		time.Sleep(100 * time.Millisecond) //睡眠100ms
		fmt.Print(".")
		w.progress += 10
	}
	fmt.Println("\n\tFinished playing",source)
	
}