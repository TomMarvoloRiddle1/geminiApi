package pkg

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// returns env pair value, when key is called without hardcoding secrets
func Getval_five(requested_key string) string {
	file_value := fmt.Sprintf("%s=", requested_key)
	return map_four()[file_value]
}

// combines prefix(key):suffix(pair) as map element
func map_four() map[string]string {
	element := make(map[string]string)
	key := prefix_env_two()
	pair := suffix_env_three()

	for i, value_key := range key {
		element[value_key] = pair[i]
	}

	return element
}

// creates suffix (pair) string slice, proportionally matched to prefix indexing
func suffix_env_three() []string {
	original_data := envdata_one()
	prefixes := prefix_env_two()

	var suf []string

	for i, v := range original_data {
		//removes suffix
		value := strings.ReplaceAll(v, prefixes[i], "")
		suf = append(suf, value)
	}

	return suf
}

// creates prefix (key) string slice, for matching
func prefix_env_two() []string {
	var prefixes []string

	for _, line_value := range envdata_one() {
		// seperator "=" included in prefix, to be removed later
		// note: can be repurposed as ":" for ie - user:pass, proxy:pass, etc.
		new_line_value := strings.SplitAfter(line_value, "=")
		if new_line_value[0] == "" {
			continue
		} else {
			prefixes = append(prefixes, new_line_value[0])
		}
	}

	return prefixes
}

// pulls data directly from file
func envdata_one() []string {
	env_rawdata, err := os.ReadFile(".env")

	if err != nil {
		fmt.Println("Issue loading .env file")
		log.Fatal(err)
	}

	list_keys := strings.Split(string(env_rawdata), "\n")
	final_list_keys := slices.Delete(list_keys, (len(list_keys) - 1), len(list_keys))
	return final_list_keys
}
