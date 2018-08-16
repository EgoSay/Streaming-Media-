/**
 * @Author codeAC
 * @Time: 2018/8/15 13:56
 * @Description
 */
package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT into video_del_rec (video_id) value(?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}
