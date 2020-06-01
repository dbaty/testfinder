package main

import "reflect"
import "testing"


func TestParsePythonFile(t *testing.T) {
	functions := parsePythonFile("testdata/test.py")
	expected := []string{
		"testdata/test.py::TestEmptyClass",
		"testdata/test.py::TestClassWithMethods",
		"testdata/test.py::TestClassWithMethods::test_method1",
		"testdata/test.py::TestClassWithMethods::test_method2",
		"testdata/test.py::test_func",
	}

	if !reflect.DeepEqual(functions, expected) {
		t.Errorf("got %s, expected %s.", functions, expected)
	}
}
