package parser

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

/*
  Cases tested:
	p.splitParentheses("(alice or bob) and carol")
	p.splitParentheses("alice or (bob and carol)")
	p.splitParentheses("(alice or bob) and (carol or dan)")
	p.splitParentheses("(alice or bob) and (carol or dan) or (elen and (frank or glenn))")
	p.splitParentheses("(alice or bob) and carol or dan or (elen and (frank or glenn))")
	p.splitParentheses("(alice or bob) and not (carol or dan) or (elen and not (frank or glenn))")
*/
func Test_splitParentheses_case1(t *testing.T) {
	p := Parser{}
	terms, err := p.splitParentheses("(alice or bob) and carol")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"(alice or bob)",
		" and carol",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitParentheses_case2(t *testing.T) {
	p := Parser{}
	terms, err := p.splitParentheses("alice or (bob and carol)")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"alice or ",
		"(bob and carol)",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitParentheses_case3(t *testing.T) {
	p := Parser{}
	terms, err := p.splitParentheses("(alice or bob) and (carol or dan)")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"(alice or bob)",
		" and ",
		"(carol or dan)",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitParentheses_case4(t *testing.T) {
	p := Parser{}
	terms, err := p.splitParentheses("(alice or bob) and (carol or dan) or (elen and (frank or glenn))")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"(alice or bob)",
		" and ",
		"(carol or dan)",
		" or ",
		"(elen and (frank or glenn))",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitParentheses_case5(t *testing.T) {
	p := Parser{}
	terms, err := p.splitParentheses("(alice or bob) and carol or dan or (elen and (frank or glenn))")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"(alice or bob)",
		" and carol or dan or ",
		"(elen and (frank or glenn))",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitParentheses_case6(t *testing.T) {
	p := Parser{}
	terms, err := p.splitParentheses("(alice or bob) and not (carol or dan) or (elen and not (frank or glenn))")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"(alice or bob)",
		" and not ",
		"(carol or dan)",
		" or ",
		"(elen and not (frank or glenn))",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitOr_case1(t *testing.T) {
	p := Parser{}
	terms, err := p.splitOr("(alice and bob) or carol")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"(alice and bob)",
		"carol",
	}
	assert.Equal(t, expected, terms)
}

func Test_splitOr_case2(t *testing.T) {
	p := Parser{}
	terms, err := p.splitOr("zero or (alice and bob) or carol or (dan and elen) or (frank and glenn)")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []string{
		"zero",
		"(alice and bob)",
		"carol",
		"(dan and elen)",
		"(frank and glenn)",
	}
	assert.Equal(t, expected, terms)
}

/*
Cases tested:
	p.Go("(alice or bob) and carol")         // (a | b) & c
	p.Go("alice or (bob and carol)")         // a | (b & c)
*/

/*
func Test_parenthesis_parser_case1(t *testing.T) {
	StrORStrCalled := false
	ExpANDStrCalled := false

	StrORStr := func(a, b string) squirrel.Or {
		assert.Equal(t, "alice", a)
		assert.Equal(t, "bob", b)

		StrORStrCalled = true

		return squirrel.Or{
			squirrel.Expr("col = '%s'", a),
			squirrel.Expr("col = '%s'", b),
		}
	}

	ExpANDStr := func(a squirrel.Sqlizer, b string) squirrel.And {
		assert.Equal(t, "carol", b)

		ExpANDStrCalled = true

		return squirrel.And{
			a,
			squirrel.Expr("col = '%s'", b),
		}
	}

	p := Parser{
		StrORStr:  StrORStr,
		ExpANDStr: ExpANDStr,
	}

	p.Go("(alice or bob) and carol")
	assert.True(t, StrORStrCalled)
	assert.True(t, ExpANDStrCalled)
}
*/

func Test_parenthesis_parser_case2(t *testing.T) {
	StrANDStrCalled := false
	StrORExpCalled := false

	StrANDStr := func(a, b string) squirrel.And {
		assert.Equal(t, "bob", a)
		assert.Equal(t, "carol", b)

		StrANDStrCalled = true

		return squirrel.And{
			squirrel.Expr("col = '%s'", a),
			squirrel.Expr("col = '%s'", b),
		}
	}

	StrORExp := func(a string, b squirrel.Sqlizer) squirrel.Or {
		assert.Equal(t, "alice", a)

		StrORExpCalled = true

		return squirrel.Or{
			squirrel.Expr("col = '%s'", a),
			b,
		}
	}

	p := Parser{
		StrANDStr: StrANDStr,
		StrORExp:  StrORExp,
	}

	_, err := p.Go("alice or (bob and carol)")
	assert.Nil(t, err)
	assert.True(t, StrANDStrCalled)
	assert.True(t, StrORExpCalled)
}
