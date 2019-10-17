package layer

import "encoding/json"

// Layer expresses Layer architecture.
type Layer struct {
	Packages []string `json:"Packages"`
	Inside   *Layer   `json:"Inside"`
	Raw      []interface{}
}

// UnmarshalJSON unmarshals JSON data by custom logic.
func (l *Layer) UnmarshalJSON(data []byte) error {
	// まず引数のdataとして渡ってきたJSONから`[]interface{}`を取得する
	var raw []interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	fillLayer(raw, l)

	return nil
}

func fillLayer(raw []interface{}, l *Layer) {
	l.Raw = raw
	for _, e := range raw {
		switch e := e.(type) {
		case string:
			l.Packages = append(l.Packages, e)
		case []interface{}:
			i := &Layer{}
			fillLayer(e, i)
			l.Inside = i
		}
	}
}
