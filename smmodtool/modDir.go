package main

import "path"

type modDirectory struct {
	path string
}

var (
	inventoryDescriptionsJson = "inventoryDescriptions.json"
	cumulativeJson            = "cum.json"
)

func (self modDirectory) getInventoryDescriptionDir(lang string) string {
	return path.Join(self.path, "Gui", "Language", lang)
}

func (self modDirectory) getShapeSetsDir() string {
	return path.Join(self.path, "Objects", "Database", "ShapeSets")
}
