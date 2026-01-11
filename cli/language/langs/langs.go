package langs

import (
	"github.com/xschemadev/xschema/language"
	"github.com/xschemadev/xschema/language/langs/typescript"
)

func RegisterBuiltins() error {
	if err := language.Register(typescript.Language()); err != nil {
		return err
	}
	return nil
}
