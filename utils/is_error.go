package utils

func IsError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
