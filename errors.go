package util

import "fmt"

func Catch(errFuncs ...func() error) error {
	for _, fn := range errFuncs {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func MultiErrors(errs ...error) (err error) {
	for _, e := range errs {
		if e != nil {
			err = fmt.Errorf("%w\n", e)
		}
	}
	return
}

func DeferMultiErrors(err *error, errFunc func() error) {
	if _err := errFunc(); _err != nil {
		*err = MultiErrors(*err, _err)
	}
}
