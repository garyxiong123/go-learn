package sleep

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	fmt.Println("Starting...")
	time.Sleep(5 * time.Second)
	fmt.Println("Finished.")
}
