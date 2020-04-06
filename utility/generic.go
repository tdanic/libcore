package utility

import (
	"fmt"
	ramdom "math/rand"
	"strings"
	"time"
	"unicode"

	uuid "github.com/satori/go.uuid"
)

/*ValidDuplidArrayStr : valida un arreglo de string si estan duplicados*/
func ValidDuplidArrayStr(strs []string) bool {
	for i, str := range strs {
		for j, straux := range strs {
			if i != j && straux == str {
				return false
			}
		}
	}
	return true
}

/*InStr : compara  varios string con uno especifico*/
func InStr(str string, strs ...string) bool {
	for _, item := range strs {
		if str == item {
			return true
		}
	}
	return false
}

/*FilterExcl : excluye los registro de un arreglo A  con el B  ejemplo
A[a,b,c,d]
B[a,b]
Result
[c,d]
*/
func FilterExcl(strs []string, excl []string) []string {
	var ret []string
	for _, item := range strs {
		ind := true
		for _, item2 := range excl {
			if item2 == item {
				ind = false
			}
		}
		if ind {
			ret = append(ret, item)
		}
	}
	return ret
}

/*FilterStr : filtra un arreglo de string mediante un metodo definido */
func FilterStr(strs []string, valid func(string) bool) (ret []string) {
	for _, s := range strs {
		if valid(s) {
			ret = append(ret, s)
		}
	}
	return
}

/*ReturnIf : retorna un if ternario
https://github.com/TheMushrr00m/go-ternary Doc
ReturnIf(<bool expression>, <result for true>, <result for false>)
ReturnIf(5 > 4, "It's true", "It's false :(")
*/
func ReturnIf(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}

	return c
}

/*StrRand : genera una cadena de caracteres ramdon*/
func StrRand(cant int) string {
	if cant <= 0 {
		cant = 1
	}
	str := ""
	for i := 0; i < cant; i++ {
		str += CharRand()
	}
	return str
}

/*CharRand : Genera una letra aleatoria.*/
func CharRand() string {
	ramdom.Seed(time.Now().UnixNano())
	return string(ramdom.Intn(122) + 97)
}

/*SubString : substring para un string en golang con runas*/
func SubString(cadena string, ini, cant int) string {
	runes := []rune(cadena)
	return string(runes[0:2])
}

/*Trim : Elimina el espacio de un string a nivel de runas.*/
func Trim(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {

			return -1
		}
		return r
	}, str)
}

/*IsSpace : valida si la cadena contiene espacio. */
func IsSpace(str string) bool {
	rest := false
	strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			rest = true
			return -1
		}
		return r
	}, str)
	return rest
}

/*GeneredUUID : Genera un codigo uuid unico */
func GeneredUUID() string {
	var err error
	return fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), err))
}
