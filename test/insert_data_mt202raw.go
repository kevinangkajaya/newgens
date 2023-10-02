package test

import (
	"fmt"
	"newgens/models"
	"newgens/repository/mysql"
	"newgens/src"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func data1Raw() (*models.MT202Raw, error) {
	rawData, err := src.ReadLines("files/202MEP_52A_57A_58A.txt")
	if err != nil {
		return nil, err
	}
	data := models.MT202Raw{
		BasicHeader:       "F01BNINSGSGAXXX7535406549",
		ApplicationHeader: "I202UOVBSGSGXXXXN",
		UserHeader: models.MT202RawUserHeader{
			ServiceIdentifier: "MEP",
			// MessageUserReference: "1000000296",
			UniqueEndToEndID: "88af5326-4938-4efc-a7f9-984a3d595f42",
		},
		Body: models.MT202RawBody{TransactionReferenceNumber: "FRND/1062/S18",
			RelatedReference:        "1CENG518990",
			ValueDateCurrencyAmount: "181126SGD29398,98",
			OrderingInstitution: []string{"/F52APartyIdendifier",
				"NATXSGSGXXX"},
			AccountWithInstitution: []string{"/SG67BNMA1554010237",
				"NATXSGSGXXX"},
			BeneficiaryInstitution: []string{"/DE66512305000018500101",
				"NATXSGSGXXX"},
		},
		RawData: strings.Join(rawData, "\n"),
	}

	return &data, nil
}

func data2Raw() (*models.MT202Raw, error) {
	rawData, err := src.ReadLines("files/202MEP_52D_57D_58A.txt")
	if err != nil {
		return nil, err
	}
	data := models.MT202Raw{
		BasicHeader:       "F01SPXBPHMMAXXX7535406549",
		ApplicationHeader: "I202SPXBSGSGXXXXN",
		UserHeader: models.MT202RawUserHeader{
			ServiceIdentifier: "MEP",
			// MessageUserReference: "MT202-00042021-3",
			UniqueEndToEndID: "18a7fbcc-16c0-4fc1-abde-8b653e1d03ba",
		},
		Body: models.MT202RawBody{TransactionReferenceNumber: "MT202-00042021-3",
			RelatedReference:        "1CENG518990",
			ValueDateCurrencyAmount: "181126SGD29398,98",
			OrderingInstitution: []string{"//SGPID123456789",
				"NEWGENS PRIVATE LIMITED COMPANY",
				"14-01 CENTRIUM SQUARE",
				"320 SERANGOON ROAD",
				"SINGAPORE 218198"},
			AccountWithInstitution: []string{
				"//SGPID123456789",
				"NEWGENS PRIVATE LIMITED COMPANY",
				"14-01 CENTRIUM SQUARE",
				"320 SERANGOON ROAD",
				"SINGAPORE 218199",
			},
			BeneficiaryInstitution: []string{
				"/DE66512305000018500101",
				"NATXSGSGXXX"},
		},
		RawData: strings.Join(rawData, "\n"),
	}

	return &data, nil
}

func TestInsertDataMt202Raw(t *testing.T) {
	db := src.ConnectMysql()
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	data1, err := data1Raw()
	if err != nil {
		t.Fatalf(err.Error())
	}
	data2, err := data2Raw()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// --------------------------------- test basic header ---------------------------------
	dataBasicHeaderEmpty := *data1
	dataBasicHeaderEmpty.BasicHeader = ""
	dataBasicHeader24 := *data1
	dataBasicHeader24.BasicHeader = "F01BNINSGSGAXXX753540654"
	dataBasicHeaderAppIdDigit := *data1
	dataBasicHeaderAppIdDigit.BasicHeader = "501BNINSGSGAXXX7535406549"
	dataBasicHeaderServiceIdLetter1 := *data1
	dataBasicHeaderServiceIdLetter1.BasicHeader = "FG1BNINSGSGAXXX7535406549"
	dataBasicHeaderServiceIdLetter2 := *data1
	dataBasicHeaderServiceIdLetter2.BasicHeader = "F0GBNINSGSGAXXX7535406549"
	var dataBasicHeaderSessionNumberLetter []*models.MT202Raw
	for i := 15; i < 19; i++ {
		temp := *data1
		temp.BasicHeader = data1.BasicHeader[0:i] + "G" + data1.BasicHeader[i+1:]
		dataBasicHeaderSessionNumberLetter = append(dataBasicHeaderSessionNumberLetter, &temp)
	}
	var dataBasicHeaderSequenceNumberLetter []*models.MT202Raw
	for i := 19; i < 25; i++ {
		temp := *data1
		temp.BasicHeader = data1.BasicHeader[0:i] + "G" + data1.BasicHeader[i+1:]
		dataBasicHeaderSequenceNumberLetter = append(dataBasicHeaderSequenceNumberLetter, &temp)
	}

	// --------------------------------- test application header ---------------------------------
	dataApplicationHeaderEmpty := *data2
	dataApplicationHeaderEmpty.ApplicationHeader = ""
	dataApplicationHeader15 := *data2
	dataApplicationHeader15.ApplicationHeader = "I20SPXBSGSGXXXXN"
	dataApplicationHeaderOutput := *data2
	dataApplicationHeaderOutput.ApplicationHeader = "O2021345160418SOGEFRPPAXXX00897254971604181345N"
	dataApplicationHeaderOutput45 := *data2
	dataApplicationHeaderOutput45.ApplicationHeader = "O21345160418SOGEFRPPAXXX00897254971604181345N"
	dataApplicationHeaderX := *data2
	dataApplicationHeaderX.ApplicationHeader = "X2021345160418SOGEFRPPAXXX00897254971604181345N"

	// --------------------------------- insert basic header ---------------------------------
	type testStruct struct {
		name     string
		input    *models.MT202Raw
		hasError bool
	}
	var tests = []*testStruct{
		// {"data1", data1, false},
		// {"data2", data2, false},
		{"dataBasicHeaderEmpty", &dataBasicHeaderEmpty, true},
		{"dataBasicHeader24", &dataBasicHeader24, true},
		{"dataBasicHeaderAppIdDigit", &dataBasicHeaderAppIdDigit, true},
		{"dataBasicHeaderServiceIdLetter1", &dataBasicHeaderServiceIdLetter1, true},
		{"dataBasicHeaderServiceIdLetter2", &dataBasicHeaderServiceIdLetter2, true},
	}
	tests = append(tests, func() []*testStruct {
		var data []*testStruct
		for i, v := range dataBasicHeaderSessionNumberLetter {
			data = append(data, &testStruct{fmt.Sprintf("dataBasicHeaderSessionNumberLetter%d", i), v, true})
		}
		return data
	}()...)
	tests = append(tests, func() []*testStruct {
		var data []*testStruct
		for i, v := range dataBasicHeaderSequenceNumberLetter {
			data = append(data, &testStruct{fmt.Sprintf("dataBasicHeaderSequenceNumberLetter%d", i), v, true})
		}
		return data
	}()...)

	// --------------------------------- insert application header ---------------------------------
	tests = append(tests, []*testStruct{
		{"dataApplicationHeaderEmpty", &dataApplicationHeaderEmpty, true},
		{"dataApplicationHeader15", &dataApplicationHeader15, true},
		{"dataApplicationHeaderOutput", &dataApplicationHeaderOutput, false},
		{"dataApplicationHeaderOutput45", &dataApplicationHeaderOutput45, true},
		{"dataApplicationHeaderX", &dataApplicationHeaderX, true},
	}...)

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
