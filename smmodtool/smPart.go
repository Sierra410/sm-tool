package main

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/hjson/hjson-go"
)

const (
	kindBlockList = "blockList"
	kindPartList  = "partList"
)

func getPartKindString(isBlock bool) string {
	if isBlock {
		return kindBlockList
	}

	return kindPartList
}

func getPartKind(s string) bool {
	return s == kindBlockList
}

type smPartDescription struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type smPart struct {
	uuid           string
	kind           bool
	descriptions   map[string]*smPartDescription
	partData       map[string]interface{}
	partDataJson   string
	unmarshalError error
}

func smPartNew(title, uuid string) *smPart {
	if !validateUuid(uuid) {
		uuid = randomUuid()
	}

	return &smPart{
		uuid:         uuid,
		descriptions: map[string]*smPartDescription{},
		partData:     map[string]interface{}{},
	}
}

func (self *smPart) marshalPartData() error {
	m := map[string]interface{}{}
	mapDeepcopy(self.partData, m)
	delete(m, "uuid")

	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	self.partDataJson = string(b)

	return nil
}

func (self *smPart) unmarshalPartData() error {
	m := map[string]interface{}{}

	err := hjson.Unmarshal([]byte(self.partDataJson), &m)
	self.unmarshalError = err
	if err != nil {
		return err
	}

	m["uuid"] = self.uuid
	self.partData = m

	return nil
}

func (self *smPart) setUuid(uuid string) bool {
	if !validateUuid(uuid) {
		return false
	}

	self.uuid = strings.ToLower(uuid)

	return true
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
