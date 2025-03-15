package application

import (
	
	"act2/src/esp32Temperature/domain"
)

type PostSensor struct {
	db domain.ISensor 

}

func NewPostSensor(db domain.ISensor)*PostSensor{
	return &PostSensor{db:db	}
} 

func (uc *PostSensor) Execute(menssage string) error{
	sensor := domain.Sensor{
		Menssage: menssage,

	}

	// Guardar 
	err := uc.db.Save(sensor.Menssage)
	if err != nil {
		return err
	}

	

	return nil
}