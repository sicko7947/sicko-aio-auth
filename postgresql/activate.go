package postgresql

import (
	"time"
)

func Activate(key, email, discordID string) (STATUSCODE, error) {

	entry := &keyMain{
		Key: key,
	}
	has, err := eg.Main().Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		switch entry.Status {
		case 0:
			entry.Email = email
			entry.Status = 1
			entry.DiscordID = discordID
			entry.ActivateTime = time.Now()

			eg.Main().ID(entry.Id).Update(entry)
			eg.Main().Insert(&keyDetails{
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
