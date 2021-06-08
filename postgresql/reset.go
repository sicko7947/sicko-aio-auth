package postgresql

func Reset(key, email string) (STATUSCODE, error) {

	entry := &keyMain{
		Key:   key,
		Email: email,
	}
	has, err := eg.Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		switch entry.Status {
		case 0:
			return REQUIRE_ACTIVATION, nil
		default:
			keyDetail := &keyDetails{
				Key:   key,
				IP:    "",
				CpuId: "",
			}
			eg.Update(keyDetail)

			return OK, nil
		}
	}
	return WRONG_CREDENTIALS, nil
}
