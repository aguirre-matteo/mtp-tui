package config

import (
  "github.com/aguirre-matteo/mtp-tui/checks"
  "github.com/spf13/viper"
)



func InitViper(user string) error {
  cfgPath, err := getCfgPath(user)
  if err != nil {
    return err
  }

  viper.SetConfigName("mtp-tui")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(cfgPath)

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
