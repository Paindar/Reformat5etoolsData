package dnd

import (
	"encoding/json"
	"os"
	"testing"
)

type spells struct {
	Spell []spell `json:"spell,omitempty"`
}

func TestReader(t *testing.T) {
	t.Run("spell_test", func(t *testing.T) {
		TEST_FILE_PATH := "./resources/data/spells-phb.json"
		raw_data, err := os.ReadFile(TEST_FILE_PATH)
		if err != nil {
			t.Error(err)
			return
		}
		var data spells
		err = json.Unmarshal(raw_data, &data)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(data.Spell[0].String())
	})
}
