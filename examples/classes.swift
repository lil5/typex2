typealias TestBool = Bool

typealias TestByte = Int

typealias TestBytes = Array<Int>

typealias TestChan = Any

typealias TestFloat32 = Float

typealias TestFloat64 = Double

typealias TestFunc = Any

typealias TestInt = Int

typealias TestInt16 = Int

typealias TestInt32 = Int

typealias TestInt64 = Int

typealias TestInt8 = Int

typealias TestInterface = Dictionary<String, Any>

typealias TestMap = Dictionary<String, String>

typealias TestMapInt = Dictionary<Int, String>

typealias TestPointer = String?

class TestPointerStruct {
	PP String?;
}

typealias TestRune = Int

typealias TestSliceSlice = Array<TestSliceString>

typealias TestSliceString = Array<String>

typealias TestString = String

class TestStruct1 {
	One Dictionary<String, String>;
}

class TestStruct2Named {
	Two String;
}

class TestStruct3Extend: TestStruct1 {
	TestStruct2Named TestStruct2Named;
}

class TestStruct5UsePackage {
	ValNotExp testStruct4NotExported;
}

typealias TestUint = Int

typealias TestUint16 = Int

typealias TestUint32 = Int

typealias TestUint64 = Int

typealias TestUint8 = Int

typealias TestUintptr = Any

class testStruct4NotExported {
	ExportMe Bool;
}

