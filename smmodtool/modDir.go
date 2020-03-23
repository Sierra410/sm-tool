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

const (
	defaultFileMode = 0644
	defaultDirMode  = 0755
)

var (
	inventoryDescriptionFilename = "inventoryDescriptions.json"
	shapeSetsFilename            = "shapeSets.json"
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

	logError := func(p string, err error) {
		logger.printfImportant("Error loading %s:\n%s\nSkipping...\n", p, err.Error())
	}

	ps := make([]*smPart, 0, 128)

	// ShapeSets
	jsonFiles, err := findJsonFiles(self.getShapeSetsDir())
	if err != nil {
		panic(err)
	}

	for _, jsonFile := range jsonFiles {
		j := map[string]interface{}{}

		err := jsonLoadFile(jsonFile, &j)
		if err != nil {
			logError(jsonFile, err)
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
		langDir := self.getInventoryDescriptionDir(lang)
		jsonFiles, err := findJsonFiles(langDir)
		if err != nil {
			logger.printf("Error reading %s\nSkipping...\n", langDir)
			continue
		}

		for _, jsonFile := range jsonFiles {
			descs := map[string]interface{}{}

			err := jsonLoadFile(jsonFile, &descs)
			if err != nil {
				logError(jsonFile, err)
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

func (self modDirectory) saveParts(parts []*smPart) error {
	descs := map[string]map[string]*smPartDescription{}
	for _, lang := range languages {
		descs[lang] = map[string]*smPartDescription{}
	}

	blockList := []interface{}{}
	partList := []interface{}{}

	// Convert to more sm-ish format
	for _, part := range parts {
		err := part.unmarshalPartData()
		if err != nil {
			logger.printlnImportant(part.uuid, "json error:\n", err.Error())
			return err
		}

		for lang, desc := range part.descriptions {
			if desc.Keywords == nil {
				desc.Keywords = []string{}
			}

			descs[lang][part.uuid] = desc
		}

		if part.kind {
			blockList = append(blockList, part.partData)
		} else {
			partList = append(partList, part.partData)
		}
	}

	// Marshall
	shapeSets, err := json.MarshalIndent(
		map[string]interface{}{
			"blockList": blockList,
			"partList":  partList,
		},
		"",
		"\t",
	)
	if err != nil {
		return err
	}

	inventoryDescriptions := map[string][]byte{}
	for lang, desc := range descs {
		b, err := json.MarshalIndent(desc, "", "\t")
		if err != nil {
			return err
		}

		inventoryDescriptions[lang] = b
	}

	shapeSetsDir := self.getShapeSetsDir()

	// All gud if we got here, delete old stuff
	jsons, err := findJsonFiles(shapeSetsDir)
	if err == nil {
		for _, jf := range jsons {
			os.Remove(jf)
		}
	}

	for _, lang := range languages {
		jsons, err := findJsonFiles(self.getInventoryDescriptionDir(lang))
		if err != nil {
			continue
		}

		for _, jf := range jsons {
			os.Remove(jf)
		}
	}

	//...and write new one
	os.MkdirAll(shapeSetsDir, defaultDirMode)
	ioutil.WriteFile(
		path.Join(shapeSetsDir, shapeSetsFilename),
		shapeSets,
		defaultFileMode,
	)

	for lang, desc := range inventoryDescriptions {
		descDir := self.getInventoryDescriptionDir(lang)
		os.MkdirAll(descDir, defaultDirMode)
		ioutil.WriteFile(
			path.Join(descDir, inventoryDescriptionFilename),
			desc,
			defaultFileMode,
		)
	}

	return nil
}

func (self modDirectory) getInventoryDescriptionDir(lang string) string {
	return path.Join(self.path, "Gui", "Language", lang)
}

func (self modDirectory) getShapeSetsDir() string {
	return path.Join(self.path, "Objects", "Database", "ShapeSets")
}

func findJsonFiles(dir string) ([]string, error) {
	files := []string{}

	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fi := range fileInfo {
		if fi.IsDir() {
			continue
		}

		if !strings.EqualFold(filepath.Ext(fi.Name()), ".json") {
			continue
		}

		files = append(files, path.Join(dir, fi.Name()))
	}

	return files, nil
}

func jsonLoadFile(p string, j interface{}) error {
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
	p.partDataJson = string(m)

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
