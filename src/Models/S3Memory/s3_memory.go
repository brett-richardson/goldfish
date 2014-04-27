package S3Memory

import "os"
import "io"
import "log"
import "io/ioutil"
import "bytes"
import "net/http"
import "github.com/eaigner/s3"
import "../../Helpers/ToMemoryString"


func Create(filename string, request *http.Request) (string, error){
  obj          := s3Obj(filename)
  content, err := addToExistingContent(obj, ToMemoryString.Do(request))
  _, err = writeToObject(obj, content)
  return content, err
}


func Fetch(filename string) (string, error){
  reader, _, err := s3Obj(filename).Reader()
  if err != nil { return "", err }
  bytes, err := ioutil.ReadAll(reader)
  content := string(bytes)
  return content, err
}


// Private ---------------------------------------------------------------------

// If file exists, add existing content to content.
func addToExistingContent(obj s3.Object, content string) (string, error){
  exists, _      := obj.Exists()
  old_content, _ := contentsOf(obj)
  if exists == true { content = old_content + content }
  return content, nil
}


func writeToObject(obj s3.Object, content string) (string, error){
  writer := obj.Writer()
  defer writer.Close()
  _, err := io.Copy(writer, bytes.NewBufferString(content))
  return content, err
}

func contentsOf(obj s3.Object) (string, error){
  reader, _, err := obj.Reader()
  if err != nil {
    log.Println(obj, reader, err)
    return "", err
  }
  bytes, err := ioutil.ReadAll(reader)
  return string(bytes), err
}


func s3Connection() *s3.S3{
  s3c := &s3.S3{
    Bucket:    os.Getenv("S3_BUCKET"),
    AccessKey: os.Getenv("S3_KEY"   ),
    Secret:    os.Getenv("S3_SECRET"),
    Path:      os.Getenv("S3_PATH"  ),
  }
  return s3c
}


func s3Obj(filename string) s3.Object{
  return s3Connection().Object(filename)
}
