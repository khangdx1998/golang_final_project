package main

import (
	"final_project/router"
)


func main() {
	r := router.SetUpRouter()
	r.Run(":5000")
}





