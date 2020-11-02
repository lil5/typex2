package storage

type TestByte = byte
type TestBytes = []byte

type TestSliceString []string
type TestSliceSlice []TestSliceString

type TestMap map[string]string
type TestMapInt map[int]string

type TestPointer *string
type TestPointerStruct struct {
	PP *string
}

type TestRune rune
type TestUintptr uintptr
type TestInterface interface{}

type TestChan chan int
type TestFunc func()
