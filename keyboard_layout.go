package elasic_keyboard_layout

// Suggester для плагина замены раскладки
// @see https://github.com/papahigh/elasticsearch-keyboard-layout
type KeyboardLayoutSuggester struct {
	name string
	
	text           string
	field          string
	language       string
	lowerCaseToken bool
	preserveCase   bool
	addOriginal    bool
	minFreq        string
	maxFreq        string
}

func (r *KeyboardLayoutSuggester) SetMaxFreq(maxFreq string) *KeyboardLayoutSuggester {
	r.maxFreq = maxFreq
	return r
}

func (r *KeyboardLayoutSuggester) SetMinFreq(minFreq string) *KeyboardLayoutSuggester {
	r.minFreq = minFreq
	return r
}

func (r *KeyboardLayoutSuggester) SetAddOriginal(addOriginal bool) *KeyboardLayoutSuggester {
	r.addOriginal = addOriginal
	return r
}

func (r *KeyboardLayoutSuggester) SetPreserveCase(preserveCase bool) *KeyboardLayoutSuggester {
	r.preserveCase = preserveCase
	return r
}

func (r *KeyboardLayoutSuggester) SetLowerCaseToken(lowerCaseToken bool) *KeyboardLayoutSuggester {
	r.lowerCaseToken = lowerCaseToken
	return r
}

func (r *KeyboardLayoutSuggester) SetLanguage(language string) *KeyboardLayoutSuggester {
	r.language = language
	return r
}

func (r *KeyboardLayoutSuggester) SetField(field string) *KeyboardLayoutSuggester {
	r.field = field
	return r
}

func (r *KeyboardLayoutSuggester) SetText(text string) *KeyboardLayoutSuggester {
	r.text = text
	return r
}

// Создает новый экзепляр KeyboardLayout
func NewKeyboardLayoutSuggester(name string) *KeyboardLayoutSuggester {
	return &KeyboardLayoutSuggester{
		name: name,
	}
}

// Получает имя
func (r *KeyboardLayoutSuggester) Name() string {
	return r.name
}

// Исходный объект запроса
type keyboardLayoutSuggesterRequest struct {
	Text           string      `json:"text"`
	KeyboardLayout interface{} `json:"keyboard_layout"`
}

// Получает объект запроса
func (r *KeyboardLayoutSuggester) Source(includeName bool) (interface{}, error) {
	keyboardLayout := &keyboardLayoutSuggesterRequest{}
	
	if r.text != "" {
		keyboardLayout.Text = r.text
	}
	
	suggester := make(map[string]interface{})
	keyboardLayout.KeyboardLayout = suggester
	
	suggester["field"] = r.field
	suggester["language"] = r.language
	
	if r.lowerCaseToken {
		suggester["lowercase_token"] = r.lowerCaseToken
	}
	if r.preserveCase {
		suggester["preserve_case"] = r.preserveCase
	}
	if r.addOriginal {
		suggester["add_original"] = r.addOriginal
	}
	if r.minFreq != "" {
		suggester["min_freq"] = r.minFreq
	}
	if r.maxFreq != "" {
		suggester["max_freq"] = r.maxFreq
	}
	
	if !includeName {
		return keyboardLayout, nil
	}
	
	source := make(map[string]interface{})
	source[r.name] = keyboardLayout
	return source, nil
}
