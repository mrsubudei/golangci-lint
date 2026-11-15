//golangcitest:args -Ecwrapmsg
package testdata

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type store struct{}

func (s *store) GetByIDs(_ []int) ([]string, error) { return nil, nil }

type service struct {
	store *store
}

func (s *service) correct(ids []int) error {
	_, err := s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "store.GetByIDs")
	}

	_, err = strconv.Atoi("1")
	if err != nil {
		return errors.Wrapf(err, "strconv.Atoi: %d", len(ids))
	}

	_, err = s.store.GetByIDs(ids)
	if err != nil {
		return fmt.Errorf("GetByIDs: %w", err)
	}

	_, err = s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "Get_By_IDs")
	}

	return nil
}

func (s *service) incorrect(ids []int) error {
	_, err := s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "GetItems") // want "message in errors.Wrap does not match function or method name"
	}

	_, err = strconv.Atoi("1")
	if err != nil {
		return errors.Wrapf(err, "atoi: %d", len(ids)) // want "message in errors.Wrapf does not match function or method name"
	}

	_, err = s.store.GetByIDs(ids)
	if err != nil {
		return fmt.Errorf("getbyids: %w", err) // want "message in fmt.Errorf does not match function or method name"
	}

	_, err = s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "ошибка") // want "message in errors.Wrap does not match function or method name"
	}

	return nil
}
