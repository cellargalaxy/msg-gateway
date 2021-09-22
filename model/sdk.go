package model

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"time"
)

type ClientConfig struct {
	Retry   int           `yaml:"retry" json:"retry"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
	Sleep   time.Duration `yaml:"sleep" json:"sleep"`
	Address string        `yaml:"address" json:"address"`
	Secret  string        `yaml:"secret" json:"-"`
}

func (this ClientConfig) String() string {
	return util.ToJsonString(this)
}

type MsgHandlerInter interface {
	GetAddress(ctx context.Context) string
	GetSecret(ctx context.Context) string
}
