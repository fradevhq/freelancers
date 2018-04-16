package freelancers

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type freeLancer struct {
	Name      string            `json:"name"`
	Title     string            `json:"title"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Desc      string            `json:"desc"`
	Website   string            `json:"website"`
	City      []string          `json:"city"`
	Keywords  []string          `json:"keywords"`
	Profiles  map[string]string `json:"profiles"`
	Available bool              `json:"available"`
}

func TestTemplateJSON(t *testing.T) {
	data1, err := ioutil.ReadFile("TEMPLATE.json")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(data1) == 0 {
		t.Errorf("Expected len(data) > 0, got %v", len(data1))
	}
	var fl1 freeLancer
	if err := json.Unmarshal(data1, &fl1); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	data2, err := json.MarshalIndent(fl1, "", "   ")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(data1) == len(data2) {
		t.Errorf("Expected\n%v, got\n%v", string(data1), string(data2))
	}
}
