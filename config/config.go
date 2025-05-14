package config

import (
	"github.com/aguirre-matteo/mtp-tui/checks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitViper(cmd *cobra.Command) error {
	user, err := cmd.Flags().GetString("user")
	if err != nil {
		return err
	}

	cfgPath, err := getCfgPath(user)
	if err != nil {
		return err
	}

	overrideCfgPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	viper.SetConfigName("mtp-tui")
	viper.SetConfigType("yaml")

	if overrideCfgPath != "unset" {
		err = assertConfigPath(overrideCfgPath)
		if err != nil {
			return err
		}
		viper.AddConfigPath(overrideCfgPath)
	}

	err = assertConfigPath(cfgPath)
	if err != nil {
		return err
	}
	viper.AddConfigPath(cfgPath)
	viper.AddConfigPath("/etc/")

	//Default settings
	viper.Set("user", user)
	defaultMnt, err := getDefaultMnt(user)
	if err != nil {
		return err
	}
	viper.SetDefault("mount.point", defaultMnt)
	viper.SetDefault("mount.options", "default_permissions,allow_other")

	err = checks.CheckAll()
	if err != nil {
		return err
	}

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return err
	}
	return nil
}
