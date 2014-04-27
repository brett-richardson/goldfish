package main

import "github.com/go-martini/martini"
import "./src/Controllers"


func main(){
  m:= martini.Classic()

  //= Middleware ===
  m.Use(martini.Recovery())
  m.Use(martini.Logger())

  //= Routes ===
  m.Get("/", func() (int, string) {
    return 200, "Goldfish is swimming, ready for memories."
  })

  m.Group("/memories", func(r martini.Router) {
    r.Get( "/:kind/:id", Memories.Show  )
    r.Post("/:kind/:id", Memories.Create)
  })

  //= Run ===
  m.Run()
}
