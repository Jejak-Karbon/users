package seeder

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
	"log"
)

func provinceTableSeeder(conn *gorm.DB) {

	var provinces = []model.Province{
		{ID:"11",Name:"ACEH"},
{ID:"12",Name:"SUMATERA UTARA"},
{ID:"13",Name:"SUMATERA BARAT"},
{ID:"14",Name:"RIAU"},
{ID:"15",Name:"JAMBI"},
{ID:"16",Name:"SUMATERA SELATAN"},
{ID:"17",Name:"BENGKULU"},
{ID:"18",Name:"LAMPUNG"},
{ID:"19",Name:"KEPULAUAN BANGKA BELITUNG"},
{ID:"21",Name:"KEPULAUAN RIAU"},
{ID:"31",Name:"DKI JAKARTA"},
{ID:"32",Name:"JAWA BARAT"},
{ID:"33",Name:"JAWA TENGAH"},
{ID:"34",Name:"DI YOGYAKARTA"},
{ID:"35",Name:"JAWA TIMUR"},
{ID:"36",Name:"BANTEN"},
{ID:"51",Name:"BALI"},
{ID:"52",Name:"NUSA TENGGARA BARAT"},
{ID:"53",Name:"NUSA TENGGARA TIMUR"},
{ID:"61",Name:"KALIMANTAN BARAT"},
{ID:"62",Name:"KALIMANTAN TENGAH"},
{ID:"63",Name:"KALIMANTAN SELATAN"},
{ID:"64",Name:"KALIMANTAN TIMUR"},
{ID:"65",Name:"KALIMANTAN UTARA"},
{ID:"71",Name:"SULAWESI UTARA"},
{ID:"72",Name:"SULAWESI TENGAH"},
{ID:"73",Name:"SULAWESI SELATAN"},
{ID:"74",Name:"SULAWESI TENGGARA"},
{ID:"75",Name:"GORONTALO"},
{ID:"76",Name:"SULAWESI BARAT"},
{ID:"81",Name:"MALUKU"},
{ID:"82",Name:"MALUKU UTARA"},
{ID:"91",Name:"PAPUA BARAT"},
{ID:"94",Name:"PAPUA"},
	}

	if err := conn.Create(&provinces).Error; err != nil {
		log.Printf("cannot seed data provinces, with error %v\n", err)
	}
	log.Println("success seed data provinces")
}
