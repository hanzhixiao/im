package service

import (
	"context"
	"im/apps/lbs/internal/config"
	"im/domain/cache"
	"im/domain/mrepo"
	"im/domain/repo"
	"im/pkg/proto/pb_lbs"
)

type LbsService interface {
	ReportLngLat(ctx context.Context, req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp, err error)
	PeopleNearby(ctx context.Context, req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp, err error)
}

type lbsService struct {
	cfg       *config.Config
	lbsRepo   mrepo.LbsRepository
	userRepo  repo.UserRepository
	locRepo   repo.UserLocationRepository
	userCache cache.UserCache
}

func NewLbsService(cfg *config.Config, lbsRepo mrepo.LbsRepository, userRepo repo.UserRepository, locRepo repo.UserLocationRepository, userCache cache.UserCache) LbsService {
	return &lbsService{cfg: cfg, lbsRepo: lbsRepo, userRepo: userRepo, locRepo: locRepo, userCache: userCache}
}
