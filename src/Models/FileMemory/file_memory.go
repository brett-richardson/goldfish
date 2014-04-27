package FileMemory

import "os"
import "fmt"
import "bufio"
import "path/filepath"
import "net/http"


func Create(filename string, request *http.Request) (string, error){
  memory    := request.PostFormValue("memory")
  datestamp := request.PostFormValue("datetime")
  err       := prepareDirectory(filename)
  var file *os.File

  if err == nil {
    file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
    defer file.Close()
  }

  if err == nil {
    addition := "==== " + datestamp + " ====" + "\n" + memory + "\n\n"
    _, err = file.WriteString(addition)
  }

  if err != nil { fmt.Println(err); return "", err }
  return memory, nil
}


func Fetch(filename string) ([]string, error){
  data, err := readLines(filename)
  if err != nil { fmt.Println(err); return nil, err }

  return data, nil
}


// Private ---------------------------------------------------------------------


func prepareDirectory(path string) error {
  dir, _ := filepath.Split(path)
  err := os.MkdirAll(dir, 0777)
  return err
}


func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil { return nil, err }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
