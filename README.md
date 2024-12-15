# loading
loading process in go

copy and paste this:

```
func waiting1(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	symbols := []string{
		"|", "||", "|||", "||", "|",
		"/", "//", "///", "//", "/",
		"-", "--", "---", "--", "-",
		"\\", "\\\\", "\\\\\\", "\\\\", "\\",
	}
	for {
		for _, symbol := range symbols {
			select {
			case <-stopChan:
				return 
			default:
				fmt.Printf("\rLoading... %s", symbol)
				time.Sleep(300 * time.Millisecond) // lower the time to make it slower
			}
		}
	}
}
```

and if you want to use it, copy and paste this:

```
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
```
