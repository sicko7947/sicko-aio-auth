package postgresql

func Ban(key string, reason string) (STATUSCODE, error) {

	entry := &keyMain{
		Key: key,
	}
	has, err := eg.Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		switch entry.Status {
		case 1:
			entry.Status = 3
			entry.Reason = reason
			eg.Update(entry)

			keyDetail := &keyDetails{
				Key: key,
			}
			eg.Delete(keyDetail)

			return OK, nil
		}
	}
	return NOT_FOUND, nil
}
