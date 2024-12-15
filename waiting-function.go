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
				time.Sleep(300 * time.Millisecond) 
			}
		}
	}
}
