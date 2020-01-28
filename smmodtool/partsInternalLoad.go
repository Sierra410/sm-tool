package main

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/hjson/hjson-go"
)

func jsonLoadFile(p string) map[string]interface{} {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil
	}

	j := map[string]interface{}{}

	hjson.Unmarshal(b, &j)

	return j
}

func jsonLoadEveryFileInDirectory(dirPath string, f func(map[string]interface{})) {
	fileInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, x := range fileInfo {
		if !x.IsDir() {
			j := jsonLoadFile(path.Join(dirPath, x.Name()))
			if j == nil {
				continue
			}

			f(j)
		}
	}
}

func toSmPart(kind string, j map[string]interface{}) *smPart {
	uuid, ok := j["uuid"]
	if !ok {
		return nil
	}

	p := smPartNew(uuid.(string))
	p.kind = kind

	delete(j, "uuid")

	m, _ := json.Marshal(j)
	p.partData = string(m)

	return p
}

func smJsonLoadParts() {
	ps := make([]*smPart, 0, 512)

	jsonLoadEveryFileInDirectory(
		"",
		func(j map[string]interface{}) {
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
		},
	)

}
