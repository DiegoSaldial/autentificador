package utils

// como el servidor tieme gmt+1 y queremos gmt-4 bolivia
// restamos 5 horas, podemos recibir por parametro la cantidad
// de horas a restar segun la zona horaria

// YA NO ES NECESARIO SE GUARDA EN UTC Y EL CLIENTE LO CONVIERTE A SU HORA LOCAL
// AUQUE HAY UN BREVE RPOBLEMITA QUE SE ENVIA CON OFFSET -04:00
// SI BIEN EN EL CLIENTE SE LO IGNORA, PODRIA TENER PROBLEMAS FUTUROS
/* func ToTZ(fecha time.Time) time.Time {
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
*/
