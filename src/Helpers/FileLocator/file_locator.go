package FileLocator

import "fmt"


func Path(kind string, id string) string{
  return fmt.Sprintf("/%s/%s.txt", kind, id)
}
