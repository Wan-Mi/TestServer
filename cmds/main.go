package main

import (
	"fmt"
	"sort"
	"time"
)

//使用
func main() {

	//jwt.Test()
	//iInterface.Test()
	//protoBuff.Test()
	//router.MakeRouter()

	//	syncGoroutine.ChanTests()

	getTimeWindowTest()
}

func getTimeWindowTest() {
	valMap := map[int64]string{
		1511236559: "aaaa",
		1511238459: "bbbb",
	}
	var timeStamps []int
	for key, _ := range valMap {
		timeStamps = append(timeStamps, int(key))
	}
	sort.Ints(timeStamps)

	getWindowDeviceID := func(tSs []int, tS int) string {
		if len(tSs) == 0 {
			return ""
		}

		tSs = append(tSs, int(time.Now().Unix()+1))
		for idx, t := range tSs {

			if idx > 0 && t > tS {
				key := int64(tSs[idx-1])
				return valMap[key]
			}
		}
		return ""
	}

	fmt.Println(getWindowDeviceID(timeStamps, int(time.Now().Unix())))
}
