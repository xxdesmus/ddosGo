package ddosgo

import (
  "os/signal"
  "syscall"
  "os"
  "fmt"
)

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- ddosGo REPORT for ",*URL)
    fmt.Println("\r- Total Requests: ",*Totalrequest)
    fmt.Println("\r- Total Response: ",*Totalresponse)
    fmt.Println("\r- Total Bounced: ",*Totalbounced)
    fmt.Println("\r- Error Distribution")
    for a,b := range(Errors){
      fmt.Printf("\r    - %v ---- %v\n",a,b)
    }
    fmt.Printf("\r- Response Distribution\n")
    for c,d := range(Responses){
      fmt.Printf("\r    - %v ---- %v\n",c,d)
    }
		os.Exit(0)
	}()
}
