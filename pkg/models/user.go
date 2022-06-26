package models

import (
	"github.com/ankogit/wwc_cocksize_bot--api/api"

	"time"
)

type UserData struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	CockSize  int
	Time      time.Time
	//ChatID    int64
}

type UserDataResponse struct {
}

func (i *UserDataResponse) ArrayToProto(userStats []UserData) *api.GetStatsResp {
	var statsEntities []*api.StatsEntity
	for _, model := range userStats {
		statsEntities = append(statsEntities, model.ToProto())
	}
	return &api.GetStatsResp{
		StatsEntities: statsEntities,
	}
}

func (i *UserData) ToProto() *api.StatsEntity {
	return &api.StatsEntity{
		Id:        i.ID,
		Username:  i.Username,
		FirstName: i.FirstName,
		LastName:  i.LastName,
		CockSize:  int32(i.CockSize),
	}
}
