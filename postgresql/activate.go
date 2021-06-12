package postgresql

import (
	"time"
)

func Activate(key string, email string) (STATUSCODE, error) {

	entry := &keyMain{
		Key: key,
	}
	has, err := eg.Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		switch entry.Status {
		case 0:
			entry.Email = email
			entry.Status = 1
			entry.ActivateTime = time.Now()

			eg.ID(entry.Id).Update(entry)
			eg.Insert(&keyDetails{
				Key: key,
			})

			return OK, nil
		case 2:
			return CANCELLED, nil
		default:
			return ACTIVATED, nil
		}
	}
	return NOT_FOUND, nil
}
