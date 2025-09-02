package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()
	h := t.Hour()
	fmt.Println("The time is", h)
	if h < 12 {
		fmt.Println("Good morning")
	} else if h < 17 {
		fmt.Println("Good afternoon")
	} else {
		fmt.Println("Good evening")
	}

	fmt.Println("The time is", h)
	switch h {
	case 16:
		fmt.Println(("Happy hour starts now!"))
	case 17:
		fmt.Println("Happy hour is over")
	default:
		fmt.Println("Happy hour is coming soon")

	}
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("You are using Windows")
	case "linux":
		fmt.Println("You are using Linux")
	default:
		fmt.Println("You are using", os)
	}
}
