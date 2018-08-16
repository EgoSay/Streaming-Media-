/**
 * @Author codeAC
 * @Time: 2018/8/13 20:37
 * @Description
 */
package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"

	VIDEO_PATH = "./videos/"
)

//控制流程
type controlChan chan string

//数据控制
type dataChan chan interface{}

//Dispatcher 和  Executor 控制(相当于生产者/消费者)
type fn func(dc dataChan) error
