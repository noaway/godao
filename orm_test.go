package godao

import "testing"

func TestParseURL(t *testing.T) {
	config, err := parseURL("postgres://1:123@localhost:10001/aaa?sslmode=false&ShowSQL=true&MaxOpenConn=1&MaxIdleConn=1")
	if err != nil {
		t.Error(err)
	}

	if config.Dialect != "postgres" {
		t.Errorf("bad dialect %s", config.Dialect)
	}

	if config.ShowSQL != true {
		t.Errorf("bad show sql %v", config.ShowSQL)
	}

	if config.URL != "postgres://1:123@localhost:10001/aaa?sslmode=false" {
		t.Errorf("bad url %s", config.URL)
	}

	if config.MaxIdleConn != 1 {
		t.Errorf("bad max idle conn %d", config.MaxIdleConn)
	}
}
