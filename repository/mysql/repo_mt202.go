package mysql

import (
	"newgens/models"
	"newgens/repository"

	"github.com/jmoiron/sqlx"
)

type mt202RepoMysql struct {
	db *sqlx.DB
}

func NewRepoMt202Mysql(db *sqlx.DB) repository.MT202Repo {
	return &mt202RepoMysql{
		db: db,
	}
}

func (r *mt202RepoMysql) GetData() ([]*models.MT202, error) {
	var data []*models.MT202

	if err := r.db.Select(&data, "SELECT * FROM mt202"); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *mt202RepoMysql) InsertData(data *models.MT202) error {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err := r.db.NamedExec(`INSERT INTO mt202 (
		SenderBIC,
		ReceiverBIC,
		Direction,
		MTType,
		UETR,
		F20,
		F21,
		F32A_ValueDate,
		F32A_Currency,
		F32A_Amount,
		F52a,
		F57a,
		F58a,
		RawData,
		created_at
	) VALUES(
		:SenderBIC,
		:ReceiverBIC,
		:Direction,
		:MTType,
		:UETR,
		:F20,
		:F21,
		:F32A_ValueDate,
		:F32A_Currency,
		:F32A_Amount,
		:F52a,
		:F57a,
		:F58a,
		:RawData,
		NOW()
	)`, data)
	if err != nil {
		return err
	}

	return nil
}
