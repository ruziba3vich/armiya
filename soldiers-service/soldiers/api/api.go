package api

import (
	"log"
	"net"

	"github.com/ruziba3vich/armiya/soldies-service/genprotos"
	"github.com/ruziba3vich/armiya/soldies-service/internal/config"
	"google.golang.org/grpc"
)

type API struct {
	adminService       genprotos.AdminServiceServer
	soldierService     genprotos.SoldierServiceServer
	groupsService      genprotos.GroupsServiceServer
	attendancesService genprotos.AttendanceServiceServer
}

func New(
	adminService genprotos.AdminServiceServer,
	soldierService genprotos.SoldierServiceServer,
	groupsService genprotos.GroupsServiceServer,
	attendancesService genprotos.AttendanceServiceServer,
) *API {
	return &API{
		adminService:       adminService,
		soldierService:     soldierService,
		groupsService:      groupsService,
		attendancesService: attendancesService,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	genprotos.RegisterAdminServiceServer(serverRegisterer, a.adminService)
	genprotos.RegisterSoldierServiceServer(serverRegisterer, a.soldierService)
	genprotos.RegisterGroupsServiceServer(serverRegisterer, a.groupsService)
	genprotos.RegisterAttendanceServiceServer(serverRegisterer, a.attendancesService)
	log.Println("SERVER HAS STARTED ON PORT", config.Server.Port)
	return serverRegisterer.Serve(listener)
}
