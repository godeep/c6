package c6

import (
	"c6/ast"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunParserTest(code string) []ast.Statement {
	var parser = NewParser(NewContext())
	return parser.ParseScss(code)
}

func BenchmarkParserComplexSelector(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var parser = NewParser(NewContext())
		parser.ParseScss(`div#myId.first-name.last-name, span, html, .first-name, .last-name { }`)
	}
}

func BenchmarkParserOverAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var parser = NewParser(NewContext())
		parser.ParseScss(`div#myId.first-name.last-name {
			.foo-bar {
				color: red;
				background: #fff;
				border-radius: 10px;
			}

			@for $i from 1 through 100 { }
			@if $i == 1 {
			} @else if $i == 2 {
			} @else if $i == 3 {
			} @else {
			}
		}`)
	}

}

func TestParserEmptyRuleSetWithUniversalSelector(t *testing.T) {
	var stmts = RunParserTest(`* { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserEmptyRuleSetWithClassSelector(t *testing.T) {
	var stmts = RunParserTest(`.first-name { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserEmptyRuleSetWithIdSelector(t *testing.T) {
	var stmts = RunParserTest(`#myId { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserEmptyRuleSetWithTypeSelector(t *testing.T) {
	var stmts = RunParserTest(`div { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserEmptyRuleSetWithTypeSelectorGroup(t *testing.T) {
	var stmts = RunParserTest(`div, span, html { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserEmptyRuleSetWithComplexSelector(t *testing.T) {
	var stmts = RunParserTest(`div#myId.first-name.last-name, span, html, .first-name, .last-name { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserNestedRuleSetSimple(t *testing.T) {
	var stmts = RunParserTest(`div, span, html { .foo { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserNestedRuleSetSimple2(t *testing.T) {
	var stmts = RunParserTest(`div, span, html { .foo { color: red; background: blue; } text-align: text; float: left; }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserNestedRuleWithParentSelector(t *testing.T) {
	var stmts = RunParserTest(`div, span, html { & { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserPropertyNameBorderWidth(t *testing.T) {
	var stmts = RunParserTest(`div { border-width: 3px 3px 3px 3px; }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserPropertyNameBorderWidthInterpolation(t *testing.T) {
	var stmts = RunParserTest(`div { border-#{ $width }: 3px 3px 3px 3px; }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserPropertyNameBorderWidthInterpolation2(t *testing.T) {
	var stmts = RunParserTest(`div { #{ $name }: 3px 3px 3px 3px; }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserPropertyNameBorderWidthInterpolation3(t *testing.T) {
	var stmts = RunParserTest(`div { #{ $name }-left: 3px; }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserImportRuleWithUrl(t *testing.T) {
	parser := NewParser(NewContext())
	stmts := parser.ParseScss(`@import url("http://foo.com/bar.css");`)

	if len(stmts) == 0 {
		t.Fatal("Returned 0 statements")
	}

	rule, ok := stmts[0].(*ast.ImportStatement)
	assert.True(t, ok, "Convert to ImportStatement OK")
	assert.NotNil(t, rule)

	// it's not a relative url
	_, ok1 := rule.Url.(ast.RelativeUrl)
	assert.False(t, ok1)

	// it's a url
	url, ok2 := rule.Url.(ast.Url)
	assert.True(t, ok2)
	assert.Equal(t, "http://foo.com/bar.css", string(url))
}

func TestParserImportRuleWithString(t *testing.T) {
	parser := NewParser(NewContext())
	stmts := parser.ParseScss(`@import "foo.css";`)

	if len(stmts) == 0 {
		t.Fatal("Returned 0 statements")
	}

	rule, ok := stmts[0].(*ast.ImportStatement)
	assert.True(t, ok, "Convert to ImportStatement OK")
	assert.NotNil(t, rule)

	// it's not a relative url
	url, ok := rule.Url.(ast.RelativeUrl)
	assert.True(t, ok)

	assert.True(t, ok)
	assert.Equal(t, "foo.css", string(url))
}

func TestParserImportRuleWithMedia(t *testing.T) {
	var stmts = RunParserTest(`@import url("foo.css") screen;`)
	assert.Equal(t, 1, len(stmts))
	fmt.Printf("Statements: %+v\n", stmts)
}

func TestParserImportRuleWithMultipleMediaTypes(t *testing.T) {
	var stmts = RunParserTest(`@import url("bluish.css") projection, tv;`)
	assert.Equal(t, 1, len(stmts))
	fmt.Printf("Statements: %+v\n", stmts)
}

func TestParserImportRuleWithMediaTypeAndColorFeature(t *testing.T) {
	var stmts = RunParserTest(`@import url(color.css) screen and (color);`)
	assert.Equal(t, 1, len(stmts))
	fmt.Printf("Statements: %+v\n", stmts)
}

func TestParserImportRuleWithMediaTypeAndMaxWidthFeature(t *testing.T) {
	var stmts = RunParserTest(`@import url(color.css) screen and (max-width: 300px);`)
	assert.Equal(t, 1, len(stmts))
	fmt.Printf("Statements: %+v\n", stmts)
}

func TestParserImportRuleWithMedia2(t *testing.T) {
	var stmts = RunParserTest(`@import url("foo.css") screen and (orientation:landscape);`)
	assert.Equal(t, 1, len(stmts))
	fmt.Printf("Statements: %+v\n", stmts)
}

func TestParserMediaQuerySimple(t *testing.T) {
	var stmts = RunParserTest(`@media screen { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryNotScreen(t *testing.T) {
	var stmts = RunParserTest(`@media not screen { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryOnlyScreen(t *testing.T) {
	var stmts = RunParserTest(`@media only screen { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryAllAndMinWidth(t *testing.T) {
	var stmts = RunParserTest(`@media all and (min-width:500px) {  .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryMinWidth(t *testing.T) {
	var stmts = RunParserTest(`@media (min-width:500px) {  .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryOrientationPortrait(t *testing.T) {
	var stmts = RunParserTest(`@media (orientation: portrait) { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryMultipleWithComma(t *testing.T) {
	var stmts = RunParserTest(`@media screen and (color), projection and (color) { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryNone(t *testing.T) {
	var stmts = RunParserTest(`@media { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryNotAndMonoChrome(t *testing.T) {
	var stmts = RunParserTest(`@media not all and (monochrome) { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryJustAll(t *testing.T) {
	var stmts = RunParserTest(`@media all { .red { color: red; } }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserMediaQueryWithVendorPrefixFeature(t *testing.T) {
	// FIXME: 'min--moz-device-pixel-ratio' will become '-moz-device-pixel-ratio'
	var stmts = RunParserTest(`@media (-webkit-min-device-pixel-ratio: 2), (min--moz-device-pixel-ratio: 2) {  }`)
	assert.Equal(t, 1, len(stmts))
	t.Logf("Stmts: %+v", stmts)
}

func TestParserIfTrueStatement(t *testing.T) {
	var stmts = RunParserTest(`@if true {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfFalseElseStatement(t *testing.T) {
	var stmts = RunParserTest(`@if false {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfFalseOrTrueElseStatement(t *testing.T) {
	var stmts = RunParserTest(`@if false or true {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfTrueAndTrueOrFalseElseStatement(t *testing.T) {
	var stmts = RunParserTest(`@if true and true or true {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfTrueAndTrueOrFalseElseStatement2(t *testing.T) {
	var stmts = RunParserTest(`@if (true and true) or true {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonGreaterThan(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) > 2 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonGreaterEqual(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) >= 2 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonLessThan(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) < 2 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonLessEqual(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) <= 2 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonEqual(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) == 6 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonUnequal(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) != 6 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserIfComparisonUnequalElseIf(t *testing.T) {
	var stmts = RunParserTest(`@if (3+3) != 6 {  } @else if (3+3) == 6 {  } @else {  }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserForStatementSimple(t *testing.T) {
	var stmts = RunParserTest(`@for $var from 1 through 20 { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserForStatementExpressionReduce(t *testing.T) {
	var stmts = RunParserTest(`@for $var from 2 * 3 through 20 * 5 + 10 { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserForStatementRangeOperator(t *testing.T) {
	var stmts = RunParserTest(`@for $var in 1 .. 10 { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserForStatementRangeOperatorWithExpression(t *testing.T) {
	var stmts = RunParserTest(`@for $var in 2 + 3 .. 10 * 10 { }`)
	assert.Equal(t, 1, len(stmts))
}

func TestParserWhileStatement(t *testing.T) {
	code := `
$i: 6;
@while $i > 0 { $i: $i - 2; }
`
	var stmts = RunParserTest(code)
	assert.Equal(t, 1, len(stmts))
}

func TestParserCSS3Gradient(t *testing.T) {
	// some test cases from htmldog
	// @see http://www.htmldog.com/guides/css/advanced/gradients/
	var buffers = []string{
		`div { background: repeating-linear-gradient(white, black 10px, white 20px); }`,
		`div { background: linear-gradient(135deg, hsl(36,100%,50%) 10%, hsl(72,100%,50%) 60%, white 90%); }`,
		`div { background: linear-gradient(black 0, white 100%); }`,
		`div { background: radial-gradient(#06c 0, #fc0 50%, #039 100%); }`,
		`div { background: linear-gradient(red 0%, green 33.3%, blue 66.7%, black 100%); }`,
		`div { background: -webkit-radial-gradient(100px 200px, circle closest-side, black, white); }`,
	}
	for _, buffer := range buffers {
		var block = RunParserTest(buffer)
		fmt.Printf("%+v\n", block)
	}
}

func TestParserPropertyListExpression(t *testing.T) {
	var buffers []string = []string{
		`div { width: 1px; }`,
		`div { width: 2px 3px; }`,
		`div { width: 4px, 5px, 6px, 7px; }`,
		`div { width: 4px, 5px 6px, 7px; }`,
		`div { width: 10px 3px + 7px 20px; }`,
		// `div { width: 10px, 3px + 7px, 20px; }`,
	}
	for _, buffer := range buffers {
		var block = RunParserTest(buffer)
		fmt.Printf("%+v\n", block)
	}
}

func TestParserFontCssSlash(t *testing.T) {
	// should be plain CSS, no division
	// TODO: verify this case
	var block = RunParserTest(`.foo { font: 12px/24px; }`)
	fmt.Printf("%+v\n", block)
}

func TestParserVariableAssignmentWithMorePlus(t *testing.T) {
	var block = RunParserTest(`$foo: 12px + 20px + 20px;`)
	fmt.Printf("%+v\n", block)
}

func TestParserVariableAssignmentWithExpressionDefaultFlag(t *testing.T) {
	var block = RunParserTest(`$foo: 12px + 20px + 20px !default;`)
	fmt.Printf("%+v\n", block)
}

func TestParserVariableAssignmentWithExpressionOptionalFlag(t *testing.T) {
	var block = RunParserTest(`$foo: 12px + 20px + 20px !optional;`)
	fmt.Printf("%+v\n", block)
}

func TestParserVariableAssignmentWithComplexExpression(t *testing.T) {
	var stmts = RunParserTest(`$foo: 12px * (20px + 20px) + 4px / 2;`)
	fmt.Printf("%+v\n", stmts[0])
}

func TestParserVariableAssignmentWithInterpolation(t *testing.T) {
	var stmts = RunParserTest(`$foo: #{ 10 + 20 }px;`)
	fmt.Printf("%+v\n", stmts[0])
}

func TestParserVariableAssignmentLengthPlusLength(t *testing.T) {
	var stmts = RunParserTest(`$foo: 10px + 20px;`)
	fmt.Printf("%+v\n", stmts)
}

func TestParserVariableAssignmentNumberPlusNumberMulLength(t *testing.T) {
	var stmts = RunParserTest(`$foo: (10 + 20) * 3px;`)
	fmt.Printf("%+v\n", stmts)
}

func TestParserVariableAssignmentWithHexColorAddOperation(t *testing.T) {
	var stmts = RunParserTest(`$foo: #000 + 10;`)
	fmt.Printf("%+v\n", stmts)
}

func TestParserVariableAssignmentWithHexColorMulOperation(t *testing.T) {
	var stmts = RunParserTest(`$foo: #010101 * 20;`)
	fmt.Printf("%+v\n", stmts)
}

func TestParserVariableAssignmentWithHexColorDivOperation(t *testing.T) {
	var stmts = RunParserTest(`$foo: #121212 / 2;`)
	fmt.Printf("%+v\n", stmts)
}

func TestParserVariableAssignmentWithPxValue(t *testing.T) {
	var stmts = RunParserTest(`$foo: 10px;`)
	fmt.Printf("%+v\n", stmts)
}

func TestParserMassiveRules(t *testing.T) {
	var buffers []string = []string{
		`div { width: auto; }`,
		`div { width: 100px }`,
		`div { width: 100pt }`,
		`div { width: 100em }`,
		`div { width: 100rem }`,
		`div { padding: 10px 10px; }`,
		`div { padding: 10px 10px 20px 30px; }`,
		`div { padding: 10px + 10px; }`,
		`div { padding: 10px + 10px * 3; }`,
		`div { color: red; }`,
		`div { color: rgb(255,255,255); }`,
		`div { color: rgba(255,255,255,0); }`,
		`div { background-image: url("../images/foo.png"); }`,
		// `div { color: #ccddee; }`,
	}
	for _, buffer := range buffers {
		fmt.Printf("Input %s\n", buffer)
		var parser = NewParser(NewContext())
		var stmts = parser.ParseScss(buffer)
		fmt.Printf("%+v\n", stmts)
	}
}

func TestParserTypeSelectorRule(t *testing.T) {
	parser := NewParser(NewContext())
	stmts := parser.ParseScss(`div { width: auto; }`)

	ruleset, ok := stmts[0].(*ast.RuleSet)
	assert.True(t, ok)

	t.Logf("%+v\n", ruleset.Selectors)
	t.Logf("%+v\n", ruleset.Block)
}

/*
func TestParserIfStatementTrueCondition(t *testing.T) {
	parser := NewParser(NewContext())
	block := parser.ParseScss(`
	div {
		@if true {
			color: red;
		}
	}
	`)
	_ = block
}
*/
