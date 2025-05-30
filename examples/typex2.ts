export type TestBool = boolean

export type TestByte = number

export type TestBytes = number[]

export type TestChan = unknown

export type TestFloat32 = number

export type TestFloat64 = number

export type TestFunc = unknown

export type TestInt = number

export type TestInt16 = number

export type TestInt32 = number

export type TestInt64 = number

export type TestInt8 = number

export type TestInterface = Record<string, any>

export type TestMap = Record<string, string>

export type TestMapInt = Record<number, string>

export type TestPointer = (string | null | undefined)

export interface TestPointerStruct {
	PP?: (string | null | undefined)
}

export type TestRune = number

export type TestSliceSlice = TestSliceString[]

export type TestSliceString = string[]

export type TestString = string

export interface TestStruct1 {
	One: Record<string, string>
}

export interface TestStruct2Named {
	id: string
	two: string
}

export interface TestStruct3Extend extends TestStruct1, TestStruct2Named {
}

export interface TestStruct5UsePackage {
	myInt: TestInt
	notExp: testStruct4NotExported
}

export interface TestStruct6NamedOmitempty {
	one?: (string | null | undefined)
}

export interface TestStruct7NamedOmit {
}

export interface TestStruct8Time {
	createdAt: string
}

export interface TestStruct9JsonFormUri {
	One: number
	two: number
	three: number
	four: number
}

export type TestUint = number

export type TestUint16 = number

export type TestUint32 = number

export type TestUint64 = number

export type TestUint8 = number

export type TestUintptr = any

interface testStruct4NotExported {
	exportMe: boolean
}

