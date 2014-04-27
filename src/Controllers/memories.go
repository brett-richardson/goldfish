package Memories

import "log"
import "net/http"
import "github.com/go-martini/martini"

import "../Models/S3Memory"
import "../Helpers/FileLocator"


func Create(request *http.Request, params martini.Params) (int, string) {
  location := FileLocator.Path(params["kind"], params["id"])
  _, err   := S3Memory.Create(location, request)

  if(err != nil){ return 422, "" }
  return 200, location
}


func Show(params martini.Params) (int, string) {
  location  := FileLocator.Path(params["kind"], params["id"])
  content, err := S3Memory.Fetch(location)
  log.Println("Showing memory: ", location)
  if(err != nil){ return 404, "File not found." }
  return 200, content
}
