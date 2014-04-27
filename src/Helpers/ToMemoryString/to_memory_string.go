package ToMemoryString

import "net/http"


func Do(request *http.Request) string{
  memory    := request.PostFormValue("memory")
  datestamp := request.PostFormValue("datetime")
  addition  := "==== " + datestamp + " ====" + "\n" + memory + "\n\n"
  return addition
}
