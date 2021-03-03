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
Function | Description
---|---:
Info | Service status logging
Error | Default error logging
ErrorAndSend | Error logging and send discord message
Fatal | Very severe error, logging and the program terminated
Debug | Logging in develop environment only  
  
## Environment variable  
Name | Description
---|---:
LOGGER_ESAY_LEVEL_DISCORDURL | Discord Webhook URL for notice
LOGGER_ESAY_LEVEL_ISDEV | Determine if in development environment  