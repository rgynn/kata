package kata

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	var result []string
	for _, person := range strings.Split(strings.ToUpper(s), ";") {
		names := strings.Split(person, ":")
		result = append(result, fmt.Sprintf("(%s, %s)", names[1], names[0]))
	}
	sort.Strings(result)
	return strings.Join(result, "")
}
