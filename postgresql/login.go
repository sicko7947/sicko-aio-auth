package postgresql

import "time"

func Login(key, ip, cpuId, timestamp string) (STATUSCODE, error) {

	entry := &keyDetails{
		Key: key,
	}

	has, err := eg.Get(entry)
	if err != nil {
		return DATABASE_ERROR, err
	}

	if has {
		time, _ := time.Parse("2006-01-02T15:04:05.000Z", timestamp)

		switch {
		case len(entry.IP) == 0 && len(entry.CpuId) == 0:
			entry.LastLoginTime = time
			entry.IP = ip
			entry.CpuId = cpuId
			eg.Update(entry)

			return OK, nil
		case ip == entry.IP && cpuId == entry.CpuId:
			entry.LastLoginTime = time
			eg.Update(entry)

			return OK, nil
		case ip != entry.IP, cpuId != entry.CpuId:
			return LOGGED_IN_OTHER_DEVICE, nil
		default:
			return NO_ACCESS, nil
		}
	}
	return NOT_FOUND, nil
}
