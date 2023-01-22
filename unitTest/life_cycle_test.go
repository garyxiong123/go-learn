package unitTest

import "testing"

//Go testing package also has a way to run function before and after
//each test function, Setup and Teardown function respectively and it can
//be used to do tasks like creating a temporary database, opening a file,
//starting a server and clean up the resources after each test.

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	println(code)
	tearDown()
}

func setUp() {
	println("gary ---- setUp-----")
}

func tearDown() {
	println("gary ---- tearDown-----")
}
