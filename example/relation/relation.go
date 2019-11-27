package relation

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/noaway/godao"
)

type Relation struct {
	ID     uint
	Uid    uint
	RelUid uint
	Ts     int64
}

func NewRelation(uid, relUid uint) (*Relation, error) {
	r := &Relation{
		Uid:    uid,
		RelUid: relUid,
		Ts:     time.Now().Unix(),
	}

	res := godao.Engine.Where("uid = ?", uid).Where("rel_uid = ?", relUid).First(r)
	if res.Error != nil {
		if res.RecordNotFound() {
			if err := godao.Engine.Create(r).Error; err != nil {
				return nil, err
			}
			return r, nil
		}
		return nil, res.Error
	}
	return r, nil
}

func DeleteRelation(uid, relUid uint) error {
	return godao.Engine.Exec("DELETE FROM relation WHERE uid = ? AND rel_uid = ?", uid, relUid).Error
}

func RelationExists(uid uint, relUids []uint) (map[uint]bool, error) {
	ret := map[uint]bool{}
	if len(relUids) == 0 {
		return ret, nil
	}
	uids := []uint{}
	err := godao.Engine.Table("relation").
		Where("uid = ?", uid).
		Where("rel_uid in (?)", relUids).
		Pluck("rel_uid", &uids).Error
	if err != nil {
		return nil, err
	}
	for _, u := range uids {
		ret[u] = true
	}
	return ret, nil
}

func GetFollowing(uid, lastRelUid, limit uint) (ret []uint, err error) {
	lastID, err := getIDByUidAndRelUid(uid, lastRelUid)
	if err != nil {
		return
	}

	ret, err = getUidList(
		godao.Engine.Where("uid = ?", uid),
		"rel_uid", lastID, limit,
	)
	return
}

func GetFollower(uid uint, lastRelUid, limit uint) (ret []uint, err error) {
	lastID, err := getIDByUidAndRelUid(lastRelUid, uid)
	if err != nil {
		return
	}

	ret, err = getUidList(
		godao.Engine.Where("rel_uid = ?", uid),
		"uid", lastID, limit,
	)
	return
}

func getIDByUidAndRelUid(uid, relUid uint) (uint, error) {
	if relUid == 0 || uid == 0 {
		return 0, nil
	}

	var id []uint
	err := godao.Engine.Table("relation").
		Where("uid = ?", uid).
		Where("rel_uid = ?", relUid).
		Pluck("id", &id).Error
	if err != nil {
		return 0, err
	}
	if len(id) > 0 {
		return id[0], nil
	} else {
		return 0, nil
	}
}

func getUidList(db *gorm.DB, col string, lastId, limit uint) ([]uint, error) {
	ret := []uint{}
	if lastId > 0 {
		db = db.Where("id > ?", lastId)
	}
	err := db.Table("relation").Limit(limit).Order("id").Pluck(col, &ret).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
