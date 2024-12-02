package utils

import (
	"os"
	"time"
)

// como el servidor tieme gmt+1 y queremos gmt-4 bolivia
// restamos 5 horas, podemos recibir por parametro la cantidad
// de horas a restar segun la zona horaria
func ToTZ(fecha time.Time) time.Time {
	e := os.Getenv("PARSE_TIME")
	if e == "1" {
		return fecha.Add(-5 * time.Hour)
	}
	return fecha
}

func ToTZNil(fecha *time.Time) *time.Time {
	if fecha == nil {
		return nil
	}
	f := ToTZ(*fecha)
	return &f
}
