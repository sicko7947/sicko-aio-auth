package postgresql

func Ban(key string, reason string) (STATUSCODE, error) {

	entry := &keyMain{
		Key: key,
	}
	has, err := eg.Main().Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		switch entry.Status {
		case 1:
			entry.Status = 3
			entry.Reason = reason
			eg.Main().ID(entry.Id).Update(entry)

			keyDetail := &keyDetails{
				Key: key,
			}
			eg.Main().Delete(keyDetail)

			return OK, nil
		}
	}
	return NOT_FOUND, nil
}
