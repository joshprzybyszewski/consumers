package main

import (
	"github.com/joshprzybyszewski/consumable/snacks"
	"github.com/joshprzybyszewski/consumers/janitor"
	"github.com/joshprzybyszewski/consumers/lawoffice"
)

func main() {
	println(`i'm a consumer`)

	snacks.EatPringles()

	lawoffice.Printer()

	janitor.CleanPaper()
}
