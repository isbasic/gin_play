package main

import (
	"fmt"

	routes "github.com/isbasic/gin_play/group-routes"
	_ "github.com/isbasic/gin_play/sample"
)

func setPort(i int) string {
	return fmt.Sprintf(":%d", i)
}

func main() {
	//r := s.SetupRouter()
	r := routes.GetRoutes()
	r.Run(setPort(8880))
}
