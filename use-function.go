import (
  "time"
  "sync"
) 
func yourfunction() //your function
  stopChan := make(chan struct{})
  var wg sync.WaitGroup
  wg.Add(1)

  // Start the animation in a goroutine
  go waiting1(stopChan, &wg)
