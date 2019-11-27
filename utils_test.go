package godao

import (
	"errors"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

func TestIsRecordNotFound(t *testing.T) {
	if !IsRecordNotFound(gorm.ErrRecordNotFound) {
		t.Error("IsRecordNotFound failed")
	}
	if IsRecordNotFound(errors.New("heheh")) {
		t.Error("IsRecordNotFound failed")
	}
}

func TestIsDuplicateEntry(t *testing.T) {
	if IsDuplicateEntry(nil) {
		t.Error("IsDuplicateEntry failed")
	}
	if IsDuplicateEntry(errors.New("hhh")) {
		t.Error("IsDuplicateEntry failed")
	}
	if !IsDuplicateEntry(&pq.Error{Code: "23505"}) {
		t.Error("IsDuplicateEntry failed")
	}
}

func TestMustAffectedRows(t *testing.T) {
	res := &gorm.DB{}
	err := MustAffectedRows(res)
	if err == nil {
		t.Error("MustAffectedRows must return error")
	}
	if err != ErrNoRowUpdated {
		t.Error("wrong err:", err, "; want:", ErrNoRowUpdated)
	}

	res.RowsAffected = 1
	err = MustAffectedRows(res)
	if err != nil {
		t.Error("wrong err:", err, "; want:", nil)
	}

	res.Error = errors.New("ttt")
	err = MustAffectedRows(res)
	if err != res.Error {
		t.Error("wrong err:", err, "; want:", res.Error)
	}
}
