package alfred

import (
	"encoding/json"
	"fmt"
)

type ScriptFilter struct {
	Items []ScriptFilterItem
}

type ScriptFilterItem struct {
	UID      string `json:"uid"`
	Arg      string
	Title    string
	Subtitle string
}

func (sf *ScriptFilter) Add(title, subtitle string) {
	var item = ScriptFilterItem{
		UID:      title,
		Arg:      title,
		Title:    title,
		Subtitle: subtitle,
	}
	sf.Items = append(sf.Items, item)
}

func (sf *ScriptFilter) Print() {
	sfJSON, err := json.MarshalIndent(sf, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sfJSON))

}
