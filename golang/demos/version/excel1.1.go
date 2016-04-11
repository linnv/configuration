package main

import (
	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2"
	"time"
)

type category struct {
	Id       int    `bson:"_id"`    //pk
	CateId   int    `bson:"CateId"` //分类的具体Id
	Name     string `bson:"Name"`
	ParentId int    `bson:"ParentId"`

	CreateTime int64 `bson:"CreateTime" json:"-"`
	UpdateTime int64 `bson:"UpdateTime" json:"-"`
	IsDeleted  bool  `bson:"IsDeleted"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("ssp_jialin").C("Category")

	excelFileName := "/Users/Jialin/Downloads/siteCategory1.0.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	var _id = 0
	for _, sheet := range xlFile.Sheets {

		var p_id int
		for i, _ := range sheet.Rows {
			if i == 0 {
				continue
			}
			t := time.Now().Unix()
			pcid, _ := sheet.Rows[i].Cells[0].Int()
			pname := sheet.Rows[i].Cells[1].String()
			cid, _ := sheet.Rows[i].Cells[2].Int()
			cname := sheet.Rows[i].Cells[3].String()

			//parent category

			println(i)
			if i < len(sheet.Rows)-6 && i > 0 {

				pcidn, _ := sheet.Rows[i-1].Cells[0].Int()
				// pcidn, _ := sheet.Rows[i+1].Cells[0].Int()
				if pcid != pcidn {
					_id++
					p_id = _id
					tmpc := &category{
						Id:         _id,
						CateId:     pcid,
						Name:       pname,
						CreateTime: t,
						UpdateTime: t,
						IsDeleted:  false,
					}

					err = c.Insert(tmpc)
					if err != nil {
						panic(err)
					}
				}
			}

			if i == len(sheet.Rows)-6 {
				println(pname)
				_id++
				p_id = _id
				t := time.Now().Unix()
				tmpc := &category{
					Id:         _id,
					CateId:     24,
					Name:       "网络购物",
					CreateTime: t,
					UpdateTime: t,
					IsDeleted:  false,
				}
				err = c.Insert(tmpc)
				if err != nil {
					panic(err)
				}
				// continue
			}

			//child category

			_id++

			tmpc := &category{
				Id:         _id,
				CateId:     cid,
				Name:       cname,
				CreateTime: t,
				UpdateTime: t,
				IsDeleted:  false,
				ParentId:   p_id,
			}
			err = c.Insert(tmpc)
			// err = c.Insert(bson.M{"CateId": cid, "Name": cname, "CreateTime": t, "UpdateTime": t, "IsDeleted": false, "ParentId": pid})
			if err != nil {
				panic(err)
			}

		}

	}
}
