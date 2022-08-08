package uteis

import (
	"log"
	"regexp"
	"time"
)

// 01 - mês
// 02 - dia
// 03 - hora
// 04 - minuto
// 05 - segundos
// 06 - ano
// -07 - Time Zone
// 15 - modelo de horas 24h, substitui o 03
// 20 - adicionado ao 06, formata o ano em 4 dígitos. Exemplo: 2006, resultado 2022.
func DataGo1(ddmmyyyyy string) (data time.Time, err error) {

	ddmmyyyyy, err = RetirarSimbolos(ddmmyyyyy)
	if err != nil {
		log.Fatal(err)
	}

	data, err = time.Parse("02012006", ddmmyyyyy)
	return

}

func DataGo2(yyyymmdd string) (data time.Time, err error) {

	yyyymmdd, err = RetirarSimbolos(yyyymmdd)
	if err != nil {
		log.Fatal(err)
	}

	data, err = time.Parse("02012006", yyyymmdd)
	return

}

func RetirarSimbolos(info string) (texto string, err error) {
	reg, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	texto = reg.ReplaceAllString(info, "")
	return
}
