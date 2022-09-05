package numerosATexto

import (
	"fmt"
	"math"
)

type Monto struct {
	Valor               float64
	TextoMonedaSingular string
	TextoMonedaPlural   string
}

func unidades(numero int) string {
	switch numero {
	case 1:
		return "UN"
	case 2:
		return "DOS"
	case 3:
		return "TRES"
	case 4:
		return "CUATRO"
	case 5:
		return "CINCO"
	case 6:
		return "SEIS"
	case 7:
		return "SIETE"
	case 8:
		return "OCHO"
	case 9:
		return "NUEVE"
	default:
		return ""
	}
}

func decenas(numero int) string {
	decena := math.Floor(float64(numero) / 10)
	unidad := numero - (int(decena) * 10)

	switch decena {
	case 1:
		switch unidad {
		case 0:
			return "DIEZ"
		case 1:
			return "ONCE"
		case 2:
			return "DOCE"
		case 3:
			return "TRECE"
		case 4:
			return "CATORCE"
		case 5:
			return "QUINCE"
		default:
			return "DIECI" + unidades(unidad)
		}
	case 2:
		switch unidad {
		case 0:
			return "VEINTE"
		case 1:
			return "VEINTIUN"
		default:
			return "VEINTI" + unidades(unidad)
		}
	case 3:
		return decenaConUnidad("TREINTA", unidad)
	case 4:
		return decenaConUnidad("CUARENTA", unidad)
	case 5:
		return decenaConUnidad("CINCUENTA", unidad)
	case 6:
		return decenaConUnidad("SESENTA", unidad)
	case 7:
		return decenaConUnidad("SETENTA", unidad)
	case 8:
		return decenaConUnidad("OCHENTA", unidad)
	case 9:
		return decenaConUnidad("NOVENTA", unidad)
	default:
		return unidades(unidad)
	}
}

func decenaConUnidad(textoDecena string, numero int) string {
	if numero > 0 && numero <= 9 {
		return fmt.Sprintf("%s Y %v", textoDecena, unidades(numero))
	} else {
		return textoDecena
	}
}

func centenas(numero int) string {
	centenas := math.Floor(float64(numero) / 100)
	residuo := numero - (int(centenas) * 100)

	switch centenas {
	case 1:
		if residuo > 1 {
			return fmt.Sprintf("%s %v", "CIENTO", decenas(residuo))
		} else {
			return "CIEN"
		}
	case 2:
		return fmt.Sprintf("%s %v", "DOSCIENTOS", decenas(residuo))
	case 3:
		return fmt.Sprintf("%s %v", "TRESCIENTOS", decenas(residuo))
	case 4:
		return fmt.Sprintf("%s %v", "CUATROCIENTOS", decenas(residuo))
	case 5:
		return fmt.Sprintf("%s %v", "QUINIENTOS", decenas(residuo))
	case 6:
		return fmt.Sprintf("%s %v", "SEISCIENTOS", decenas(residuo))
	case 7:
		return fmt.Sprintf("%s %v", "SETECIENTOS", decenas(residuo))
	case 8:
		return fmt.Sprintf("%s %v", "OCHOCIENTOS", decenas(residuo))
	case 9:
		return fmt.Sprintf("%s %v", "NOVECIENTOS", decenas(residuo))
	default:
		return decenas(residuo)
	}
}

func miles(numero int) string {
	miles := math.Floor(float64(numero) / 1000)
	residuo := numero - (int(miles) * 1000)

	textoMiles := ""
	textoCentenas := centenas(residuo)

	if miles > 1 {
		textoMiles = fmt.Sprintf("%s %s", centenas(int(miles)), "MIL")
	} else if miles == 1 {
		textoMiles = "MIL"
	} else {
		return textoCentenas
	}

	return textoMiles + " " + textoCentenas
}

func millones(numero int) string {
	millones := math.Floor(float64(numero) / 1000000)
	residuo := numero - (int(millones) * 1000000)

	textoMillones := ""
	textoMiles := miles(residuo)

	if millones > 1 {
		textoMillones = fmt.Sprintf("%s %s", miles(int(millones)), "MILLONES")
	} else if millones == 1 {
		textoMillones = "UN MILLON"
	} else {
		return textoMiles
	}

	return textoMillones + " " + textoMiles
}

func (monto Monto) Convertir() string {
	enteros := math.Floor(monto.Valor)
	centavos := math.Round(monto.Valor*100) - (math.Floor(monto.Valor) * 100)

	if enteros == 0 {
		if centavos != 0 && centavos <= 9 {
			return fmt.Sprintf("CERO %s 0%d/100", monto.TextoMonedaPlural, int(centavos))
		} else if centavos > 9 {
			return fmt.Sprintf("CERO %s %d/100", monto.TextoMonedaPlural, int(centavos))
		}
		return fmt.Sprintf("CERO %s 00/100", monto.TextoMonedaPlural)
	} else if enteros == 1 {
		if centavos != 0 && centavos <= 9 {
			return fmt.Sprintf("%s %s 0%d/100", millones(int(enteros)), monto.TextoMonedaSingular, int(centavos))
		} else if centavos > 9 {
			return fmt.Sprintf("%s %s %d/100", millones(int(enteros)), monto.TextoMonedaSingular, int(centavos))
		}
		return fmt.Sprintf("%s %s 00/100", millones(int(enteros)), monto.TextoMonedaSingular)
	} else {
		if centavos != 0 && centavos <= 9 {
			return fmt.Sprintf("%s %s 0%d/100", millones(int(enteros)), monto.TextoMonedaPlural, int(centavos))
		} else if centavos > 9 {
			return fmt.Sprintf("%s %s %d/100", millones(int(enteros)), monto.TextoMonedaPlural, int(centavos))
		}
		return fmt.Sprintf("%s %s 00/100", millones(int(enteros)), monto.TextoMonedaPlural)
	}
}