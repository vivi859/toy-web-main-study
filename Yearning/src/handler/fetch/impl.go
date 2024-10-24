package fetch

import (
	"Yearning-go/src/handler/common"
	"Yearning-go/src/model"
	"fmt"
)

const (
	UNDO_EXPR = "username =? AND work_id =? AND `status` =? "
)

type referOrder struct {
	Data model.CoreSqlOrder `json:"data"`
	SQLs string             `json:"sqls"`
	Tp   int                `json:"tp"`
}

type PageSizeRef struct {
	WorkId   string `json:"work_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type _FetchBind struct {
	IDC      string             `json:"idc"`
	Tp       string             `json:"tp"`
	Source   string             `json:"source"`
	SourceId string             `json:"source_id"`
	DataBase string             `json:"data_base"`
	Table    string             `json:"table"`
	Rows     []common.FieldInfo `json:"rows"`
	Idx      []common.IndexInfo `json:"idx"`
	Hide     bool               `json:"hide"`
}

type advisorFrom struct {
	SourceID string   `json:"source_id"`
	Schema   string   `json:"data_base"`
	Tables   []string `json:"tables"`
	SQL      string   `json:"sql"`
	Desc     string   `json:"desc"`
}

type ShowCreateTable struct {
	CreateTable string `gorm:"column:Create Table"`
}

func (a *advisorFrom) Go() (tables []string, err error) {
	var dataSource model.CoreDataSource
	model.DB().Model(model.CoreDataSource{}).Where("source_id =?", a.SourceID).First(&dataSource)
	db, err := dataSource.ConnectDB(a.Schema)
	if err != nil {
		return nil, err
	}
	defer model.Close(db)
	for _, i := range a.Tables {
		var result ShowCreateTable
		err = db.Raw(fmt.Sprintf("SHOW CREATE TABLE %s.%s", a.Schema, i)).Scan(&result).Error
		if err != nil {
			return nil, fmt.Errorf("failed to execute query: %v", err)
		}
		tables = append(tables, result.CreateTable)
	}
	return tables, nil
}

func (u *_FetchBind) FetchTableFieldsOrIndexes() error {
	var s model.CoreDataSource

	model.DB().Where("source_id =?", u.SourceId).First(&s)

	db, err := s.ConnectDB(u.DataBase)
	if err != nil {
		return err
	}

	defer model.Close(db)

	if err := db.Raw(fmt.Sprintf("SHOW FULL FIELDS FROM `%s`.`%s`", u.DataBase, u.Table)).Scan(&u.Rows).Error; err != nil {
		return err
	}

	if err := db.Raw(fmt.Sprintf("SHOW INDEX FROM `%s`.`%s`", u.DataBase, u.Table)).Scan(&u.Idx).Error; err != nil {
		return err
	}
	return nil
}
