package Memories

import "log"
import "strings"
import "net/http"
import "github.com/go-martini/martini"
import "../Models/FileMemory"
import "../Helpers/FileLocator"


func Create(request *http.Request, params martini.Params) (int, string) {
  location := FileLocator.Path(params["kind"], params["id"])
  _, err   := FileMemory.Create(location, request)

  if(err != nil){ return 422, "" }
  return 200, location
}


func Show(params martini.Params) (int, string) {
  location  := FileLocator.Path(params["kind"], params["id"])
  data, err := FileMemory.Fetch(location)
  log.Println("Showing memory: ", location)

  if(err != nil){ return 404, "" }
  return 200, strings.Join(data, "\n")
}
