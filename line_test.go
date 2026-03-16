package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Tests of Line functions
type LineTestSuite struct {
	suite.Suite
}

func (suite *LineTestSuite) TestLineSplit() {
	testLine := newLine("Hello", 0)
	testSplitLine := testLine.split(4)
	wantRemain := newLine("Hell", 0)
	wantSplit := newLine("o", 0)
	suite.Equal(wantRemain.content, testLine.content, "What remains after split")
	suite.Equal(wantSplit.content, testSplitLine.content, "What was split")
}

func (suite *LineTestSuite) TestLineJoin() {
	testLine1 := newLine("Hello, ", 0)
	testLine2 := newLine("World!", 0)
	expectedLine := newLine("Hello, World!", 0)
	testLine1.join(testLine2)
	suite.Equal(expectedLine.content, testLine1.content, "Joining rows")
}

func (suite *LineTestSuite) TestInsertChar() {
	testLine := newLine("Howdy", 0)
	testLine.insertChar('R', 2)
	expectedLine := newLine("HoRwdy", 0)
	suite.Equal(expectedLine.content, testLine.content, "Inserting character")
}

func (suite *LineTestSuite) TestDeleteChar() {
	testLine := newLine("HoRwdy", 0)
	testLine.deleteChar(2)
	expectedLine := newLine("Howdy", 0)
	suite.Equal(expectedLine.content, testLine.content)
}

func (suite *LineTestSuite) TestNextWord() {
	testLine := newLine("Hello and goodbye", 0)
	suite.Equal(6, testLine.nextWord(2))
	suite.Equal(16, testLine.nextWord(13))
}

func (suite *LineTestSuite) TestPrevWord() {
	testLine := newLine("Hello and goodbye", 0)
	suite.Equal(0, testLine.prevWord(3))
	suite.Equal(0, testLine.prevWord(7))
	suite.Equal(6, testLine.prevWord(13))
}

func (suite *LineTestSuite) TestWordEnd() {
	testLine := newLine("Hello and goodbye", 0)
	suite.Equal(4, testLine.wordEnd(1))
	suite.Equal(16, testLine.wordEnd(14))
	suite.Equal(8, testLine.wordEnd(6))
}

func (suite *LineTestSuite) TestWordStart() {
	testLine := newLine("Hello and goodbye", 0)
	suite.Equal(0, testLine.wordStart(1))
	suite.Equal(10, testLine.wordStart(14))
	suite.Equal(6, testLine.wordStart(8))
}

// Run Line test suite
func TestLineTestSuite(t *testing.T) {
	suite.Run(t, new(LineTestSuite))
}
