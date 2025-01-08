package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"github.com/charmbracelet/lipgloss"
)

func main() {

	resp, err := http.Get("https://zenquotes.io/api/today")

	if err != nil {
		log.Fatal(err)
	}

	// must close body after request
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var res []map[string]interface{}
	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatal(err)
	}

	var styleA = lipgloss.NewStyle().
		Bold(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("63")).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#000000")).
		PaddingTop(2).
		PaddingBottom(2).
		PaddingLeft(2).
		PaddingRight(2).
		Width(50)

	quote := res[0]["q"].(string)
	author := res[0]["a"].(string)

	fmt.Println(styleA.Render(quote, "\n", "\n", "Author - ", author, "\n", "\n", "https://zenquotes.io"))
	fmt.Println("")
}
