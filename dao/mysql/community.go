package mysql

import (
	"database/sql"

	"bluebell/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`

	communityList = []*models.Community{}

	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (detail *models.CommunityDetail, err error) {
	detail = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id = ?`

	if err = db.Get(detail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}

	return
}
