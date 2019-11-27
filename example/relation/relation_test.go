package relation

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/noaway/godao"
)

func TestNewRelation(t *testing.T) {
	wrapTest(t, func(as *assert.Assertions) {
		r, err := NewRelation(1, 100)
		if as.NoError(err, "关注失败") {
			as.True(r.ID > 0, "r.ID should large than 0")
			as.True(r.Ts > 0, "r.Ts should large than 0")
			as.Equal(uint(1), r.Uid, "r.Uid should equal to 1")
			as.Equal(uint(100), r.RelUid, "r.RelUid should equal to 100")
		}

		newR, err := NewRelation(1, 100)
		if as.NoError(err, "关注失败") {
			as.Equal(r, newR, "多次关注,结果应该一样")
		}
	})
}

func TestDeleteRelation(t *testing.T) {
	wrapTest(t, func(as *assert.Assertions) {
		// delete relation not exists
		err := DeleteRelation(1, 100)
		as.NoError(err)

		r, err := NewRelation(1, 100)
		if !as.NoError(err) {
			as.FailNow(err.Error())
		}

		err = DeleteRelation(r.Uid, r.RelUid)
		as.NoError(err)

		var count int
		err = godao.Engine.Model(new(Relation)).
			Where("uid = ?", r.Uid).
			Where("rel_uid = ?", r.RelUid).
			Count(&count).Error
		as.NoError(err)
		as.Equal(0, count, "delete failed")
	})
}

func TestRelationExists(t *testing.T) {
	wrapTest(t, func(as *assert.Assertions) {
		r, err := RelationExists(1, nil)
		as.NoError(err)
		as.Len(r, 0)

		r, err = RelationExists(1, []uint{})
		as.NoError(err)
		as.Len(r, 0)

		r, err = RelationExists(1, []uint{100, 101})
		as.NoError(err)
		as.Len(r, 0)

		NewRelation(1, 100)
		NewRelation(1, 101)
		r, err = RelationExists(1, []uint{100, 101, 102})
		as.NoError(err)
		as.Len(r, 2)
		as.True(r[100])
		as.True(r[101])
		as.False(r[102])
	})
}

func TestGetFollowing(t *testing.T) {
	wrapTest(t, func(as *assert.Assertions) {
		noFollowing := [][3]uint{
			{1, 0, 0},
			{1, 1, 0},
			{1, 0, 10},
			{1, 1, 10},
		}
		for _, v := range noFollowing {
			r, err := GetFollowing(v[0], v[1], v[2])
			if as.NoError(err) {
				as.Len(r, 0)
			}
		}

		firstFollowing := uint(100)
		for i := 0; i < 14; i++ {
			NewRelation(1, firstFollowing)
			firstFollowing++
		}

		r, err := GetFollowing(1, 0, 10)
		if as.NoError(err) {
			if as.Len(r, 10) {
				as.Equal(uint(100), r[0])
				as.Equal(uint(109), r[9])
			}
		}

		r, err = GetFollowing(1, 109, 10)
		if as.NoError(err) {
			if as.Len(r, 4) {
				as.Equal(uint(110), r[0])
				as.Equal(uint(113), r[3])
			}
		}
	})
}

func TestGetFollower(t *testing.T) {
	wrapTest(t, func(as *assert.Assertions) {
		noFollower := [][3]uint{
			{1, 0, 0},
			{1, 1, 0},
			{1, 0, 10},
			{1, 1, 10},
		}
		for _, v := range noFollower {
			r, err := GetFollower(v[0], v[1], v[2])
			if as.NoError(err) {
				as.Len(r, 0)
			}
		}

		firstFollower := uint(100)
		for i := 0; i < 14; i++ {
			NewRelation(firstFollower, 1)
			firstFollower++
		}

		r, err := GetFollower(1, 0, 10)
		if as.NoError(err) {
			if as.Len(r, 10) {
				as.Equal(uint(100), r[0])
				as.Equal(uint(109), r[9])
			}
		}

		r, err = GetFollower(1, 109, 10)
		if as.NoError(err) {
			if as.Len(r, 4) {
				as.Equal(uint(110), r[0])
				as.Equal(uint(113), r[3])
			}
		}
	})
}

func wrapTest(t *testing.T, fn func(*assert.Assertions)) {
	godao.Engine.AutoMigrate(&Relation{})
	godao.Engine.Delete(&Relation{})
	fn(assert.New(t))
}

func init() {
	godao.InitTestORM()
}
