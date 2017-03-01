package Togo

import (
	"os"
)

func BaseName(path, suffix string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	return fi.Name(), nil
}

func Chdir(dir string) error {
	return os.Chdir(dir)
}

//func Chgrp(filename string,group interface{})error{
//	uid := os.Geteuid()
//	switch group.(type) {
//	case string:
//		grp := os.
//	case int:
//		return os.Chown(filename,uid,group.(int))
//	default:
//		return errors.New("unsupported group type")
//
//	}
//	os.Chown()
//}

func Chmod(filename string, mode int) error {
	return os.Chmod(filename, os.FileMode(mode))
}


//func Chown(filename string,user interface{})error{
//	switch user.(type) {
//		case string:
//			grp := os.
//		case int:
//			return os.Chown(filename,uid,group.(int))
//		default:
//			return errors.New("unsupported user type")
//		}
//	return os.Chown()
//}