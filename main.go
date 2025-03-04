package main

import (
	"fmt"
	"os"

	"howett.net/plist"
)

type iTermColor struct {
	Background struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Background Color"`
	Foreground struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Foreground Color"`
	Cursor struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Cursor Color"`
	CursorText struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Cursor Text Color"`
	Selection struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Selection Color"`
	SelectedText struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Selected Text Color"`
	Bold struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Bold Color"`
	// black
	ANSI0 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 0 Color"`
	// red
	ANSI1 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 1 Color"`
	// green
	ANSI2 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 2 Color"`
	// yellow
	ANSI3 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 3 Color"`
	// blue
	ANSI4 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 4 Color"`
	// magenta
	ANSI5 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 5 Color"`
	// cyan
	ANSI6 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 6 Color"`
	// white
	ANSI7 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 7 Color"`
	// BRIGHT VARIANTS (as above)
	ANSI8 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 8 Color"`
	ANSI9 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 9 Color"`
	ANSI10 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 10 Color"`
	ANSI11 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 11 Color"`
	ANSI12 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 12 Color"`
	ANSI13 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 13 Color"`
	ANSI14 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 14 Color"`
	ANSI15 struct {
		Red   float64 `plist:"Red Component"`
		Green float64 `plist:"Green Component"`
		Blue  float64 `plist:"Blue Component"`
		Alpha int64   `plist:"Alpha Component"`
	} `plist:"Ansi 15 Color"`
}

func iterm2hex(R float64, G float64, B float64) string {
	R, G, B = R*255, G*255, B*255
	return fmt.Sprintf("%x%x%x", int(R), int(G), int(B))
}

func main() {
	// init our variable to hold the values
	var itermcolors iTermColor
	// read a file
	// TODO: make this import from os.args
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		err := fmt.Errorf("No file specified.\nUsage: iterm2hex <file>")
		panic(err)
	}
	file := argsWithoutProg[0]
	buf, _ := os.Open(file)
	defer buf.Close()

	// now we read and marshall the plist entries
	decoder := plist.NewDecoder(buf)
	err := decoder.Decode(&itermcolors)
	if err != nil {
		panic(err)
	}

	// convert everything to hex
	// TODO: find a shorthand for this
	bg := iterm2hex(itermcolors.Background.Red, itermcolors.Background.Green, itermcolors.Background.Blue)
	fg := iterm2hex(itermcolors.Foreground.Red, itermcolors.Foreground.Green, itermcolors.Foreground.Blue)
	ansi0 := iterm2hex(itermcolors.ANSI0.Red, itermcolors.ANSI0.Green, itermcolors.ANSI0.Blue)
	ansi1 := iterm2hex(itermcolors.ANSI1.Red, itermcolors.ANSI1.Green, itermcolors.ANSI1.Blue)
	ansi2 := iterm2hex(itermcolors.ANSI2.Red, itermcolors.ANSI2.Green, itermcolors.ANSI2.Blue)
	ansi3 := iterm2hex(itermcolors.ANSI3.Red, itermcolors.ANSI3.Green, itermcolors.ANSI3.Blue)
	ansi4 := iterm2hex(itermcolors.ANSI4.Red, itermcolors.ANSI4.Green, itermcolors.ANSI4.Blue)
	ansi5 := iterm2hex(itermcolors.ANSI5.Red, itermcolors.ANSI5.Green, itermcolors.ANSI5.Blue)
	ansi6 := iterm2hex(itermcolors.ANSI6.Red, itermcolors.ANSI6.Green, itermcolors.ANSI6.Blue)
	ansi7 := iterm2hex(itermcolors.ANSI7.Red, itermcolors.ANSI7.Green, itermcolors.ANSI7.Blue)
	ansi8 := iterm2hex(itermcolors.ANSI8.Red, itermcolors.ANSI8.Green, itermcolors.ANSI8.Blue)
	ansi9 := iterm2hex(itermcolors.ANSI9.Red, itermcolors.ANSI9.Green, itermcolors.ANSI9.Blue)
	ansi10 := iterm2hex(itermcolors.ANSI10.Red, itermcolors.ANSI10.Green, itermcolors.ANSI10.Blue)
	ansi11 := iterm2hex(itermcolors.ANSI11.Red, itermcolors.ANSI11.Green, itermcolors.ANSI11.Blue)
	ansi12 := iterm2hex(itermcolors.ANSI12.Red, itermcolors.ANSI12.Green, itermcolors.ANSI12.Blue)
	ansi13 := iterm2hex(itermcolors.ANSI13.Red, itermcolors.ANSI13.Green, itermcolors.ANSI13.Blue)
	ansi14 := iterm2hex(itermcolors.ANSI14.Red, itermcolors.ANSI14.Green, itermcolors.ANSI14.Blue)
	ansi15 := iterm2hex(itermcolors.ANSI15.Red, itermcolors.ANSI15.Green, itermcolors.ANSI15.Blue)
	cursor := iterm2hex(itermcolors.Cursor.Red, itermcolors.Cursor.Green, itermcolors.Cursor.Blue)
	cursorText := iterm2hex(itermcolors.CursorText.Red, itermcolors.CursorText.Green, itermcolors.CursorText.Blue)
	selection := iterm2hex(itermcolors.Selection.Red, itermcolors.Selection.Green, itermcolors.Selection.Blue)
	selectedText := iterm2hex(itermcolors.SelectedText.Red, itermcolors.SelectedText.Green, itermcolors.SelectedText.Blue)
	bold := iterm2hex(itermcolors.Bold.Red, itermcolors.Bold.Green, itermcolors.Bold.Blue)

	// print the results
	fmt.Println("Name                  Value")
	fmt.Println("----                  -----")
	fmt.Printf("%s            #%s\n", "Background", bg)
	fmt.Printf("%s            #%s\n", "Foreground", fg)
	fmt.Printf("%s             #%s\n", "0 (black)", ansi0)
	fmt.Printf("%s               #%s\n", "1 (red)", ansi1)
	fmt.Printf("%s             #%s\n", "2 (green)", ansi2)
	fmt.Printf("%s            #%s\n", "3 (yellow)", ansi3)
	fmt.Printf("%s              #%s\n", "4 (blue)", ansi4)
	fmt.Printf("%s           #%s\n", "5 (magenta)", ansi5)
	fmt.Printf("%s              #%s\n", "6 (cyan)", ansi6)
	fmt.Printf("%s             #%s\n", "7 (white)", ansi7)
	fmt.Printf("%s      #%s\n", "8 (bright black)", ansi8)
	fmt.Printf("%s        #%s\n", "9 (bright red)", ansi9)
	fmt.Printf("%s     #%s\n", "10 (bright green)", ansi10)
	fmt.Printf("%s    #%s\n", "11 (bright yellow)", ansi11)
	fmt.Printf("%s      #%s\n", "12 (bright blue)", ansi12)
	fmt.Printf("%s   #%s\n", "13 (bright magenta)", ansi13)
	fmt.Printf("%s      #%s\n", "14 (bright cyan)", ansi14)
	fmt.Printf("%s     #%s\n", "15 (bright white)", ansi15)
	fmt.Printf("%s                #%s\n", "Cursor", cursor)
	fmt.Printf("%s           #%s\n", "Cursor Text", cursorText)
	fmt.Printf("%s             #%s\n", "Selection", selection)
	fmt.Printf("%s         #%s\n", "Selected Text", selectedText)
	fmt.Printf("%s                  #%s\n", "Bold", bold)
}
