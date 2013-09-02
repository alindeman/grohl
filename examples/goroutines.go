package main

import (
  ".."
  "strconv"
  "runtime"
  "fmt"
)

// quick test for concurrent log messages
func main() {
  fmt.Printf("%d -> %d\n", runtime.NumCPU, runtime.GOMAXPROCS(1000))

  logger := scrolls.NewLogger(nil)
  chans := make([]chan bool, 1000)

  for i := range chans {
    chans[i] = make(chan bool)
    go LogLikeCrazy(strconv.Itoa(i), logger, chans[i])
  }

  for _, ch := range chans {
    <- ch
  }
}

func LogLikeCrazy(key string, logger scrolls.Logger, stopped chan bool) {
  for i := 0; i < 10; i += 1 {
    logger.Log(scrolls.LogData{"key": key, "i": strconv.Itoa(i+1)})
  }
  stopped <- true
}