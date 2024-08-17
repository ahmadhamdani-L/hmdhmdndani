package configs

import (
	"strings"
	"github.com/spf13/viper"
)

var (
	fang       *viper.Viper
)

func init() {

	fang = viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.SetConfigType("yaml")
}
