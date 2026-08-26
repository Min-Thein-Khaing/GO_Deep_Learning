package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Channel 1 မှ data ရောက်လာသည်"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Channel 2 မှ data ရောက်လာသည်"
	}()

	// Channel နှစ်ခုထဲမှ အရင်ဆုံးရောက်လာမည့်တစ်ခုကို စောင့်ယူခြင်း
	newSeleted := Selected(ch1, ch2, 3 * time.Second)
	fmt.Println(newSeleted)
}

func Selected(ch1, ch2 <-chan string,timeout time.Duration)string  {

	select {
	case msg1 := <-ch1:
		return msg1
	case msg2 := <-ch2:
		return msg2
	case <-time.After(3 * time.Second):
		return "Timeout"
	}

}