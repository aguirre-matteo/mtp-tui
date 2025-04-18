package checks

import (
  "os"
  "os/user"
  "strconv"
  "github.com/spf13/viper"
)



func mkMnt(mountpoint, owner string) error {
  err := os.Mkdir(mountpoint, 0755)
	if err != nil {
		return err
	}

	usr, err := user.Lookup(owner)
	if err != nil {
		os.RemoveAll(mountpoint)
		return err
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		os.RemoveAll(mountpoint)
		return err
	}
	gid, err := strconv.Atoi(usr.Gid)
	if err != nil {
		os.RemoveAll(mountpoint)
		return err
	}

	err = os.Chown(mountpoint, uid, gid)
	if err != nil {
		os.RemoveAll(mountpoint)
		return err
	}
  return nil
}



func checkMnt() error {
  mountpoint := viper.GetString("mount.point")
  user := viper.GetString("user")
  _, err := os.Stat(mountpoint)

  if os.IsNotExist(err) {
    err = mkMnt(mountpoint, user)
  }
  return err
}
