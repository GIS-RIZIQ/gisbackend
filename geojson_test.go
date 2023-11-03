package gisbackend

import (
	"context"
	"fmt"
	"testing"

	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("riziq", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "duar",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGOSTRING", "geo", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Point",
		Name:   "titik",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGOSTRING", "geo", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGOSTRING", "geo")
	pass, _ := pasproj.HashPass("riziq")
	data := RegisterStruct{
		Username: "riziq",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}
