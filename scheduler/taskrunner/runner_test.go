/**
 * @Author codeAC
 * @Time: 2018/8/13 21:13
 * @Description
 */
package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TestRunner_StartAll(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}
		return nil
	}
	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor Received: % v", d)
			default:
				break forloop
			}
		}
		return nil
	}
	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
