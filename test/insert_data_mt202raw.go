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
			RelatedReference:          "1CENG518990",
			ValueDateCurrencyAmount:   "181126SGD29398,98",
			OrderingInstitutionOption: 'A',
			OrderingInstitution: []string{"/F52APartyIdendifier",
				"NATXSGSGXXX"},
			AccountWithInstitutionOption: 'A',
			AccountWithInstitution: []string{"/SG67BNMA1554010237",
				"NATXSGSGXXX"},
			BeneficiaryInstitutionOption: 'A',
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
			RelatedReference:          "1CENG518990",
			ValueDateCurrencyAmount:   "181126SGD29398,98",
			OrderingInstitutionOption: 'D',
			OrderingInstitution: []string{"//SGPID123456789",
				"NEWGENS PRIVATE LIMITED COMPANY",
				"14-01 CENTRIUM SQUARE",
				"320 SERANGOON ROAD",
				"SINGAPORE 218198"},
			AccountWithInstitutionOption: 'D',
			AccountWithInstitution: []string{
				"//SGPID123456789",
				"NEWGENS PRIVATE LIMITED COMPANY",
				"14-01 CENTRIUM SQUARE",
				"320 SERANGOON ROAD",
				"SINGAPORE 218199",
			},
			BeneficiaryInstitutionOption: 'A',
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
	var dataApplicationHeaderMessageTypeLetter []*models.MT202Raw
	for i := 1; i < 4; i++ {
		temp := *data2
		temp.ApplicationHeader = data1.ApplicationHeader[0:i] + "D" + data1.ApplicationHeader[i+1:]
		dataApplicationHeaderMessageTypeLetter = append(dataApplicationHeaderMessageTypeLetter, &temp)
	}
	dataApplicationHeaderPriority := *data2
	dataApplicationHeaderPriority.ApplicationHeader = "I202SPXBSGSGXXXXJ"
	dataApplicationHeaderDeliveryMonitoringU1 := *data2
	dataApplicationHeaderDeliveryMonitoringU1.ApplicationHeader = "I202SPXBSGSGXXXXU1"
	dataApplicationHeaderDeliveryMonitoringU2 := *data2
	dataApplicationHeaderDeliveryMonitoringU2.ApplicationHeader = "I202SPXBSGSGXXXXU2"
	dataApplicationHeaderDeliveryMonitorinN1 := *data2
	dataApplicationHeaderDeliveryMonitorinN1.ApplicationHeader = "I202SPXBSGSGXXXXN1"
	dataApplicationHeaderDeliveryMonitoringN2 := *data2
	dataApplicationHeaderDeliveryMonitoringN2.ApplicationHeader = "I202SPXBSGSGXXXXN2"
	dataApplicationHeaderDeliveryObselescencePeriodU1002 := *data2
	dataApplicationHeaderDeliveryObselescencePeriodU1002.ApplicationHeader = "I202SPXBSGSGXXXXU1002"
	dataApplicationHeaderDeliveryObselescencePeriodU3003 := *data2
	dataApplicationHeaderDeliveryObselescencePeriodU3003.ApplicationHeader = "I202SPXBSGSGXXXXU3003"
	dataApplicationHeaderDeliveryObselescencePeriodU10J2 := *data2
	dataApplicationHeaderDeliveryObselescencePeriodU10J2.ApplicationHeader = "I202SPXBSGSGXXXXU10J2"

	// --------------------------------- test user header ---------------------------------
	dataUserHeaderServiceIdentifierEmpty := *data1
	dataUserHeaderServiceIdentifierEmpty.UserHeader.ServiceIdentifier = ""
	dataUserHeaderServiceIdentifier4 := *data1
	dataUserHeaderServiceIdentifier4.UserHeader.ServiceIdentifier = "EBAA"
	dataUserHeaderServiceIdentifierDigit := *data1
	dataUserHeaderServiceIdentifierDigit.UserHeader.ServiceIdentifier = "E4A"
	dataUserHeaderUniqueEndToEndID34 := *data1
	dataUserHeaderUniqueEndToEndID34.UserHeader.UniqueEndToEndID = data1.UserHeader.UniqueEndToEndID[:34]
	dataUserHeaderUniqueEndToEndIDRegex1 := *data1
	dataUserHeaderUniqueEndToEndIDRegex1.UserHeader.UniqueEndToEndID = "88af5326-3938-4efc-a7f9-984a3d595f42"
	dataUserHeaderUniqueEndToEndIDRegex2 := *data1
	dataUserHeaderUniqueEndToEndIDRegex2.UserHeader.UniqueEndToEndID = "88af5326-4938-2efc-a7f9-984a3d595f42"
	dataUserHeaderUniqueEndToEndIDRegex3 := *data1
	dataUserHeaderUniqueEndToEndIDRegex3.UserHeader.UniqueEndToEndID = "88af5326-4938-4efc-c7f9-984a3d595f42"

	// --------------------------------- test body ---------------------------------
	dataBodyTransactionReferenceNumberEmpty := *data1
	dataBodyTransactionReferenceNumberEmpty.Body.TransactionReferenceNumber = ""
	dataBodyTransactionReferenceNumber17 := *data1
	dataBodyTransactionReferenceNumber17.Body.TransactionReferenceNumber = "FRND/1062/S18XXXX"
	dataBodyRelatedReferenceEmpty := *data1
	dataBodyRelatedReferenceEmpty.Body.RelatedReference = ""
	dataBodyRelatedReference17 := *data1
	dataBodyRelatedReference17.Body.RelatedReference = "1CENG518990XXXXXX"
	dataBodyValueDateCurrencyAmount9 := *data1
	dataBodyValueDateCurrencyAmount9.Body.ValueDateCurrencyAmount = "181126SGD"
	dataBodyValueDateCurrencyAmountDateTime1 := *data1
	dataBodyValueDateCurrencyAmountDateTime1.Body.ValueDateCurrencyAmount = "311131SGD263,33"
	dataBodyValueDateCurrencyAmountDateTime2 := *data1
	dataBodyValueDateCurrencyAmountDateTime2.Body.ValueDateCurrencyAmount = "31Nov31SGD263,33"
	dataBodyValueDateCurrencyAmountCurrency1 := *data1
	dataBodyValueDateCurrencyAmountCurrency1.Body.ValueDateCurrencyAmount = "181126SG5263,33"
	dataBodyValueDateCurrencyAmountCurrency2 := *data1
	dataBodyValueDateCurrencyAmountCurrency2.Body.ValueDateCurrencyAmount = "181126SGD263,33,3"
	dataBodyOrderingInstitutionOptB := *data1
	dataBodyOrderingInstitutionOptB.Body.OrderingInstitutionOption = 'B'
	dataBodyOrderingInstitutionEmpty := *data1
	dataBodyOrderingInstitutionEmpty.Body.OrderingInstitution = []string{}
	dataBodyOrderingInstitution3 := *data1
	dataBodyOrderingInstitution3.Body.OrderingInstitution = []string{"A", "B", "C"}
	dataBodyOrderingInstitutionNoSlash := *data1
	dataBodyOrderingInstitutionNoSlash.Body.OrderingInstitution = []string{"BDAD", "NATXSGSGXXX"}
	dataBodyOrderingInstitutionManySlash := *data1
	dataBodyOrderingInstitutionManySlash.Body.OrderingInstitution = []string{"//BD/AD", "NATXSGSGXXX"}
	dataBodyOrderingInstitution1stPart2 := *data1
	dataBodyOrderingInstitution1stPart2.Body.OrderingInstitution = []string{"/JJ/BD", "NATXSGSGXXX"}
	dataBodyOrderingInstitution1stPartDigit := *data1
	dataBodyOrderingInstitution1stPartDigit.Body.OrderingInstitution = []string{"/9/BD", "NATXSGSGXXX"}
	dataBodyOrderingInstitutionOnlySlash := *data1
	dataBodyOrderingInstitutionOnlySlash.Body.OrderingInstitution = []string{"//", "NATXSGSGXXX"}
	dataBodyOrderingInstitutionIdentifierCode7 := *data1
	dataBodyOrderingInstitutionIdentifierCode7.Body.OrderingInstitution = []string{"/F52APartyIdendifier", "NATXSGS"}
	dataBodyOrderingInstitutionIdentifierCodeDigit1 := *data1
	dataBodyOrderingInstitutionIdentifierCodeDigit1.Body.OrderingInstitution = []string{"/F52APartyIdendifier", "NATXS7SGXXX"}
	dataBodyOrderingInstitutionIdentifierCodeDigit2 := *data1
	dataBodyOrderingInstitutionIdentifierCodeDigit2.Body.OrderingInstitution = []string{"/F52APartyIdendifier", "NATXSS7GXXX"}
	dataBodyAccountWithInstitutionB := *data1
	dataBodyAccountWithInstitutionB.Body.AccountWithInstitutionOption = 'B'
	dataBodyAccountWithInstitutionAddressEmpty := *data2
	dataBodyAccountWithInstitutionAddressEmpty.Body.AccountWithInstitution = []string{"//SGPID123456789", ""}

	// --------------------------------- insert basic header ---------------------------------
	type testStruct struct {
		name     string
		input    *models.MT202Raw
		hasError bool
	}
	var tests = []*testStruct{
		{"data1", data1, false},
		{"data2", data2, false},
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
	tests = append(tests, func() []*testStruct {
		var data []*testStruct
		for i, v := range dataApplicationHeaderMessageTypeLetter {
			data = append(data, &testStruct{fmt.Sprintf("dataApplicationHeaderMessageTypeLetter%d", i), v, true})
		}
		return data
	}()...)
	tests = append(tests, []*testStruct{
		{"dataApplicationHeaderPriority", &dataApplicationHeaderPriority, true},
		{"dataApplicationHeaderDeliveryMonitoringU1", &dataApplicationHeaderDeliveryMonitoringU1, false},
		{"dataApplicationHeaderDeliveryMonitoringU2", &dataApplicationHeaderDeliveryMonitoringU2, true},
		{"dataApplicationHeaderDeliveryMonitorinN1", &dataApplicationHeaderDeliveryMonitorinN1, true},
		{"dataApplicationHeaderDeliveryMonitoringN2", &dataApplicationHeaderDeliveryMonitoringN2, false},
		{"dataApplicationHeaderDeliveryObselescencePeriodU1002", &dataApplicationHeaderDeliveryObselescencePeriodU1002, true},
		{"dataApplicationHeaderDeliveryObselescencePeriodU3003", &dataApplicationHeaderDeliveryObselescencePeriodU3003, false},
		{"dataApplicationHeaderDeliveryObselescencePeriodU10J2", &dataApplicationHeaderDeliveryObselescencePeriodU10J2, true},
	}...)

	// --------------------------------- insert user header ---------------------------------
	tests = append(tests, []*testStruct{
		{"dataUserHeaderServiceIdentifierEmpty", &dataUserHeaderServiceIdentifierEmpty, false},
		{"dataUserHeaderServiceIdentifier4", &dataUserHeaderServiceIdentifier4, true},
		{"dataUserHeaderServiceIdentifierDigit", &dataUserHeaderServiceIdentifierDigit, true},
		{"dataUserHeaderUniqueEndToEndID34", &dataUserHeaderUniqueEndToEndID34, true},
		{"dataUserHeaderUniqueEndToEndIDRegex1", &dataUserHeaderUniqueEndToEndIDRegex1, false},
		{"dataUserHeaderUniqueEndToEndIDRegex2", &dataUserHeaderUniqueEndToEndIDRegex2, true},
		{"dataUserHeaderUniqueEndToEndIDRegex3", &dataUserHeaderUniqueEndToEndIDRegex3, true},
	}...)

	// --------------------------------- insert body ---------------------------------
	tests = append(tests, []*testStruct{
		{"dataBodyTransactionReferenceNumberEmpty", &dataBodyTransactionReferenceNumberEmpty, true},
		{"dataBodyTransactionReferenceNumber17", &dataBodyTransactionReferenceNumber17, true},
		{"dataBodyRelatedReferenceEmpty", &dataBodyRelatedReferenceEmpty, true},
		{"dataBodyRelatedReference17", &dataBodyRelatedReference17, true},
		{"dataBodyValueDateCurrencyAmount9", &dataBodyValueDateCurrencyAmount9, true},
		{"dataBodyValueDateCurrencyAmountDateTime1", &dataBodyValueDateCurrencyAmountDateTime1, true},
		{"dataBodyValueDateCurrencyAmountDateTime2", &dataBodyValueDateCurrencyAmountDateTime2, true},
		{"dataBodyValueDateCurrencyAmountCurrency1", &dataBodyValueDateCurrencyAmountCurrency1, true},
		{"dataBodyValueDateCurrencyAmountCurrency2", &dataBodyValueDateCurrencyAmountCurrency2, true},
		{"dataBodyOrderingInstitutionOptB", &dataBodyOrderingInstitutionOptB, true},
		{"dataBodyOrderingInstitutionEmpty", &dataBodyOrderingInstitutionEmpty, true},
		{"dataBodyOrderingInstitution3", &dataBodyOrderingInstitution3, true},
		{"dataBodyOrderingInstitutionNoSlash", &dataBodyOrderingInstitutionNoSlash, true},
		{"dataBodyOrderingInstitutionManySlash", &dataBodyOrderingInstitutionManySlash, true},
		{"dataBodyOrderingInstitution1stPart2", &dataBodyOrderingInstitution1stPart2, true},
		{"dataBodyOrderingInstitution1stPartDigit", &dataBodyOrderingInstitution1stPartDigit, true},
		{"dataBodyOrderingInstitutionOnlySlash", &dataBodyOrderingInstitutionOnlySlash, true},
		{"dataBodyOrderingInstitutionIdentifierCode7", &dataBodyOrderingInstitutionIdentifierCode7, true},
		{"dataBodyOrderingInstitutionIdentifierCodeDigit1", &dataBodyOrderingInstitutionIdentifierCodeDigit1, true},
		{"dataBodyOrderingInstitutionIdentifierCodeDigit2", &dataBodyOrderingInstitutionIdentifierCodeDigit2, false},
		{"dataBodyAccountWithInstitutionB", &dataBodyAccountWithInstitutionB, false},
		{"dataBodyAccountWithInstitutionAddressEmpty", &dataBodyAccountWithInstitutionAddressEmpty, true},
	}...)

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.hasError {
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
