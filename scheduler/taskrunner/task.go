/**
 * @Author codeAC
 * @Time: 2018/8/13 20:37
 * @Description
 */
package taskrunner

import (
	"errors"
	"log"
	"os"
	"project/project/videoServer/scheduler/dbops"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
	}
	return nil
}
func VideoClearDispatcher(dc dataChan) error {
	record, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error:%v", err)
		return err
	}
	//表示没有
	if len(record) == 0 {
		return errors.New("All tasks finished ")
	}
	for _, id := range record {
		dc <- id
	}
	return nil
}
func VideoClearExecutor(dc dataChan) error {
	//错误集合
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			//为什么不直接用传进来的vid??
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
		errMap.Range(func(key, value interface{}) bool {
			err = value.(error)
			if err != nil {
				return false
			}
			return true
		})
	}
	return err
}
