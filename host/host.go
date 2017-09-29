package host;

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"path"
)

type Host struct {
	domain string 
	props map[string]string
}

func NewHost(domain string) (*Host) {
	u := new(Host)
	u.domain = domain
	u.props = make(map[string]string)
	return u
}

func (t *Host) PrintProps() {
	for key, val := range t.props {
		fmt.Printf(" %s = %s\n", key, val)
	}
}

func (t *Host) HasProp(name string) bool {
	if _, val := t.props[name]; val {
		return true
	}
	return false
}

func (t *Host) SetProp(name string, val string) (*Host) {
	t.props[name] = val
	return t
}

func (t *Host) GetProp(name string) string {
	return t.props[name]
}

func LoadFromFile(file string) (*Host) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	u := NewHost(path.Base(file))

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		data := strings.SplitN(line, "=", 2)
		
		if len(data) == 2 {
			u.SetProp(strings.Trim(data[0], " "),  strings.Trim(data[1], " "))
		} else {
			log.Println("Unable to parse file ")
			log.Println(file)
			log.Println(data)
		}
	}

	return u
}








