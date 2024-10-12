package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/database"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/model"
)

var (
	ErrHospitalNotFound = errors.New("hospital not found")
)

type HospitalRepository interface {
	CreateHospital(ctx context.Context, hospital *model.Hospital) error 
	GetHospitals(ctx context.Context, from int, count int) (*[]model.Hospital, error)
	GetHopitalById(ctx context.Context, id int) (*model.Hospital, error)
	GetHospitalRooms(ctx context.Context, id int) (*[]string, error)
	UpdateHospital(ctx context.Context, id int, hospital *model.Hospital) error
	DeleteHospital(ctx context.Context, id int) error
} 

type HospitalRepo struct {
	db *pgxpool.Pool
}

func NewHospitalRepository() *HospitalRepo {
	return &HospitalRepo{
		db: database.GetHospitalDB(),
	}
}

func (r *HospitalRepo) CreateHospital(ctx context.Context, hospital *model.Hospital) error {
	query := "INSERT INTO hospitals (name, address, contact_phone, rooms) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(ctx, query, hospital.Name, hospital.Address, hospital.ContactPhone, hospital.Rooms)
	if err != nil {
		return err
	}

	return nil
}

func (r *HospitalRepo) GetHospitals(ctx context.Context, from int, count int) (*[]model.Hospital, error) {
	query := "SELECT * FROM hospitals ORDER BY id LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(ctx, query, count, from)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrHospitalNotFound
		}

		return nil, err
	}
	defer rows.Close()

	var hospitals []model.Hospital

	for rows.Next() {
		var hospital model.Hospital

		err := rows.Scan(&hospital.Id, &hospital.Name, &hospital.Address, &hospital.ContactPhone, &hospital.Rooms)
		if err != nil {
			return nil, err
		}

		hospitals = append(hospitals, hospital)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &hospitals, nil
}

func (r *HospitalRepo) GetHopitalById(ctx context.Context, id int) (*model.Hospital, error) {
	query := "SELECT * FROM hospitals WHERE id = $1"

	var hospital model.Hospital
	
	err := r.db.QueryRow(ctx, query, id).Scan(&hospital.Id, &hospital.Name, &hospital.Address, &hospital.ContactPhone, &hospital.Rooms)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrHospitalNotFound
		}

		return nil, err
	}
	
	return &hospital, nil
}

func (r *HospitalRepo) GetHospitalRooms(ctx context.Context, id int) (*[]string, error) {
	query := "SELECT rooms FROM hospitals WHERE id = $1"

	var rooms []string

	err := r.db.QueryRow(ctx, query, id).Scan(&rooms)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrHospitalNotFound
		}

		return nil, err
	}

	return &rooms, nil
}

func (r *HospitalRepo) UpdateHospital(ctx context.Context, id int, hospital *model.Hospital) error {
	query := "UPDATE hospitals SET name = $1, address = $2, contact_phone = $3, rooms = $4 WHERE id = $5"

	commandTag, err := r.db.Exec(ctx, query, hospital.Name, hospital.Address, hospital.ContactPhone, hospital.Rooms, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrHospitalNotFound
	}

	return nil
}

func (r *HospitalRepo) DeleteHospital(ctx context.Context, id int) error {
	query := "DELETE FROM hospitals WHERE id = $1"

	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrHospitalNotFound
	}

	return nil
}