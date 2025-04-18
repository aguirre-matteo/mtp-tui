package config

import (
  "os/user"
)



func getUserHome(username string) (string, error) {
  usr, err := user.Lookup(username)
  if err != nil {
    return "", err
  }

  return usr.HomeDir, nil
}



func getCfgPath(user string) (string, error) {
  if user == "root" {
    return "/etc/", nil
  } else {
    home, err := getUserHome(user)
    if err != nil {
      return "", err
    }
    return home + "/.config/", nil
  }
}



func getDefaultMnt(user string) (string, error) {
  if user == "root" {
    return "/mtp/", nil
  } else {
    home, err := getUserHome(user)  
    if err != nil {
      return "", err
    }

    return home + "/mtp/", nil
  }
}
