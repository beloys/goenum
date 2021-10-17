package test_data

import "strconv"

var priorityTestIntEnum = []TestIntEnum{
	TestIntEnumZero,
	TestIntEnumOne,
}

func (e TestIntEnum) String() string {
	return strconv.Itoa(int(e))
}
func (e TestIntEnum) In(values ...TestIntEnum) bool {
	for _, v := range values {
		if v == e {
			return true
		}
	}
	return false
}

func (e TestIntEnum) Priority() int {
	for k, v := range priorityTestIntEnum {
		if v == e {
			return k
		}
	}
	return -1
}
