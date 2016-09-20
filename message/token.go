package main

import "fmt"

//type Token struct {
//Key, Value string
//Expire     int64
//}

// format the key
func formatUidKey(uid int64) string {
	return fmt.Sprintf("uid_%d", uid)
}

// format the token
func formatTokenKey(token string) string {
	return fmt.Sprintf("token_%s", token)
}

// check if the token is valid
func IsTokenValid(uid int64, token string) (ok bool, err error) {
	uidReal, err := UseStorage.GetUid(token)
	if err != nil {
		return false, err
	}
	return uidReal == uid, nil
}
