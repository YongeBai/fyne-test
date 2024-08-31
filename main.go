package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello")
	in := widget.NewEntry()

	in.OnChanged = func(context string) {
		out.SetText("Hello " + context)
	}

	return out, in
}

func updateClock(clock *widget.Label) {
	format := time.Now().Format("15:04:05")
	clock.SetText(format)
}

func tidyUp() {
	fmt.Println("tidyUp")
}

func main() {
	a := app.New()
	w1 := a.NewWindow("entry")
	w1.SetContent(container.NewVBox(makeUI()))
	w1.SetMaster()
	w1.Show()

	w2 := a.NewWindow("clock")
	w2.SetContent(widget.NewLabel(""))
	go func() {
		for range time.Tick(time.Second) {
			updateClock(w2.Content().(*widget.Label))
		}
	}()
	w2.Show()	

	a.Run()
	tidyUp()
}

