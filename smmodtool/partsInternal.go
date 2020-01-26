package main

import "regexp"

const (
	languageBrazilian = "Brazilian"
	languageChinese   = "Chinese"
	languageEnglish   = "English"
	languageFrench    = "French"
	languageGerman    = "German"
	languageItalian   = "Italian"
	languageJapanese  = "Japanese"
	languageKorean    = "Korean"
	languagePolish    = "Polish"
	languageRussian   = "Russian"
	languageSpanish   = "Spanish"
)

type smPartDescription struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type smPart struct {
	Uuid         string
	Descriptions map[string]*smPartDescription
	PartData     string
}

func (self *smPart) setTitle(title, lang string) {
	pd, ok := self.Descriptions[lang]
	if ok {
		pd.Title = title
	} else {
		self.Descriptions[lang] = &smPartDescription{
			Title: title,
		}
	}
}

func (self *smPart) setDescription(desc, lang string) {
	pd, ok := self.Descriptions[lang]
	if ok {
		pd.Description = desc
	} else {
		self.Descriptions[lang] = &smPartDescription{
			Description: desc,
		}
	}
}

func (self *smPart) getTitle(lang string) string {
	pd, ok := self.Descriptions[lang]
	if ok {
		return pd.Title
	}

	return ""
}

func (self *smPart) getDescription(lang string) string {
	pd, ok := self.Descriptions[lang]
	if ok {
		return pd.Description
	}

	return ""
}

func (self *smPart) matches(r *regexp.Regexp) bool {
	if r.MatchString(self.Uuid) {
		return true
	}

	for _, desc := range self.Descriptions {
		if r.MatchString(desc.Title) {
			return true
		}

		if r.MatchString(desc.Description) {
			return true
		}

		for _, keyword := range desc.Keywords {
			if r.MatchString(keyword) {
				return true
			}
		}
	}

	return false
}
