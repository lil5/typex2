export type TestBool = boolean

export type TestByte = number

export type TestBytes = number[]

export type TestFloat32 = number

export type TestFloat64 = number

export type TestInt = number

export type TestInt16 = number

export type TestInt32 = number

export type TestInt64 = number

export type TestInt8 = number

export type TestMap = Record<string, string>

export type TestMapInt = Record<number, string>

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

export type TestUint = number

export type TestUint16 = number

export type TestUint32 = number

export type TestUint64 = number

export type TestUint8 = number

interface testStruct4NotExported {
	exportMe: boolean
}

