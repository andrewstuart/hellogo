// Package feed provides ...

package feed

import (
  "encoding/json"
)

type Item struct {
  Title string
  URL string
}

type Response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}
