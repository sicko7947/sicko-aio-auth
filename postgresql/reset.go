package postgresql

func Reset(key, email, discordID string) (STATUSCODE, error) {
	entry := &keyMain{
		Key:       key,
		Email:     email,
		DiscordID: discordID,
	}
	has, err := eg.Main().Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		switch entry.Status {
		case 0:
			return REQUIRE_ACTIVATION, nil
		default:
			keyDetail := &keyDetails{
				Key: key,
			}

			if has, _ := eg.Main().Get(keyDetail); has {
				eg.Main().Table(new(keyDetails)).ID(keyDetail.Id).Update(map[string]interface{}{
					"IP":    "",
					"CpuId": "",
				})
			}
			return OK, nil
		}
	}
	return WRONG_CREDENTIALS, nil
}
