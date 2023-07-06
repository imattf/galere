package main

import (
	"fmt"

	"github.com/imattf/galere/models"
)

func main() {
	gs := models.GalleryService{}
	fmt.Println(gs.Images(1))
	fmt.Println(gs.Images(2))

}
