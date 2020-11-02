package storage

type TestStruct1 struct {
	One map[string]string
}

type TestStruct2Named struct {
	ID  string `json:"id"`
	Two string `json:"two"`
}

type TestStruct3Extend struct {
	TestStruct1
	TestStruct2Named
}

type testStruct4NotExported struct {
	ExportMe bool `json:"exportMe"`
}

type TestStruct5UsePackage struct {
	MyInt     TestInt                `json:"myInt"`
	ValNotExp testStruct4NotExported `json:"notExp"`
}
