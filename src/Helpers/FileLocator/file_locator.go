package FileLocator

import "fmt"
import "log"

func Path(kind string, id string) string{
  location := fmt.Sprintf("./store/%s/%s.txt", kind, id)
  log.Println("Location generated: ", location)
  return location
}
