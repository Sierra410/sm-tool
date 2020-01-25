package main

import "regexp"

const (
	languageBrazilian = iota
	languageChinese
	languageEnglish
	languageFrench
	languageGerman
	languageItalian
	languageJapanese
	languageKorean
	languagePolish
	languageRussian
	languageSpanish
)

var lanugages = []string{
	"Brazilian",
	"Chinese",
	"English",
	"French",
	"German",
	"Italian",
	"Japanese",
	"Korean",
	"Polish",
	"Russian",
	"Spanish",
}

type partDescription struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type smPart struct {
	Uuid         string
	Descriptions map[int]partDescription
	PartData     map[string]interface{}
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

func newSmPart() {

}
