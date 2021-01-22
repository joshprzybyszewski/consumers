package janitor

import (
	"fmt"

	"github.com/joshprzybyszewski/consumable/papers"
	"github.com/joshprzybyszewski/consumable/toiletpaper"
)

func CleanPaper() {
	width, height := papers.TPSize(toiletpaper.GetRandomBrand())
	fmt.Printf("cleaning a paper that is %v x %v large\n", width, height)
}
