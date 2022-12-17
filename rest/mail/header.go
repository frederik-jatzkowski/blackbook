package mail

import (
	"fmt"
	"sync"
	"time"
)

var messageIdCounter = 0
var messageIdCounterLock sync.Mutex

func buildHeaders(from string, to string, subject string, host string) string {
	messageIdCounterLock.Lock()
	defer messageIdCounterLock.Unlock()

	messageIdCounter++

	return fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Message-ID: <%d-%d@%s>\r\n"+
			"Subject: Aktivierungscode\r\n"+
			"\r\n",
		from,
		to,
		messageIdCounter,
		time.Now().UnixNano(),
		host,
	)
}
