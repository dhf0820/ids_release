package common

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
)

func NormalizePhone(phone string) string {
	fmt.Printf("Normalizing : %s\n", phone)
	reg, err := regexp.Compile("[^0-9]*")
	if err != nil {
		log.Errorf("regexp compile failed: %v", err)
	}
	return reg.ReplaceAllString(phone, "")
}

//https://www.golangprograms.com/regular-expression-to-validate-email-address.html
func ValidateEmail(email string) bool {
	fmt.Printf("Validating : %s\n", email)
	str := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	re := regexp.MustCompile(str)
	return re.MatchString(email)
}

func SetDashToBlank(inStr string) string {
	m1 := regexp.MustCompile(`^-$`)
	return m1.ReplaceAllString(inStr, "")
}

//func GetFieldByName(flds []Field, name string) (*Field, error) {
//	for _, fld := range flds {
//		if fld.Name == name {
//			return &fld, nil
//		}
//	}
//	return nil, status.Errorf(
//		codes.NotFound,
//		fmt.Sprintf("Field name: %s was not found", name),
//	)
//}
