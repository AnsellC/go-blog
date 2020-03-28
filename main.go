package main

import (
	"os"
)

func main() {
	a := App{}
	a.Init()
	a.Serve(os.Getenv("APP_HTTP_PORT"))
}
