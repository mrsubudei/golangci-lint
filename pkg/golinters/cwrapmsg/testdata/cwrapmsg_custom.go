//golangcitest:args -Ecwrapmsg
//golangcitest:config_path testdata/cwrapmsg.yml
package testdata

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type customStore struct{}

func (s *customStore) GetByIDs(_ []int) ([]string, error) { return nil, nil }

type customService struct {
	store *customStore
}

func (s *customService) correct(ids []int) error {
	_, err := s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "store.getbyids")
	}

	_, err = strconv.Atoi("1")
	if err != nil {
		return errors.Wrapf(err, "atoi: %d", len(ids))
	}

	_, err = s.store.GetByIDs(ids)
	if err != nil {
		return fmt.Errorf("get_by_ids: %w", err)
	}

	_, err = s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "ошибка")
	}

	return nil
}

func (s *customService) incorrect(ids []int) error {
	_, err := s.store.GetByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "getitems") // want "message in errors.Wrap does not match function or method name"
	}

	_, err = strconv.Atoi("1")
	if err != nil {
		return fmt.Errorf("itoa: %w", err) // want "message in fmt.Errorf does not match function or method name"
	}

	return nil
}
