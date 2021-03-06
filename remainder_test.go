package gommand

import "testing"

// TestRemainder is used to test the remainder argument parser.
func TestRemainder(t *testing.T) {
	r := NewRouter(&RouterConfig{
		PrefixCheck: StaticPrefix("%"),
	})
	r.SetCommand(&Command{
		Name: "remainder",
		ArgTransformers: []ArgTransformer{
			{
				Function:  StringTransformer,
				Remainder: true,
			},
		},
		Function: func(ctx *Context) error {
			last := ctx.Args[0].(string)
			if last != "\"hello\"" {
				t.Log("String is", last)
				t.FailNow()
			}
			return nil
		},
	})
	r.AddErrorHandler(func(_ *Context, err error) bool {
		t.Log(err)
		t.FailNow()
		return true
	})
	r.CommandProcessor(nil, 0, mockMessage("%remainder   \"hello\""), true)
}
