# zenquotes
Quote generator written in Go. Grabs quotes from https://zenquotes.io \
Uses charmbracelet/lipgloss library for pretty output https://github.com/charmbracelet/lipgloss 

Right now it only prints the quote of the day. I'm planning to add the ability to grab random quotes, and quotes by person soon. Complete with command line switches and everything.

## Quick Start

git clone https://github.com/lowerclasswarfare/zenquotes.git \
go get https://github.com/charmbracelet/lipgloss \
go build zenquotes.go \
./zenquotes \
??? \
profit
