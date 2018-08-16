/**
 * @Author codeAC
 * @Time: 2018/8/13 21:29
 * @Description
 */
package dbops

import "log"

//查找要删除的资源
func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmOut, err := dbConn.Prepare("select video_id from video_del_rec limit ?")
	var ids []string
	if err != nil {
		return ids, err
	}
	rows, err := stmOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error: %v", err)
		return ids, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stmOut.Close()
	return ids, nil
}

//删除资源记录
func DelVideoDeletionRecord(vid string) error {
	stmDel, err := dbConn.Prepare("DELETE from video_del_rec where video_id = ?")
	if err != nil {
		return err
	}
	_, err = stmDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord error: %v", err)
	}
	defer stmDel.Close()
	return nil
}
