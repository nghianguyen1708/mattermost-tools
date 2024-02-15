package config

import (
	"errors"
	"strings"

	"mattermost-tools/assets"
	string_helper "mattermost-tools/pkg/string-helper"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/fs"
)

const EnvPrefix = "APP__"

type AppConfig struct {
	HttpPort int      `koanf:"httpPort"`
	Db       DbConfig `koanf:"db"`
	Jwt      struct {
		PublicKey string `koanf:"publicKey"`
	} `koanf:"jwt"`
	ModelGeneration struct {
		Path             string   `koanf:"path"`
		IgnoredTables    []string `koanf:"ignoredTables"`
		ImmutableColumns []string `koanf:"immutableColumns"`
	} `koanf:"modelGeneration"`
	Mattermost MattermostConfig `koanf:"mattermost"`
}

type MattermostConfig struct {
	WebhookUrl    string `koanf:"webhookUrl"`
	MattermostApi string `koanf:"mattermostApi"`
	AdminToken    string `koanf:"adminToken"`
	UserToken     string `koanf:"userToken"`
}

type DbConfig struct {
	User        string `koanf:"user"`
	Password    string `koanf:"password"`
	DbName      string `koanf:"dbName"`
	Port        string `koanf:"port"`
	Host        string `koanf:"host"`
	EnableSsl   bool   `koanf:"enableSsl"`
	AutoMigrate bool   `koanf:"autoMigrate"`
}

func InitConfig() (AppConfig, error) {
	var config AppConfig
	k := koanf.New(".")
	configProvider := fs.Provider(assets.EmbeddedFiles, "config.yaml")
	if err := k.Load(configProvider, yaml.Parser()); err != nil {
		return config, errors.New("cannot read config from file")
	}
	if err := k.Load(
		env.Provider(
			EnvPrefix, ".", func(s string) string {
				return string_helper.SnakeToCamel(
					strings.Replace(
						strings.ToLower(
							strings.TrimPrefix(s, EnvPrefix),
						), "__", ".", -1,
					),
				)
			},
		), nil,
	); err != nil {
		return config, err
	}

	if err := k.Unmarshal("", &config); err != nil {
		return config, err
	}
	return config, nil
}
