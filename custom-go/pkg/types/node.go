package types

import (
	"custom-go/pkg/utils"
	"github.com/joho/godotenv"
)

const (
	nodeConfigFilepath = ""
	nodeEnvFilepath    = ""
)

func init() {
	_ = utils.ReadStructAndCacheFile(nodeConfigFilepath, &WdgGraphConfig)
	_ = godotenv.Overload(nodeEnvFilepath)
}
