# easy-logging-level
Easy to design logging level with standard logger

## Getting Started
To install the library, run:  
```
go get github.com/im-jinsu/easy-logging-level/logger
```  
The following is a simple example which logging with standard logger
```go
package main

import (
	"log"

	"github.com/im-jinsu/easy-logging-level/logger"
)

func main() {
	// Set stdlogger with lumberjack
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger.Info("hello world")
}
```  
## Level  
값 | 의미
---|---:
Info | 
Error | 
ErrorAndSend | 
Fatal | 
Debug |  