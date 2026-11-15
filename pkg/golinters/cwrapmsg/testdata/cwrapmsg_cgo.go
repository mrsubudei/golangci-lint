//golangcitest:args -Ecwrapmsg
package testdata

/*
 #include <stdio.h>
 #include <stdlib.h>

 void myprint(char* s) {
 	printf("%d\n", s);
 }
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/pkg/errors"
)

func _() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

type cgoStore struct{}

func (s *cgoStore) GetByIDs(_ []int) ([]string, error) { return nil, nil }

func _(s *cgoStore) error {
	_, err := s.GetByIDs(nil)
	if err != nil {
		return errors.Wrap(err, "GetByIDs")
	}

	_, err = s.GetByIDs(nil)
	if err != nil {
		return errors.Wrap(err, "GetItems") // want "message in errors.Wrap does not match function or method name"
	}

	_, err = s.GetByIDs(nil)
	if err != nil {
		return fmt.Errorf("getItems: %w", err) // want "message in fmt.Errorf does not match function or method name"
	}

	return nil
}
