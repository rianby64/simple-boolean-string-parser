package parser

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func Test_parser_one_not(t *testing.T) {
	NotStrCalled := false
	NotStr := func(a string) squirrel.Sqlizer {
		assert.Equal(t, "alice", a)

		NotStrCalled = true

		r := squirrel.Expr("col <> '%s'", a)
		return r
	}

	p := Parser{
		NotStr: NotStr,
	}

	p.Go("not alice")
	assert.True(t, NotStrCalled)
}

func Test_parser_two_ands(t *testing.T) {
	StrANDStrCalled := false
	StrANDStr := func(a, b string) squirrel.And {
		assert.Equal(t, "alice", a)
		assert.Equal(t, "bob", b)

		StrANDStrCalled = true

		r := squirrel.And{
			squirrel.Expr("col = '%s'", a),
			squirrel.Expr("col = '%s'", b),
		}
		return r
	}

	p := Parser{
		StrANDStr: StrANDStr,
	}

	p.Go("alice and bob")
	assert.True(t, StrANDStrCalled)
}

func Test_parser_three_ands(t *testing.T) {
	StrANDStrCalled := false
	ExpANDStrCalled := false

	StrANDStr := func(a, b string) squirrel.And {
		assert.Equal(t, "alice", a)
		assert.Equal(t, "bob", b)

		StrANDStrCalled = true

		return squirrel.And{
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
		StrANDStr: StrANDStr,
		ExpANDStr: ExpANDStr,
	}

	p.Go("alice and bob and carol")
	assert.True(t, StrANDStrCalled)
	assert.True(t, ExpANDStrCalled)
}

func Test_parser_two_ors(t *testing.T) {
	StrORStrCalled := false
	StrORStr := func(a, b string) squirrel.Or {
		assert.Equal(t, "alice", a)
		assert.Equal(t, "bob", b)

		StrORStrCalled = true

		r := squirrel.Or{
			squirrel.Expr("col = '%s'", a),
			squirrel.Expr("col = '%s'", b),
		}
		return r
	}

	p := Parser{
		StrORStr: StrORStr,
	}

	p.Go("alice or bob")
	assert.True(t, StrORStrCalled)
}

func Test_parser_three_ors(t *testing.T) {
	StrORStrCalled := false
	ExpORStrCalled := false

	StrORStr := func(a, b string) squirrel.Or {
		assert.Equal(t, "alice", a)
		assert.Equal(t, "bob", b)

		StrORStrCalled = true

		return squirrel.Or{
			squirrel.Expr("col = '%s'", a),
			squirrel.Expr("col = '%s'", b),
		}
	}

	ExpORStr := func(a squirrel.Sqlizer, b string) squirrel.Or {
		assert.Equal(t, "carol", b)

		ExpORStrCalled = true

		return squirrel.Or{
			a,
			squirrel.Expr("col = '%s'", b),
		}
	}

	p := Parser{
		StrORStr: StrORStr,
		ExpORStr: ExpORStr,
	}

	p.Go("alice or bob or carol")
	assert.True(t, StrORStrCalled)
	assert.True(t, ExpORStrCalled)
}

func Test_parser_case1(t *testing.T) {
	StrANDStrCalled := false
	ExpANDStrCalled := false
	ExpORStrCalled := false

	StrANDStr := func(a, b string) squirrel.And {
		assert.Equal(t, "alice", a)
		assert.Equal(t, "bob", b)

		StrANDStrCalled = true

		return squirrel.And{
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

	ExpORStr := func(a squirrel.Sqlizer, b string) squirrel.Or {
		assert.Equal(t, "dan", b)

		ExpORStrCalled = true

		return squirrel.Or{
			a,
			squirrel.Expr("col = '%s'", b),
		}
	}

	p := Parser{
		StrANDStr: StrANDStr,
		ExpANDStr: ExpANDStr,
		ExpORStr:  ExpORStr,
	}

	p.Go("alice and bob and carol or dan")
	assert.True(t, StrANDStrCalled)
	assert.True(t, ExpANDStrCalled)
	assert.True(t, ExpORStrCalled)
}
