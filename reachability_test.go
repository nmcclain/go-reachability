package reachability

import (
	"fmt"
	"testing"
	"time"
)

var testGoodHost = "google.com"
var testFailHost = "127.0.99.99"
var testBadDnsHost = "baddns.cooooom"
var testPort = 80

func TestIsReachable(t *testing.T) {
	if err := IsReachable(testGoodHost, testPort); err != nil {
		t.Fatalf("IsReachable failed: %v", err)
	}
}
func TestIsReachableShouldError(t *testing.T) {
	err := IsReachable("", testPort)
	want := "foo"
	if err == nil {
		t.Fatalf("empty host should have failed")
	} else if err.Error() != "must specify a host" {
		t.Fatalf("empty host failed with wrong error [want: %v | got: %v]", want, err)
	}

	want = fmt.Sprintf("TCP connection error: dial tcp %s:%d: i/o timeout", testFailHost, testPort)
	if err := IsReachableTimeout(testFailHost, testPort, time.Second*1); err == nil {
		t.Fatalf("IsReachable should have failed but didnt: %v", err)
	} else if err.Error() != want {
		t.Fatalf("IsReachable failed with wrong error [want: %v | got: %v]", want, err)
	}

	want = fmt.Sprintf("TCP connection error: dial tcp: i/o timeout")
	if err := IsReachableTimeout(testGoodHost, testPort, time.Millisecond); err == nil {
		t.Fatalf("IsReachable should have timed out but didnt: %v", err)
	} else if err.Error() != want {
		t.Fatalf("IsReachable timed out with wrong error [want: %v | got: %v]", want, err)
	}

	want = fmt.Sprintf("TCP connection error: dial tcp: lookup %s: no such host", testBadDnsHost)
	if err := IsReachable(testBadDnsHost, testPort); err == nil {
		t.Fatalf("IsReachable should have failed on dns but didnt: %v", err)
	} else if err.Error() != want {
		t.Fatalf("IsReachable failed on dns with wrong error [want: %v | got: %v]", want, err)
	}
}
