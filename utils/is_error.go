package utils

import "log"

func IsError(err error) error {
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
