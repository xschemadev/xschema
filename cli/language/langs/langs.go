package langs

import (
	"github.com/xschemadev/xschema/language"
	_ "github.com/xschemadev/xschema/language/langs/python"
	"github.com/xschemadev/xschema/language/langs/typescript"
)

func RegisterBuiltins() error {
	if err := language.Register(typescript.Language()); err != nil {
		return err
	}
	return nil
}
