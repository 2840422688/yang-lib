package Utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var AccessTokenKey = []byte("www.innyang.cn")
var RefreshTokenKey = []byte("innyang.cn")

const ACCESS_TOKEN_EXPIRE_TIME = time.Hour * 2
const REFRESH_TOKEN_EXPIRE_TIME = time.Hour * 48

type cliam struct {
	UserAccount string `json:"user_account"`
	User_id     int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenAccessToken(userAccount string, User_id int) (string, int64, error) {
	cliams := cliam{
		userAccount,
		User_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TOKEN_EXPIRE_TIME)),
			Issuer:    "SSO",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams).SignedString(AccessTokenKey)
	if err != nil {
		return "", 0.0, err
	}
	return token, int64(ACCESS_TOKEN_EXPIRE_TIME.Seconds()), nil
}

func GenRefreshToken(userAccount string, User_id int) (string, int64, error) {
	cliams := cliam{
		userAccount,
		User_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(REFRESH_TOKEN_EXPIRE_TIME)),
			Issuer:    "SSO",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams).SignedString(RefreshTokenKey)
	if err != nil {
		return "", 0.0, err
	}
	return token, int64(REFRESH_TOKEN_EXPIRE_TIME.Seconds()), nil
}
func ParseToken(tokenString string) (*cliam, error) {
	// 解析token
	if !strings.HasPrefix(tokenString, "Bearer%20") {
		return nil, errors.New("非法令牌")
	}
	tokenString = strings.Replace(tokenString, "Bearer%20", "", 1)
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &cliam{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return AccessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*cliam); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//func SplicingSql(loopStruct interface{}, queryTable string, findAll bool) string {
//	sql := " "
//	queryName, queryValue := []interface{}{}, []interface{}{}
//	kind := reflect.TypeOf(loopStruct).Kind().String()
//	if kind == "struct" {
//		fieldNum := reflect.ValueOf(loopStruct).NumField()
//		for i := 0; i < fieldNum; i++ {
//			//nameOfValue := reflect.TypeOf(loopStruct).Field(i).Name
//			nameOfValue := reflect.TypeOf(loopStruct).Field(i).Tag.Get("gorm")
//			typeOfValue := reflect.TypeOf(loopStruct).Field(i).Type.Kind().String()
//			switch typeOfValue {
//			case "string":
//				valueString := reflect.ValueOf(loopStruct).Field(i).String()
//				if valueString != "" {
//					queryName = append(queryName, nameOfValue)
//					queryValue = append(queryValue, valueString)
//				}
//				break
//			case "int":
//				valueInt := reflect.ValueOf(loopStruct).Field(i).Int()
//				if valueInt != 0 {
//					queryName = append(queryName, nameOfValue)
//					queryValue = append(queryValue, valueInt)
//				}
//				break
//			}
//		}
//	}
//
//	if len(queryValue) == 0 && len(queryName) == 0 {
//		sql = ""
//		return sql
//	}
//	sqlWhere := ""
//	sqlSelect := ""
//	for i := 0; i < len(queryName); i++ {
//		if len(queryName)-i > 1 {
//			sqlSelect = sql + queryName[i].(string) + " , "
//		}
//		sqlSelect = sql + queryName[i].(string) + " "
//		if findAll {
//			sqlSelect = " * "
//		}
//		typeV := reflect.TypeOf(queryValue[i]).Kind().String()
//		switch typeV {
//		case "string":
//			if len(queryName)-i <= 1 {
//				sqlWhere = sqlWhere + strings.ToLower(queryName[i].(string)) + " = " + "'" + queryValue[i].(string) + "'"
//				break
//			}
//			sqlWhere = sqlWhere + queryName[i].(string) + " = " + queryValue[i].(string) + " AND "
//			break
//		case "int64":
//			if len(queryName)-i <= 1 {
//				sqlWhere = sqlWhere + strings.ToLower(queryName[i].(string)) + " = " + strconv.Itoa(int(queryValue[i].(int64)))
//				break
//			}
//			sqlWhere = sqlWhere + strings.ToLower(queryName[i].(string)) + " = " + strconv.Itoa(int(queryValue[i].(int64))) + " AND "
//			break
//		}
//	}
//	sql = "SELECT " + sqlSelect + " FROM " + queryTable + " WHERE " + sqlWhere + ";"
//	return sql
//}

func SplicingSql(loopStruct interface{}, queryTable string, findAll bool) string {
	var sql strings.Builder

	// 获取结构体类型
	structType := reflect.TypeOf(loopStruct)

	// 如果不是结构体，直接返回空字符串
	if structType.Kind() != reflect.Struct {
		return ""
	}

	// 初始化 WHERE 子句
	var conditions []string

	// 遍历结构体的字段
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("gorm")
		value := reflect.ValueOf(loopStruct).Field(i)

		// 根据字段类型处理值
		switch value.Kind() {
		case reflect.String:
			if strValue := value.String(); strValue != "" {
				conditions = append(conditions, tag+" = '"+strValue+"'")
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intValue := value.Int(); intValue != 0 {
				conditions = append(conditions, tag+" = '"+strconv.FormatInt(intValue, 10)+"'")
			}
		}
	}

	// 如果没有查询条件，则返回空字符串
	if len(conditions) == 0 {
		return ""
	}

	// 构建 SELECT 子句
	sqlSelect := "*"
	if !findAll {
		sqlSelect = strings.Join(getGormTags(structType), ", ")
	}

	// 构建 WHERE 子句
	sqlWhere := strings.Join(conditions, " AND ")

	// 构建 SQL 查询语句
	sql.WriteString("SELECT " + sqlSelect + " FROM " + queryTable + " WHERE " + sqlWhere + ";")
	return sql.String()
}

// 获取结构体的字段的 GORM 标签
func getGormTags(structType reflect.Type) []string {
	var tags []string
	for i := 0; i < structType.NumField(); i++ {
		tag := structType.Field(i).Tag.Get("gorm")
		tags = append(tags, tag)
	}
	return tags
}

// 生成BC加密密文
func GenBcryptStr(content []byte, rounds int) (string, error) {
	result, err := bcrypt.GenerateFromPassword(content, rounds)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

//hmac加密
func HmacEncode(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
