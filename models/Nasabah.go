package models

import (
	"fmt"
	"go_rest/config"
)

func GetAll(Nasabah *[]Nasabah) (err error) {
	if err := config.DB.Find(Nasabah).Error; err != nil {
		return err
	}
	return nil
}
func Create(Nasabah *Nasabah) (err error) {
	if err = config.DB.Create(Nasabah).Error; err != nil {
		return err
	}
	return nil
}

func Detail(Nasabah *Nasabah, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(Nasabah).Error; err != nil {
		return err
	}
	return nil
}

func Update(Nasabah *Nasabah, id string) (err error) {
	fmt.Println(Nasabah)
	config.DB.Save(Nasabah)
	return nil
}

func Delete(Nasabah *Nasabah, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(Nasabah)
	return nil
}
