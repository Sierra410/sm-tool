package main

import (
	"regexp"

	"github.com/hjson/hjson-go"
)

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
	uuid         string
	kind         string
	descriptions map[string]*smPartDescription
	partData     string
	partDataJson map[string]interface{}
}

func smPartNew(uuid string) *smPart {
	if !validateUuid(uuid) {
		uuid = randomUuid()
	}

	return &smPart{
		uuid:         uuid,
		descriptions: map[string]*smPartDescription{},
		partDataJson: map[string]interface{}{},
	}
}

func (self *smPart) unmarshalPartData() error {
	json := map[string]interface{}{}

	err := hjson.Unmarshal([]byte(self.partData), &json)
	if err != nil {
		return err
	}

	json["uuid"] = self.uuid
	self.partDataJson = json

	return nil
}

func (self *smPart) setTitle(title, lang string) {
	pd, ok := self.descriptions[lang]
	if ok {
		pd.Title = title
	} else {
		self.descriptions[lang] = &smPartDescription{
			Title: title,
		}
	}
}

func (self *smPart) setDescription(desc, lang string) {
	pd, ok := self.descriptions[lang]
	if ok {
		pd.Description = desc
	} else {
		self.descriptions[lang] = &smPartDescription{
			Description: desc,
		}
	}
}

func (self *smPart) getTitle(lang string) string {
	pd, ok := self.descriptions[lang]
	if ok {
		return pd.Title
	}

	return ""
}

func (self *smPart) getDescription(lang string) string {
	pd, ok := self.descriptions[lang]
	if ok {
		return pd.Description
	}

	return ""
}

func (self *smPart) matches(r *regexp.Regexp) bool {
	if r.MatchString(self.uuid) {
		return true
	}

	for _, desc := range self.descriptions {
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
