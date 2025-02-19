package nineutils

import "os/user"

func GetCurrentUser() (string, string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", "", err
	}
	username := currentUser.Username
	group, err := user.LookupGroupId(currentUser.Gid)
	if err != nil {
		return "", "", err
	}
	groupname := group.Name

	return username, groupname, nil
}
