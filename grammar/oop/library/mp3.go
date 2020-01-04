package library

import (
	"fmt"
	"time"
)
type MP3Player struct {
	stat int
	progress int
}
func (p *MP3Player)Play(source string){
	fmt.Println("\tPlaying MP3 music",source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) //睡眠100ms
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\n\tFinished playing",source)
}