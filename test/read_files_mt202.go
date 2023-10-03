package test

import (
	"newgens/config"
	"newgens/models"
	"newgens/repository/mysql"
	"newgens/src"
	"strings"
	"testing"
)

func TestReadFilesMt202(t *testing.T) {
	configs := config.GetConfig()

	db := src.ConnectMysql(configs)
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	data1, err := src.ReadLines("files/202MEP_52A_57A_58A.txt")
	if err != nil {
		t.Fatal(err)
	}

	data2, err := src.ReadLines("files/202MEP_52D_57D_58A.txt")
	if err != nil {
		t.Fatal(err)
	}

	data1Ready, err := models.NewMt202RawFromFile(strings.Join(data1, "\n"))
	if err != nil {
		t.Fatal(err)
	}

	data2Ready, err := models.NewMt202RawFromFile(strings.Join(data2, "\n"))
	if err != nil {
		t.Fatal(err)
	}

	type testStruct struct {
		name     string
		input    *models.MT202Raw
		hasError bool
	}
	var tests = []*testStruct{
		{"data1", data1Ready, false},
		{"data2", data2Ready, false},
	}

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !configs.InsertTestingToDatabase && !tt.hasError {
				t.Skip("Not run")
			}
			mt202, err := models.NewMT202FromRaw(tt.input)
			if err == nil {
				if tt.hasError {
					t.Fatalf("Convert data success, but expected fail")
				} else {
					t.Logf("%s success convert", tt.name)
				}
			} else {
				if tt.hasError {
					t.Logf("%s: %s", tt.name, err)
				} else {
					t.Fatalf("Convert data failed: %s, but expected success", err)
				}
			}

			if mt202 != nil {
				err = mt202RepoMysql.InsertData(mt202)
				if err == nil {
					if tt.hasError {
						t.Fatalf("Insert data success, but expected fail")
					} else {
						t.Logf("%s success insert", tt.name)
					}
				} else {
					if tt.hasError {
						t.Logf("%s: %s", tt.name, err)
					} else {
						t.Fatalf("Insert data failed: %s, but expected success", err)
					}
				}
			}

		})
	}
}
