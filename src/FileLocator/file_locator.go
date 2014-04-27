package FileLocator

import "fmt"

func Path(kind string, id string) string{
  return fmt.Sprintf("store/%s/%s", kind, id)
}
