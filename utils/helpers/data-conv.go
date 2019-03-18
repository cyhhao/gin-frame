package helpers

import (
	"strconv"
	"fmt"
	"math/big"
	"time"
	"reflect"
)

func HidePhoneNumber(mobile interface{}) (string) {
	mobileStr, isStr := mobile.(string)
	if !isStr {
		mobileStr = strconv.FormatInt(mobile.(int64), 10)
	}
	if len(mobileStr) < 11 {
		return mobileStr
	}
	return fmt.Sprintf("%s****%s", mobileStr[0:3], mobileStr[7:11])
}

func String2BigInt(numStr string) (*big.Int, bool) {
	n := new(big.Int)
	return n.SetString(numStr, 10)
}

const TimeFormatMs = "2006-01-02T15:04:05.000"
const TimeFormatS = "2006-01-02T15:04:05"
const TimeFormatSpaceS = "2006-01-02 15:04:05"

func TimeStrToStamp(format string, timeStr string) (int64, error) {
	thisTime, err := time.Parse(format, timeStr)
	if err != nil {
		return 0, err
	}
	return thisTime.Unix(), nil

}

func TimeStrToTime(format string, timeStr string) (time.Time, error) {
	return time.Parse(format, timeStr)

}

func ExtractList(list interface{}, key string) (interface{}) {
	var ans = make([]string, 0)
	data := reflect.ValueOf(list)
	for i := 0; i < data.Len(); i++ {
		item := data.Index(i).Interface()
		value := reflect.ValueOf(item)
		ans = append(ans, value.FieldByName(key).Interface().(string))
	}
	return ans
}
