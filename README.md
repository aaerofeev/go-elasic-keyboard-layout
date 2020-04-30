# Extension for github.com/olivere/elastic add support for keyboard-layout plugin

Usage:
```
suggester := go-elastic-keyboard-layout.NewKeyboardLayoutSuggester(keyboardSuggesterName)
suggester.SetAddOriginal(true).
	SetField("query.shingle").
	SetLanguage("russian").
	SetLowerCaseToken(true).
	SetPreserveCase(true).
	SetText(phrase)
```

https://github.com/papahigh/elasticsearch-keyboard-layout
