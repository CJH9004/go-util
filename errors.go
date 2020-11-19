package util

func Catch(errFuncs ...func() error) error {
	for _, fn := range errFuncs {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
