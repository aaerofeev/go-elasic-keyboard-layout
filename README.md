# Extension for elastic add support for keyboard-layout plugin

## Links:
- https://github.com/olivere/elastic/tree/release-branch.v7
- https://github.com/papahigh/elasticsearch-keyboard-layout

## Usage:
```
suggester := go-elastic-keyboard-layout.NewKeyboardLayoutSuggester(keyboardSuggesterName)
suggester.SetAddOriginal(true).
	SetField("query.shingle").
	SetLanguage("russian").
	SetLowerCaseToken(true).
	SetPreserveCase(true).
	SetText(phrase)
```
