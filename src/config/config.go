package config

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ProcessCaptions(source string) (map[string]string, error) {
	captions, err := ParseCaptions(string(source))
	if err != nil {
		return nil, err
	}

	caps, err := captions.BuildCaptions()
	return caps, err
}

type CaptionsConfiguration struct {
	Captions  map[string]string
	Variables map[string]string
}

func ParseCaptions(source string) (*CaptionsConfiguration, error) {
	c := new(CaptionsConfiguration)
	c.Captions = make(map[string]string)
	c.Variables = make(map[string]string)

	source = strings.Replace(source, "\r\n", "\n", -1)
	for i, line := range strings.Split(source, "\n") {
		line = strings.TrimLeftFunc(line, unicode.IsSpace) // remove leading space

		// ignore comments and empty lines
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || line == "" {
			continue
		}

		parts := strings.SplitN(line, ": ", 2)

		if len(parts) != 2 {
			return nil, errors.New("Wrong format in line " + strconv.Itoa(i))
		}

		if strings.HasPrefix(line, "$") {
			name := strings.TrimLeft(parts[0], "$")
			c.Variables[name] = parts[1]
			continue
		}

		c.Captions[parts[0]] = parts[1]
	}
	return c, nil
}

var variableRegex = regexp.MustCompile(`\$\(([\w\.]+)\)`)

func (c *CaptionsConfiguration) BuildCaptions() (map[string]string, error) {
	captions := make(map[string]string)
	for name, contents := range c.Captions {
		captions[name] = variableRegex.ReplaceAllStringFunc(contents, func(match string) string {
			varName := strings.TrimLeft(strings.TrimRight(match, ")"), "$(")
			if repl, ok := c.Variables[varName]; ok {
				return repl
			}
			fmt.Println("[WARNING] Could not find variable" + varName + " in caption " + name + ", IGNORING!")
			return match // just leave it as it is
		})
	}
	return captions, nil
}
