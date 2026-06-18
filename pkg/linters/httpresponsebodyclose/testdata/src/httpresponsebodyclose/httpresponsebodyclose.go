package httpresponsebodyclose

import (
	"io"
	"net/http"
)

// flagged: Body.Close() called manually, not deferred.
func fetchManualClose(url string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req) // want `HTTP response Body\.Close\(\) should be deferred immediately after error check to prevent resource leaks`
	if err != nil {
		return err
	}
	_, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	return nil
}

// not flagged: Body.Close() is deferred correctly.
func fetchDeferred(url string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.ReadAll(resp.Body)
	return nil
}

// flagged: Body.Close() return value captured but not deferred.
func fetchManualCloseWithErr(url string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req) // want `HTTP response Body\.Close\(\) should be deferred immediately after error check to prevent resource leaks`
	if err != nil {
		return err
	}
	_, _ = io.ReadAll(resp.Body)
	closeErr := resp.Body.Close()
	return closeErr
}

// flagged: Body.Close() captured in if-init, not deferred.
func fetchIfInitClose(url string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req) // want `HTTP response Body\.Close\(\) should be deferred immediately after error check to prevent resource leaks`
	if err != nil {
		return err
	}
	if closeErr := resp.Body.Close(); closeErr != nil {
		return closeErr
	}
	return nil
}

// not flagged: closure is analyzed independently and its defer must not be
// attributed to the outer function.
func fetchInClosure(url string) {
	fn := func() error {
		resp, err := http.DefaultClient.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		_, _ = io.ReadAll(resp.Body)
		return nil
	}
	_ = fn()
}

// flagged: same response variable reassigned after a manual close; the first
// response acquisition must still be reported.
func fetchReassignAfterManualClose(url string) error {
	resp, err := http.DefaultClient.Get(url) // want `HTTP response Body\.Close\(\) should be deferred immediately after error check to prevent resource leaks`
	if err != nil {
		return err
	}
	resp.Body.Close()
	resp, err = http.DefaultClient.Get(url + "/retry")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.ReadAll(resp.Body)
	return nil
}

// not flagged: no Body.Close() call at all (different concern, out of scope).
func fetchNoClose(url string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, _ = io.ReadAll(resp.Body)
	return nil
}
