package types

import (
	"encoding/json"
)

type multiString []string

func (ms *multiString) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		switch data[0] {
		case '"':
			var s string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}

			*ms = multiString([]string{})
		case '[':
			var s []string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*ms = multiString(s)
		}
	}
	return nil
}

/*func (ms multiString) MarshalJSON() ([]byte, error) {

	if len(ms) == 1 && (ms)[0] == "" {
		var arr []string
		return json.Marshal(arr)
	} else {

		return json.Marshal(ms)
	}
}*/
