package postgresql

func Polling(key, ip, cpuId string) (STATUSCODE, error) {

	entry := &keyDetails{
		Key:   key,
		IP:    ip,
		CpuId: cpuId,
	}

	has, err := eg.Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if !has {
		return KEY_STATUS_NOT_MATCH, nil
	}
	return OK, nil
}
