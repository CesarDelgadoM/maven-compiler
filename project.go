package main

import "fmt"

//projects database
type project map[string][]string

func (p project) Keys() []string {

	keys := make([]string, 0, len(projects))
	for k := range projects {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	return keys
}

func (p project) Key(index int) string {

	keys := p.Keys()
	return keys[index-1]
}

func (p project) Value(key string) []string {

	return p[key]
}
