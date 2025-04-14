use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub type TestBool = bool;

#[derive(Serialize, Deserialize)]
pub type TestByte = u8;

#[derive(Serialize, Deserialize)]
pub type TestBytes = Vec<u8>;

#[derive(Serialize, Deserialize)]
pub type TestChan<T0> = T0;

#[derive(Serialize, Deserialize)]
pub type TestFloat32 = f32;

#[derive(Serialize, Deserialize)]
pub type TestFloat64 = f64;

#[derive(Serialize, Deserialize)]
pub type TestFunc<T0> = T0;

#[derive(Serialize, Deserialize)]
pub type TestInt = usize;

#[derive(Serialize, Deserialize)]
pub type TestInt16 = i16;

#[derive(Serialize, Deserialize)]
pub type TestInt32 = i32;

#[derive(Serialize, Deserialize)]
pub type TestInt64 = i64;

#[derive(Serialize, Deserialize)]
pub type TestInt8 = i8;

#[derive(Serialize, Deserialize)]
pub type TestInterface<T0> = Map<String, T0>;

#[derive(Serialize, Deserialize)]
pub type TestMap = Map<String, String>;

#[derive(Serialize, Deserialize)]
pub type TestMapInt = Map<usize, String>;

#[derive(Serialize, Deserialize)]
pub type TestPointer = Option<String>;

#[derive(Serialize, Deserialize)]
pub struct TestPointerStruct {
	PP: Option<String>,
}

#[derive(Serialize, Deserialize)]
pub type TestRune = i32;

#[derive(Serialize, Deserialize)]
pub type TestSliceSlice = Vec<TestSliceString>;

#[derive(Serialize, Deserialize)]
pub type TestSliceString = Vec<String>;

#[derive(Serialize, Deserialize)]
pub type TestString = String;

#[derive(Serialize, Deserialize)]
pub struct TestStruct1 {
	One: Map<String, String>,
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct2Named {
	id: String,
	two: String,
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct3Extend {
	TestStruct1: TestStruct1,
	TestStruct2Named: TestStruct2Named,
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct5UsePackage {
	myInt: TestInt,
	notExp: testStruct4NotExported,
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct6NamedOmitempty {
	one: Option<String>,
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct7NamedOmit {
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct8Time {
	createdAt: String,
}

#[derive(Serialize, Deserialize)]
pub struct TestStruct9JsonFormUri {
	One: usize,
	two: usize,
	three: usize,
	four: usize,
}

#[derive(Serialize, Deserialize)]
pub type TestUint = usize;

#[derive(Serialize, Deserialize)]
pub type TestUint16 = u16;

#[derive(Serialize, Deserialize)]
pub type TestUint32 = u32;

#[derive(Serialize, Deserialize)]
pub type TestUint64 = u64;

#[derive(Serialize, Deserialize)]
pub type TestUint8 = u8;

#[derive(Serialize, Deserialize)]
pub type TestUintptr<T0> = T0;

#[derive(Serialize, Deserialize)]
pub mod struct testStruct4NotExported {
	exportMe: bool,
}

