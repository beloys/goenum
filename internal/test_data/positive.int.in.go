package test_data

//go:generate goenum -p
type TestIntEnum int

const (
	TestIntEnumZero TestIntEnum = 0
	TestIntEnumOne  TestIntEnum = 1
)
