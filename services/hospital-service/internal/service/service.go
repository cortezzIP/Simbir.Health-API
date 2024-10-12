package service

import (
	"context"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/cortezzIP/Simbir.Health-API/gen/go/hospital-service"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/model"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/repository"
)

type HospitalServiceServer struct {
	pb.UnimplementedHospitalServer
	repo *repository.HospitalRepo
}

func NewHospitalServiceServer(repo *repository.HospitalRepo) *HospitalServiceServer {
	return &HospitalServiceServer{
		repo: repo,
	}
}

func (s *HospitalServiceServer) CreateHospital(ctx context.Context, req *pb.CreateHospitalRequest) (*emptypb.Empty, error) {
	newHospital := model.Hospital{
		Name: req.Hospital.GetName(),
		Address: req.Hospital.GetAddress(),
		ContactPhone: req.Hospital.GetContactPhone(),
		Rooms: req.Hospital.GetRooms(),
	}

	err := s.repo.CreateHospital(ctx, &newHospital)
	if err != nil {
		slog.Error("Failed to create hospital: " + err.Error())
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}

func (s *HospitalServiceServer) GetHospitals(ctx context.Context, req *pb.GetAllHospitalsRequest) (*pb.GetAllHospitalsResponse, error) {
	resp, err := s.repo.GetHospitals(ctx, int(req.GetFrom()), int(req.GetCount()))
	if err != nil {
		slog.Error("Failed to get hospitals: " + err.Error())
		return nil, err
	}

	var hospitals []*pb.HospitalInfo

	for _, elem := range *resp {
		hospital := &pb.HospitalInfo{
			Id: int64(elem.Id),
			Name: elem.Name,
			Address: elem.Address,
			ContactPhone: elem.ContactPhone,
			Rooms: elem.Rooms,
		}
		hospitals = append(hospitals, hospital)
	}

	return &pb.GetAllHospitalsResponse{Hospitals: hospitals}, nil
}

func (s *HospitalServiceServer) GetHospitalById(ctx context.Context, req *pb.GetHospitalByIdRequest) (*pb.GetHospitalByIdResponse, error) {
	resp, err := s.repo.GetHopitalById(ctx, int(req.GetId()))
	if err != nil {
		slog.Error("Failed to get hospital by id: " + err.Error())
		return nil, err
	}

	hospital := &pb.HospitalInfo{
		Id: int64(resp.Id),
		Name: resp.Name,
		Address: resp.Address,
		ContactPhone: resp.ContactPhone,
		Rooms: resp.Rooms,
	}

	return &pb.GetHospitalByIdResponse{Hospital: hospital}, nil
}

func (s *HospitalServiceServer) GetHospitalRooms(ctx context.Context, req *pb.GetHospitalsRoomsRequest) (*pb.GetHospitalsRoomsResponse, error) {
	resp, err := s.repo.GetHospitalRooms(ctx, int(req.GetId()))
	if err != nil {
		slog.Error("Failed to get hospitals rooms: " + err.Error())
		return nil, err
	}

	return &pb.GetHospitalsRoomsResponse{Rooms: *resp}, nil
}

func (s *HospitalServiceServer) UpdateHospital(ctx context.Context, req *pb.UpdateHospitalRequest) (*emptypb.Empty, error) {
	newData := model.Hospital{
		Name: req.Hospital.GetName(),
		Address: req.Hospital.GetAddress(),
		ContactPhone: req.Hospital.GetContactPhone(),
		Rooms: req.Hospital.GetRooms(),
	}

	err := s.repo.UpdateHospital(ctx, int(req.GetId()), &newData)
	if err != nil {
		slog.Error("Failed to update hospital: " + err.Error())
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *HospitalServiceServer) DeleteHospital(ctx context.Context, req *pb.DeleteHospitalRequest) (*emptypb.Empty, error) {
	err := s.repo.DeleteHospital(ctx, int(req.GetId()))
	if err != nil {
		slog.Error("Failed to delete hospital: " + err.Error())
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}