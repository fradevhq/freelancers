package freelancers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
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

func testProfile(t *testing.T, path string) *freeLancer {
	data1, err := ioutil.ReadFile(path)
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
	return &fl1
}

func TestTemplateJSON(t *testing.T) {
	_ = testProfile(t, filepath.Join("TEMPLATE.json"))
	// TODO: Add more tempalte tests here to make sure all worksre
}

func TestProfiles(t *testing.T) {
	files, err := ioutil.ReadDir("./profiles")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		testProfile(t, filepath.Join("./profiles", f.Name()))
	}

}
