package gocli

import "github.com/stretchr/testify/mock"

// NewParserMock is the function that returns a new ParserMock.
func NewParserMock() (r *ParserMock) {
	r = &ParserMock{}
	return
}

// ParserMock is the mock that implements the Parser interface.
type ParserMock struct {
	mock.Mock
}

// Parse is the method that parses the input.
func (m *ParserMock) Parse(args string) (i Input, err error) {
	// args
	a := m.Called(args)

	// return
	i = a.Get(0).(Input)
	err = a.Error(1)
	return
}