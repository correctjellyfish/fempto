package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Tests of Row functions
type RowTestSuite struct {
	suite.Suite
}

func (suite *RowTestSuite) TestRowSplit() {
	testRow := newRow("Hello", 0)
	testSplitRow := testRow.split(4)
	wantRemain := newRow("Hell", 0)
	wantSplit := newRow("o", 0)
	suite.Equal(wantRemain.content, testRow.content, "What remains after split")
	suite.Equal(wantSplit.content, testSplitRow.content, "What was split")
}

func (suite *RowTestSuite) TestRowJoin() {
	testRow1 := newRow("Hello, ", 0)
	testRow2 := newRow("World!", 0)
	expectedRow := newRow("Hello, World!", 0)
	testRow1.join(testRow2)
	suite.Equal(expectedRow.content, testRow1.content, "Joining rows")
}

func (suite *RowTestSuite) TestInsertChar() {
	testRow := newRow("Howdy", 0)
	testRow.insertChar('R', 2)
	expectedRow := newRow("HoRwdy", 0)
	suite.Equal(expectedRow.content, testRow.content, "Inserting character")
}

func (suite *RowTestSuite) TestDeleteChar() {
	testRow := newRow("HoRwdy", 0)
	testRow.deleteChar(2)
	expectedRow := newRow("Howdy", 0)
	suite.Equal(expectedRow.content, testRow.content)
}

func (suite *RowTestSuite) TestNextWord() {
	testRow := newRow("Hello and goodbye", 0)
	suite.Equal(6, testRow.nextWord(2))
	suite.Equal(16, testRow.nextWord(13))
}

func (suite *RowTestSuite) TestPrevWord() {
	testRow := newRow("Hello and goodbye", 0)
	suite.Equal(0, testRow.prevWord(3))
	suite.Equal(0, testRow.prevWord(7))
	suite.Equal(6, testRow.prevWord(13))
}

func (suite *RowTestSuite) TestWordEnd() {
	testRow := newRow("Hello and goodbye", 0)
	suite.Equal(4, testRow.wordEnd(1))
	suite.Equal(16, testRow.wordEnd(14))
	suite.Equal(8, testRow.wordEnd(6))
}

func (suite *RowTestSuite) TestWordStart() {
	testRow := newRow("Hello and goodbye", 0)
	suite.Equal(0, testRow.wordStart(1))
	suite.Equal(10, testRow.wordStart(14))
	suite.Equal(6, testRow.wordStart(8))
}

// Run Row test suite
func TestRowTestSuite(t *testing.T) {
	suite.Run(t, new(RowTestSuite))
}
