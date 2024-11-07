typedef TestBool = bool

typedef TestByte = int

typedef TestBytes = List<int>

typedef TestChan = dynamic

typedef TestFloat32 = double

typedef TestFloat64 = double

typedef TestFunc = dynamic

typedef TestInt = int

typedef TestInt16 = int

typedef TestInt32 = int

typedef TestInt64 = int

typedef TestInt8 = int

typedef TestInterface = Map<String, dynamic>

typedef TestMap = Map<String, String>

typedef TestMapInt = Map<int, String>

typedef TestPointer = String?

class TestPointerStruct {
	PP String?;
}

typedef TestRune = int

typedef TestSliceSlice = List<TestSliceString>

typedef TestSliceString = List<String>

typedef TestString = String

class TestStruct1 {
	One Map<String, String>;
}

class TestStruct2Named {
	two String;
}

class TestStruct3Extend extends TestStruct1 {
	TestStruct2Named TestStruct2Named;
}

class TestStruct5UsePackage {
	notExp testStruct4NotExported;
}

class TestStruct6NamedOmitempty {
	one String?;
}

class TestStruct7NamedOmit {
}

typedef TestUint = int

typedef TestUint16 = int

typedef TestUint32 = int

typedef TestUint64 = int

typedef TestUint8 = int

typedef TestUintptr = dynamic

class testStruct4NotExported {
	exportMe bool;
}

