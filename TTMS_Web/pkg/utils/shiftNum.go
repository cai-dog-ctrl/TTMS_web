package utils

import (
	"strconv"
)

func ShiftToNum64(num string)int64{
	i,_:=strconv.ParseInt(num,10,64)
	return i
}
func ShiftToNum(num string)int{
	i, _ := strconv.Atoi(num)
	return i
}
func ShiftToStringFromInt64(num int64)string{
	return strconv.FormatInt(num,10)
}
func ShiftToStringFromInt32(num int)string{
  return strconv.Itoa(num)
}

