package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GeneratorPassword(passwordLength int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	password := make([]byte, passwordLength)

	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(password)
}

func GeneratorValidCNPJ() string {
	digitos := make([]int, 12)
	for i := range 12 {
		digitos[i] = rand.Intn(10)
	}

	soma := 0
	pesos := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := range 12 {
		soma += digitos[i] * pesos[i]
	}
	resto := soma % 11
	if resto < 2 {
		digitos = append(digitos, 0)
	} else {
		digitos = append(digitos, 11-resto)
	}

	soma = 0
	pesos = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := range 13 {
		soma += digitos[i] * pesos[i]
	}
	resto = soma % 11
	if resto < 2 {
		digitos = append(digitos, 0)
	} else {
		digitos = append(digitos, 11-resto)
	}

	resultado := ""
	for _, digito := range digitos {
		resultado += strconv.Itoa(digito)
	}
	return resultado
}

func GenerateValidCPF() string {
	digitos := make([]int, 9)
	for i := range 9 {
		digitos[i] = rand.Intn(10)
	}

	soma := 0
	for i := range 9 {
		soma += digitos[i] * (10 - i)
	}
	resto := soma % 11
	if resto < 2 {
		digitos = append(digitos, 0)
	} else {
		digitos = append(digitos, 11-resto)
	}

	soma = 0
	for i := range 10 {
		soma += digitos[i] * (11 - i)
	}
	resto = soma % 11
	if resto < 2 {
		digitos = append(digitos, 0)
	} else {
		digitos = append(digitos, 11-resto)
	}

	resultado := ""
	for _, digito := range digitos {
		resultado += strconv.Itoa(digito)
	}
	return resultado
}

func FormatCNPJ(cnpj string) string {
	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:])
}

func FormatCPF(cpf string) string {
	return fmt.Sprintf("%s.%s.%s-%s", cpf[:3], cpf[3:6], cpf[6:9], cpf[9:])
}
