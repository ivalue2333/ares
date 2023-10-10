package ares

func TaskRunMany(total, batchSize int, taskFunc func() error) {
	Info("TaskRunMany start")
	execSize := 0
	for j := 0; j < total/batchSize; j++ {
		err := taskFunc()
		if err != nil {
			Info("TaskRunMany err:%+v", err)
			return
		}
		execSize += batchSize
		Info("TaskRunMany percent:%d/%d, %d%s.", execSize, total, execSize*100/total, "%")
	}
	Info("TaskRunMany done")
}
