package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Save and restore original notifyUser.
	savedNotifyUser := notifyUser
	defer func() { notifyUser = savedNotifyUser }()

	// Save and restore original bytesInUse.
	savedBytesInUse := bytesInUse
	defer func() { bytesInUse = savedBytesInUse }()

	// Install the test's fake notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	// Install the test's fake bytesInUse.
	// ...simulate a 980MB-used condition...
	bytesInUse = func(username string) int64 {
		return 980000000
	}

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
