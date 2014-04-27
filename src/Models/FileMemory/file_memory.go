package FileMemory

import "os"
import "fmt"
import "bufio"
import "path/filepath"
import "net/http"


func Create(filename string, request *http.Request) (string, error){
  err := prepareDirectory(filename)
  if err == nil { return "Couldn't prepare directory.", err }
  var file *os.File

  file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
  defer file.Close()
  if err =! nil { return "Couldn't open file.", err }

  addition := ToMemoryString.Do(request)
  _, err = file.WriteString(addition)
  if err != nil { return "Couldn't write to object.", err }
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
