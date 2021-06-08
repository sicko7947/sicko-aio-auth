package postgresql

func Deactivate(key string) (STATUSCODE, error) {

	entry := &keyDetails{
		Key: key,
	}

	has, err := eg.Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		entry.IP = ""
		entry.MAC = ""
		eg.Update(entry)

		return OK, nil
	}

	return NOT_FOUND, nil
}
