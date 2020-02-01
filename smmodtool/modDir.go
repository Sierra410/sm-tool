package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hjson/hjson-go"
)

var (
	inventoryDescriptionsJson = "inventoryDescriptions.json"
	shapeSets                 = "shapeSets.json"
)

type modDirectory struct {
	path string
}

func (self *modDirectory) setDir(p string) error {
	if !fileExists(path.Join(p, "description.json")) {
		return errors.New("description.json not found!")
	}

	self.path = p

	return nil
}

func (self modDirectory) loadParts() []*smPart {
	defer func() {
		if err := recover(); err != nil {
			logger.printlnImportant("LOADING OF THE MOD FAILED!\n", err)
		}
	}()

	ps := make([]*smPart, 0, 128)

	// ShapeSets
	dir := self.getShapeSetsDir()
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, fi := range fileInfo {
		if fi.IsDir() {
			continue
		}

		j := map[string]interface{}{}

		jsonFile := path.Join(dir, fi.Name())
		err := jsonLoadFile(jsonFile, &j)
		if err != nil {
			continue
		}

		for k, v := range j {
			if k != "blockList" && k != "partList" {
				continue
			}

			for _, part := range v.([]interface{}) {
				p := toSmPart(k, part.(map[string]interface{}))
				if p == nil {
					continue
				}

				ps = append(ps, p)
			}
		}
	}

	// Description
	for _, lang := range languages {
		dir := self.getInventoryDescriptionDir(lang)
		fileInfo, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, fi := range fileInfo {
			if fi.IsDir() {
				continue
			}

			descs := map[string]interface{}{}

			jsonFile := path.Join(dir, fi.Name())
			err := jsonLoadFile(jsonFile, &descs)
			if err != nil {
				continue
			}

			for _, p := range ps {
				desc := smPartDescription{}

				for k, v := range descs {
					if strings.EqualFold(k, p.uuid) {
						v := v.(map[string]interface{})

						title, ok := v["title"]
						if ok {
							desc.Title = title.(string)
						}

						description, ok := v["description"]
						if ok {
							desc.Description = description.(string)
						}

						keywords, ok := v["keywords"]
						if ok {
							kwi := keywords.([]interface{})
							kws := make([]string, len(kwi))

							for _, x := range kwi {
								kws = append(kws, x.(string))
							}

							desc.Keywords = kws
						}

						goto descSet
					}
				}

				continue

			descSet:
				_, ok := p.descriptions[lang]
				if ok {
					logger.printfImportant("[%s/%s] description overwritten!\n", p.uuid, lang)
				}

				p.descriptions[lang] = &desc
			}
		}
	}

	return ps
}

func (self modDirectory) getInventoryDescriptionDir(lang string) string {
	return path.Join(self.path, "Gui", "Language", lang)
}

func (self modDirectory) getShapeSetsDir() string {
	return path.Join(self.path, "Objects", "Database", "ShapeSets")
}

func jsonLoadFile(p string, j interface{}) error {
	if !strings.EqualFold(filepath.Ext(p), ".json") {
		return errors.New("Not .json")
	}

	b, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	hjson.Unmarshal(b, j)

	return nil
}

func toSmPart(kind string, j map[string]interface{}) *smPart {
	uuid, ok := j["uuid"]
	if !ok {
		return nil
	}

	p := smPartNew("", uuid.(string))
	p.kind = getPartKind(kind)

	delete(j, "uuid")

	m, _ := json.MarshalIndent(j, "", "\t")
	p.partData = string(m)

	return p
}

func fileExists(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return !info.IsDir()
}
