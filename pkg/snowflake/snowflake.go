package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	if st, err = time.Parse("2006-01-02", startTime); err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000 // 设置时间戳开始时间
	node, err = sf.NewNode(machineId)
	return
}

func GenInt64Id() int64 {
	return node.Generate().Int64()
}
