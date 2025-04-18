package checks



func CheckAll() error {
  err := checkDependencies()
  if err != nil {
    return err
  }

  err = checkMnt()
  if err != nil {
    return err
  }
  return nil
}
