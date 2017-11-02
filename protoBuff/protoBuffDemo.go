package protoBuff

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func Test() {
	testLogInfo := &LoginInfo{
		AppToken: "xxxooo",
		UserId:   123456,
		DeviceId: "abcdefg",
	}

	protoLogInfo := genProto(testLogInfo)
	originLogInfo := reaProto(protoLogInfo)
	fmt.Println(fmt.Sprintf("origin log info:%+v \n", originLogInfo))
}

func genProto(testLogInfo *LoginInfo) []byte {
	pbOut, err := proto.Marshal(testLogInfo)
	if err != nil {
		println("generate proto error")
	}
	return pbOut
}

func reaProto(pbOut []byte) LoginInfo {
	pbLogin := LoginInfo{}
	if proto.Unmarshal(pbOut, &pbLogin) != nil {
		println("realease proto error")
	}

	return pbLogin
}
