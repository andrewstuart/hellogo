package main

import (
  "log"
  "net/http"
  "encoding/json"
  "fmt"
  "os"
)

//Item is a reddit item
type Item struct {
  Title string
  URL string
  Score int
  Author string
}

//Response is neat.
type Response struct {
  Data struct {
    Children []struct {
      Data Item
    }
    After string
  }
}

func main() {

  var subreddit string

  if(len(os.Args) >= 2) {
    subreddit = os.Args[1]
  } else {
    fmt.Println("Usage:")
    message := fmt.Sprintf("%s [subreddit] [after]", os.Args[0])
    fmt.Println(message)
    return
  }

  str := fmt.Sprintf("http://reddit.com/r/%s.json?count=20", subreddit)

  if(len(os.Args) >= 3) {
    str = str + "&after=" + os.Args[2]
  }

  resp, err := http.Get(str)

  if(err != nil) {
    log.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    log.Fatal(resp.Status)
  }

  r := new(Response)

  err = json.NewDecoder(resp.Body).Decode(r)

  if err != nil {
    log.Fatal(err)
  }

  var i = 0
  for _, child := range r.Data.Children {
    i = i + 1
    ret := fmt.Sprintf("%d. (%d) /u/%s: %s", i, child.Data.Score, child.Data.Author, child.Data.Title)
    fmt.Println(ret)
  }
  fmt.Println(r.Data.After)
}
