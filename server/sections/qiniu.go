package sections

import "github.com/QN-Resources/server/config"

type qiniu struct {
	AccessKey string `yaml:"AccessKey,omitempty"`
	SecretKey string `yaml:"SecretKey,omitempty"`
}

func (s *qiniu) SectionName() string {
	return "qiniu"
}

var Qiniu = &qiniu{}

func init() {
	config.Load(Qiniu)
}
