package main

import (
	"fmt"
	"newgens/models"
	"newgens/repository/mysql"
	"newgens/src"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
)

func data1() (*models.MT202, error) {
	F32A_ValueDate, err := time.Parse("060102", "181126")
	if err != nil {
		return nil, err
	}
	F32A_Amount, err := decimal.NewFromString(src.ReplaceCommaWithDot("29398,98"))
	if err != nil {
		return nil, err
	}
	rawData, err := src.ReadLines("files/202MEP_52A_57A_58A.txt")
	if err != nil {
		return nil, err
	}
	data := models.MT202{
		SenderBIC:      "BNINSGSGAXXX",
		ReceiverBIC:    "UOVBSGSGXXXX",
		Direction:      "I",
		MTType:         "MEP",
		UETR:           "88af5326-4938-4efc-a7f9-984a3d595f42",
		F20:            "FRND/1062/S18",
		F21:            "1CENG518990",
		F32A_ValueDate: F32A_ValueDate,
		F32A_Currency:  "SGD",
		F32A_Amount:    F32A_Amount,
		F52a:           "/F52APartyIdendifier.NATXSGSGXXX",
		F57a:           "SG67BNMA1554010237.NATXSGSGXXX",
		F58a:           "/DE66512305000018500101.NATXSGSGXXX",
		RawData:        strings.Join(rawData, "\n"),
	}

	return &data, nil
}

func data2() (*models.MT202, error) {
	F32A_ValueDate, err := time.Parse("060102", "181126")
	if err != nil {
		return nil, err
	}
	F32A_Amount, err := decimal.NewFromString(src.ReplaceCommaWithDot("29398,98"))
	if err != nil {
		return nil, err
	}
	rawData, err := src.ReadLines("files/202MEP_52D_57D_58A.txt")
	if err != nil {
		return nil, err
	}
	data := models.MT202{
		SenderBIC:      "SPXBPHMMAXXX",
		ReceiverBIC:    "SPXBSGSGXXXX",
		Direction:      "I",
		MTType:         "MEP",
		UETR:           "18a7fbcc-16c0-4fc1-abde-8b653e1d03ba",
		F20:            "MT202-00042021-3",
		F21:            "1CENG518990",
		F32A_ValueDate: F32A_ValueDate,
		F32A_Currency:  "SGD",
		F32A_Amount:    F32A_Amount,
		F52a:           "//SGPID123456789.NEWGENS PRIVATE LIMITED COMPANY.14-01 CENTRIUM SQUARE.320 SERANGOON ROAD.SINGAPORE 218198",
		F57a:           "//SGPID123456789.NEWGENS PRIVATE LIMITED COMPANY.14-01 CENTRIUM SQUARE.320 SERANGOON ROAD.SINGAPORE 218199",
		F58a:           "/DE66512305000018500101.NATXSGSGXXX",
		RawData:        strings.Join(rawData, "\n"),
	}

	return &data, nil
}

func TestInsertDataMt202(t *testing.T) {
	db := src.ConnectMysql()
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	data1, err := data1()
	if err != nil {
		t.Fatalf(err.Error())
	}
	data2, err := data2()
	if err != nil {
		t.Fatalf(err.Error())
	}

	dataSenderBICEmpty := *data1
	dataSenderBICEmpty.SenderBIC = ""
	dataSenderBIC5 := *data1
	dataSenderBIC5.SenderBIC = "BNINS"
	dataSenderBIC11 := *data1
	dataSenderBIC11.SenderBIC = "BNINSGSGAXX"
	dataReceiverBIC8thCharacter := *data2
	dataReceiverBIC8thCharacter.ReceiverBIC = "SPXBSGSG4XXX"
	dataReceiverBIC7 := *data2
	dataReceiverBIC7.ReceiverBIC = "SPXBSGS"
	dataReceiverBIC13 := *data2
	dataReceiverBIC13.ReceiverBIC = "SPXBSGSGXXXXX"
	dataDirectionXX := *data1
	dataDirectionXX.Direction = "XX"
	dataDirectionO := *data1
	dataDirectionO.Direction = "O"
	dataMTType5 := *data1
	dataMTType5.MTType = "MEPXX"
	dataUETR35 := *data1
	dataUETR35.UETR = dataUETR35.UETR[:35]
	dataF20Empty := *data1
	dataF20Empty.F20 = ""
	dataF20X17 := *data1
	dataF20X17.F20 = "FRND/1062/S18ABCD"
	dataF21Empty := *data1
	dataF21Empty.F21 = ""
	dataF21X18 := *data1
	dataF21X18.F21 = "1CENG518990ABCDEFG"
	dataF32A_ValueDateZero := *data1
	dataF32A_ValueDateZero.F32A_ValueDate = time.Time{}
	dataF32A_Currency2 := *data1
	dataF32A_Currency2.F32A_Currency = "SG"
	dataF32A_Currency4 := *data1
	dataF32A_Currency4.F32A_Currency = "SGDD"
	dataF58aEmpty := *data1
	dataF58aEmpty.F58a = ""
	dataRawDataEmpty := *data1
	dataRawDataEmpty.RawData = ""

	var tests = []struct {
		name     string
		input    *models.MT202
		hasError bool
	}{
		// {"data1", data1, false},
		// {"data2", data2, false},
		{"dataSenderBICEmpty", &dataSenderBICEmpty, true},
		{"dataSenderBIC5", &dataSenderBIC5, true},
		{"dataSenderBIC11", &dataSenderBIC11, true},
		{"dataReceiverBIC8thCharacter", &dataReceiverBIC8thCharacter, true},
		{"dataReceiverBIC7", &dataReceiverBIC7, true},
		{"dataReceiverBIC13", &dataReceiverBIC13, true},
		{"dataDirectionXX", &dataDirectionXX, true},
		// {"dataDirectionO", &dataDirectionO, false},
		{"dataMTType5", &dataMTType5, true},
		{"dataUETR35", &dataUETR35, true},
		{"dataF20Empty", &dataF20Empty, true},
		{"dataF20X17", &dataF20X17, true},
		{"dataF21Empty", &dataF21Empty, true},
		{"dataF21X18", &dataF21X18, true},
		{"dataF32A_ValueDateZero", &dataF32A_ValueDateZero, true},
		{"dataF32A_Currency2", &dataF32A_Currency2, true},
		{"dataF32A_Currency4", &dataF32A_Currency4, true},
		{"dataF58aEmpty", &dataF58aEmpty, true},
		{"dataRawDataEmpty", &dataRawDataEmpty, true},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mt202RepoMysql.InsertData(tt.input)
			if err == nil {
				if tt.hasError {
					t.Fatalf("Insert data success, but expected fail")
				}
			} else {
				if !tt.hasError {
					t.Fatalf("Insert data failed: %s, but expected success", err)
				}
			}

		})
	}
}

func TestGetDataMt202(t *testing.T) {
	db := src.ConnectMysql()
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	mt202, err := mt202RepoMysql.GetData()
	if err != nil {
		t.Fatalf("Get data failed: %s", err)
	}
	for _, v := range mt202 {
		fmt.Println(v)
	}
}
