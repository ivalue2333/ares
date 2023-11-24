package runs

import "github.com/ivalue2333/ares/logs"

func TaskRunMany(total, batchSize int, taskFunc func() error) {
	logs.Info("TaskRunMany start")
	execSize := 0
	for j := 0; j < total/batchSize; j++ {
		err := taskFunc()
		if err != nil {
			logs.Info("TaskRunMany err:%+v", err)
			return
		}
		execSize += batchSize
		logs.Info("TaskRunMany percent:%d/%d, %d%s.", execSize, total, execSize*100/total, "%")
	}
	logs.Info("TaskRunMany done")
}
