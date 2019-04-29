package gotree

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

// Define Element types to be used as Tree Node values.

// Int type represents an integer.
type Int int

func (i Int) Equals(e Element) bool {
	v, ok := e.(Int)
	if !ok {
		return false
	}
	return i == v
}

func (i Int) Less(e Element) bool {
	v, ok := e.(Int)
	if !ok {
		log.Printf("%+v is not an Int type.\n", e)
		return false
	}
	return i < v
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func TestIntegerElement(t *testing.T) {
	testCases := []struct {
		thisElem       Int
		equalElem      Int
		notEqualElem   Int
		lessElem       Int
		notLessElem    Int
		expectedString string
	}{
		{
			thisElem:       Int(-11),
			equalElem:      Int(-11),
			notEqualElem:   Int(-12),
			lessElem:       Int(-100),
			notLessElem:    Int(10),
			expectedString: "-11",
		},
		{
			thisElem:       Int(5),
			equalElem:      Int(5),
			notEqualElem:   Int(-5),
			lessElem:       Int(1),
			notLessElem:    Int(6),
			expectedString: "5",
		},
		{
			thisElem:       Int(0),
			equalElem:      Int(0),
			notEqualElem:   Int(1),
			lessElem:       Int(-1),
			notLessElem:    Int(100),
			expectedString: "0",
		},
	}
	t.Run("Integer Element", func(t *testing.T) {
		for _, tc := range testCases {
			if !tc.equalElem.Equals(tc.thisElem) {
				t.Errorf("Expected equality between %+v and %+v", tc.equalElem, tc.thisElem)
			}
			if tc.notEqualElem.Equals(tc.thisElem) {
				t.Errorf("Expected inequality between %+v and %+v", tc.notEqualElem, tc.thisElem)
			}
			if !tc.lessElem.Less(tc.thisElem) {
				t.Errorf("Expected %+v to be smaller than %+v", tc.lessElem, tc.thisElem)
			}
			if tc.notLessElem.Less(tc.thisElem) {
				t.Errorf("Expected %+v to be greater than %+v", tc.notLessElem, tc.thisElem)
			}
			if tc.thisElem.String() != tc.expectedString {
				t.Errorf("Expected %+v's string representation to be %s", tc.thisElem, tc.expectedString)
			}
		}
	})
}

// Person is a struct that contains a string and an int.
type Person struct {
	name string
	age  int
}

func (p Person) Equals(e Element) bool {
	v, ok := e.(Person)
	if !ok {
		return false
	}
	return p.name == v.name && p.age == v.age
}

func (p Person) Less(e Element) bool {
	v, ok := e.(Person)
	if !ok {
		log.Printf("%+v is not an Int type.\n", e)
		return false
	}
	if p.name == v.name {
		return p.age < v.age
	}
	return p.name < v.name
}

func (p Person) String() string {
	return fmt.Sprintf("<%s:%d>", p.name, p.age)
}

func TestStructElement(t *testing.T) {
	testCases := []struct {
		thisElem       Person
		notEqualElem   Person
		equalElem      Person
		lessElem       Person
		notLessElem    Person
		expectedString string
	}{
		{
			thisElem: Person{
				name: "John Doe",
				age:  31,
			},
			equalElem: Person{
				name: "John Doe",
				age:  31,
			},
			notEqualElem: Person{
				name: "John Doe",
				age:  30,
			},
			lessElem: Person{
				name: "John Doe",
				age:  10,
			},
			notLessElem: Person{
				name: "John Doe",
				age:  51,
			},
			expectedString: "<John Doe:31>",
		},
		{
			thisElem: Person{
				name: "John Doe",
				age:  21,
			},
			equalElem: Person{
				name: "John Doe",
				age:  21,
			},
			notEqualElem: Person{
				name: "John D",
				age:  21,
			},
			lessElem: Person{
				name: "Jack Doe",
				age:  21,
			},
			notLessElem: Person{
				name: "Julie Doe",
				age:  21,
			},
			expectedString: "<John Doe:21>",
		},
	}

	t.Run("Struct Element", func(t *testing.T) {
		for _, tc := range testCases {
			if !tc.equalElem.Equals(tc.thisElem) {
				t.Errorf("Expected equality between %+v and %+v", tc.equalElem, tc.thisElem)
			}
			if tc.notEqualElem.Equals(tc.thisElem) {
				t.Errorf("Expected inequality between %+v and %+v", tc.notEqualElem, tc.thisElem)
			}
			if !tc.lessElem.Less(tc.thisElem) {
				t.Errorf("Expected %+v to be smaller than %+v", tc.lessElem, tc.thisElem)
			}
			if tc.notLessElem.Less(tc.thisElem) {
				t.Errorf("Expected %+v to be greater than %+v", tc.notLessElem, tc.thisElem)
			}
			if tc.thisElem.String() != tc.expectedString {
				t.Errorf("Expected %+v's string representation to be %s", tc.thisElem, tc.expectedString)
			}
		}
	})
}
