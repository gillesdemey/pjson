package main

import (
  "os"
  "log"
  "encoding/json"
  "io/ioutil"
  // "github.com/mgutz/ansi"
)

func main() {

  in, _ := ioutil.ReadAll(os.Stdin)

  var f interface{}

  err := json.Unmarshal(in, &f)

  if err != nil {
    log.Fatal("invalid JSON structure")
  }

  m := f.(map[string]interface{})

  out, _ := json.MarshalIndent(m, "", "  ")

  os.Stdout.Write(out)
}