package main

import "github.com/go-martini/martini"
import "fmt"
import "./src/FileLocator"


func main(){
  m:= martini.Classic()

  m.Get("/", func() (int, string) {
    return 200, "Goldfish is swimming. Ready for Goldfish memories."
  })

  m.Group("/memories", func(r martini.Router) {
    r.Post("/:kind/:id", CreateMemory)
    r.Get( "/:kind/:id", ShowMemory  )
  })

  m.Run()
}


func CreateMemory(params martini.Params) (int, string) {
  kind := params["kind"]
  id   := params["id"]
  path := FileLocator.Path(kind, id)

  return 422, fmt.Sprintf("Could not save memory. Path:%s", path)
}


func ShowMemory(params martini.Params) (int, string) {
  kind := params["kind"]
  id   := params["id"]
  path := FileLocator.Path(kind, id)

  return 422, fmt.Sprintf("Could not find memory. Path:%s", path)
}
