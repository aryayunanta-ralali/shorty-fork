// Package validator
package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type RuleFunc func(field string, rule string, message string, value interface{}) error

func Rules(envi string) map[string]RuleFunc {

	emailPattern := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$"

	rules := map[string]RuleFunc{
		"phone_number": func(field string, rule string, message string, value interface{}) error {
			v := util.ToString(value)
			if len(v) == 0 {
				return nil
			}

			pattern := `^\+?([ -]?\d+)+|\(\d+\)([ -]\d+)$`
			rgx, err := regexp.Compile(pattern)
			if err != nil {
				return fmt.Errorf("The %s invalid phone number", field)
			}

			ok := rgx.MatchString(v)

			if !ok {
				return fmt.Errorf("The %s invalid phone number", field)
			}

			return nil
		},

		"reference_id": func(field string, rule string, message string, value interface{}) error {
			v := util.ToString(value)

			if len(v) == 0 {
				return nil
			}

			msg := fmt.Sprintf(`The %s is invalid format`, field)

			pattern := `^[0-9a-zA-Z\-_\\/]+$`
			rgx, err := regexp.Compile(pattern)

			if len(strings.Trim(message, " ")) > 0 {
				msg = message
			}

			if err != nil {
				return fmt.Errorf(msg)
			}

			ok := rgx.MatchString(v)

			if !ok {
				return fmt.Errorf(msg)
			}

			return nil
		},

		"subscriber_id": func(field string, rule string, message string, value interface{}) error {
			v := util.ToString(value)

			if len(v) == 0 {
				return nil
			}

			msg := fmt.Sprintf(`The %s is invalid format`, field)

			pattern := `^[0-9a-zA-Z]+$`

			err := Match(v, field, pattern, msg)
			err2 := Match(v, field, emailPattern, msg)

			if err != nil && err2 != nil {
				return fmt.Errorf(msg)
			}

			return nil
		},
	}

	return rules

}

// Match regular expression validation
func Match(value interface{}, key, format, msg string) error {

	rgx, e := regexp.Compile(format)

	if e != nil {
		return fmt.Errorf("%s invalid rule regular expression %s: %s", key, format, e.Error())
	}

	val, ok := value.(string)

	if !ok {
		return fmt.Errorf("%s invalid type, expected string found %s", key, reflect.TypeOf(value))
	}

	if !rgx.MatchString(val) {
		if msg == "" {
			return fmt.Errorf("%s has invalid format value", key)
		}

		return fmt.Errorf(msg)
	}

	return nil
}
