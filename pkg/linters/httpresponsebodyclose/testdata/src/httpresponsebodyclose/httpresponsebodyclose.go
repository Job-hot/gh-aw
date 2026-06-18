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
