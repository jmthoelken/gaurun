package gaurun

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Sirupsen/logrus"
)

func init() {
	LogAccess = logrus.New()
	LogError = logrus.New()
	LogAccess.Formatter = new(GaurunFormatter)
	LogError.Formatter = new(GaurunFormatter)
	LogAccess.Out = ioutil.Discard
	LogError.Out = ioutil.Discard
}

func BenchmarkLogPushIOSOmitempty(b *testing.B) {
	req := RequestGaurunNotification{
		Platform: PlatFormIos,
	}
	errPush := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		LogPush(uint64(100), StatusAcceptedPush, "xxx", 0.123, req, errPush)
	}
}

func BenchmarkLogPushIOSFull(b *testing.B) {
	req := RequestGaurunNotification{
		Platform:         PlatFormIos,
		Badge:            1,
		Sound:            "foo",
		ContentAvailable: true,
		Expiry:           100,
	}
	errPush := fmt.Errorf("error")
	for i := 0; i < b.N; i++ {
		LogPush(uint64(100), StatusAcceptedPush, "xxx", 0.123, req, errPush)
	}
}
