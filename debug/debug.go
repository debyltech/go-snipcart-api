package debug

import "fmt"

type Debug struct {
	Enabled bool `json:"enabled"`
}

func (d *Debug) Printf(s string, a ...any) {
	if !d.Enabled {
		return
	}

	fmt.Printf(s, a...)
}

func (d *Debug) Println(a ...any) {
	if !d.Enabled {
		return
	}

	fmt.Println(a...)
}
