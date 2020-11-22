package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type DefaultSleeper struct {

}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	  for i := countdownStart; i > 0; i-- {
                fmt.Fprintln(out, i)
        }

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
