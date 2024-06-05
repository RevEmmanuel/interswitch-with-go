package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Configuration struct {
	AuthURL                      string
	AuthHeader                   string
	ValidatePurchaseRecurrentUrl string
	AUTHTOKEN                    string
	TOKENIZETRANSACTIONURL       string
	AUTHDATA                     string
	TRANSACTIONREF               string
}

var Config Configuration

func LoadConfig() {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	Config = Configuration{
		AuthURL:                      config["authURL"],
		AuthHeader:                   config["authHeader"],
		ValidatePurchaseRecurrentUrl: config["https://qa.interswitchng.com/api/v2/purchases/recurrents"],
		AUTHTOKEN:                    config["Bearer eyJhbGciOiJSUzI1NiJ9.eyJtZXJjaGFudF9jb2RlIjoiTVg2MDcyIiwicmVxdWVzdG9yX2lkIjoiMTIzODA4NTk1MDMiLCJpbmNvZ25pdG9fcmVxdWVzdG9yX2lkIjoiMTIzODA4NTk1MDMiLCJwYXlhYmxlX2lkIjoiMzM1OTciLCJjbGllbnRfZGVzY3JpcHRpb24iOm51bGwsImNsaWVudF9pZCI6IklLSUFCMjNBNEUyNzU2NjA1QzFBQkMzM0NFM0MyODdFMjcyNjdGNjYwRDYxIiwiYXVkIjpbImFwaS1nYXRld2F5IiwiYXJiaXRlciIsImNhZXNhciIsImhpbXMtcG9ydGxldCIsImluY29nbml0byIsImlzdy1jb2xsZWN0aW9ucyIsImlzdy1jb3JlIiwiaXN3LWluc3RpdHV0aW9uIiwiaXN3LWxlbmRpbmctc2VydmljZSIsImlzdy1wYXBlIiwiaXN3LXBhcHJzIiwiaXN3LXBhcHNzIiwiaXN3LXBheW1lbnRnYXRld2F5IiwiaXN3LXBvc3Qtb2ZmaWNlIiwia3ljLXNlcnZpY2UiLCJwYXNzcG9ydCIsInBvc3RpbGlvbi1hcGkiLCJwcm9qZWN0LXgtY29uc3VtZXIiLCJwcm9qZWN0LXgtbWVyY2hhbnQiLCJxdC1zZXJ2aWNlIiwicXVpY2t0ZWxsZXItZXRsci1yZXF1ZXJ5IiwicmVjdXJyZW50LWJpbGxpbmctYXBpIiwidHJhbnNmZXItc2VydmljZS1hZG1pbiIsInRyYW5zZmVyLXNlcnZpY2UtY29yZSIsInZhdWx0Iiwidm91Y2hlci1hcGkiLCJ3YWxsZXQiLCJ3ZWJwYXktcG9ydGxldCJdLCJjbGllbnRfYXV0aG9yaXphdGlvbl9kb21haW4iOiJNWDYwNzIiLCJzY29wZSI6WyJjbGllbnRzIiwicHJvZmlsZSJdLCJhcGlfcmVzb3VyY2VzIjpbInJpZC1QT1NUL2FwaS92MS9wdXJjaGFzZXMiLCJyaWQtUE9TVC9hcGkvdjEvcHVyY2hhc2VzLyoqIiwicmlkLVBVVC9hcGkvdjEvcHVyY2hhc2VzIiwicmlkLVBVVC9hcGkvdjEvcHVyY2hhc2VzLyoqIiwicmlkLUdFVC9hcGkvdjEvcHVyY2hhc2VzIiwicmlkLUdFVC9hcGkvdjEvcHVyY2hhc2VzLyoqIiwicmlkLURFTEVURS9hcGkvdjEvcHVyY2hhc2VzIiwicmlkLURFTEVURS9hcGkvdjEvcHVyY2hhc2VzLyoqIiwicmlkLVBPU1QvYXBpL3YyL3B1cmNoYXNlcyIsInJpZC1QT1NUL2FwaS92Mi9wdXJjaGFzZXMvKioiLCJyaWQtUFVUL2FwaS92Mi9wdXJjaGFzZXMiLCJyaWQtUFVUL2FwaS92Mi9wdXJjaGFzZXMvKioiLCJyaWQtR0VUL2FwaS92Mi9wdXJjaGFzZXMiLCJyaWQtR0VUL2FwaS92Mi9wdXJjaGFzZXMvKioiLCJyaWQtREVMRVRFL2FwaS92Mi9wdXJjaGFzZXMiLCJyaWQtREVMRVRFL2FwaS92Mi9wdXJjaGFzZXMvKioiLCJyaWQtUE9TVC9hcGkvdjMvcHVyY2hhc2VzIiwicmlkLVBPU1QvYXBpL3YzL3B1cmNoYXNlcy8qKiIsInJpZC1QVVQvYXBpL3YzL3B1cmNoYXNlcyIsInJpZC1QVVQvYXBpL3YzL3B1cmNoYXNlcy8qKiIsInJpZC1HRVQvYXBpL3YzL3B1cmNoYXNlcyIsInJpZC1HRVQvYXBpL3YzL3B1cmNoYXNlcy8qKiIsInJpZC1ERUxFVEUvYXBpL3YzL3B1cmNoYXNlcyIsInJpZC1ERUxFVEUvYXBpL3YzL3B1cmNoYXNlcy8qKiJdLCJleHAiOjMyNzg1MTAzNjUsImNsaWVudF9uYW1lIjoiUjdqSmhyRWd5TCIsImNsaWVudF9sb2dvIjpudWxsLCJqdGkiOiI5MGE5NWYzYy0zOGVlLTQ5YmYtYjRlMS0zNjBiZjRhMjZjNmMifQ.KVBRAXnQxN8G8ssK5xBVq8TxvHAj9IuunqpRT3q1Fc2c5eWkEGOom7jRUOMEE6rIcqh7Sqe5fi-PGpfj8wTKFl9kLNRb7w4a8lXWYkd_iudHpVutY16cLO2UcJyLwS0lgCeEF7IJD2_K1-d2i6MdrYNqkYG1VHUKbNfp_i1z1jcw5AEkOwKG2R32MsoKd70ZuGvFyH577Qno-5YbmFEk7b02aJe98bNneFw6hAC1UWhUZWyYkAHiFS67PGGxh46ZN8GXcg-JG_VlZeBXnnhMudek9wuxFmDh5d7BmvAi1lmf2Bs8lAu7meoxbPbtKoB-qVDOzLuF6QJcnYggRzz9Wg"],
		TOKENIZETRANSACTIONURL:       config["https://qa.interswitchng.com/api/v2/purchases/validations/recurrents"],
		AUTHDATA:                     config["Uk4wiexIGkIQtjIXOuO4O7kE+3VTP5fxBY+O1YeQI5uuDFDbdojwahSNeNG/ta0tZW4zuJvynZQIqeaDU/jOMEDEhRsf/++EYK8ApWlg0MHWO1tE6wURnPvrGMU4Mpf7lg8UB7mWbkmssO/tpM95KUAwQ5MVSxOBB3O7rFGYGoqgPK/wHLqMSUuYNYpLbKJsWdvIMO5fiDoUoMqf+FVIzpL18tWvmEw6E2gev6o46CwBEdUBXS52W8e8GqoBdxaa/KSSLeYge88Um8d88VVJH62fYycdWDj6Gubg0uUWNRm0xl8GyMcU9RpGv1She8t53c8VDcUFnEg19/WH1CN+1w=="],
		TRANSACTIONREF:               config["x14NPCcOnfkOipa"],
	}
}
