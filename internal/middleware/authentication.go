package middleware

import (
	"fmt"
	"lion-super-app/configs"
	"lion-super-app/internal/abstraction"
	res "lion-super-app/pkg/util/response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	jwtKey := configs.Jwt().SecretKey()

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}
		var id int
		destructID := token.Claims.(jwt.MapClaims)["id"]
		if destructID != nil {
			id = int(destructID.(float64))
		} else {
			id = 0
		}

		var name string
		destructName, ok := token.Claims.(jwt.MapClaims)["name"]
		if ok && destructName != nil {
			name, ok = destructName.(string)
			if !ok {
				name = ""
			}
		} else {
			name = ""
		}

		cc := c.(*abstraction.Context)
		cc.Auth = &abstraction.AuthContext{
			ID: id,
			Name: name,
		}

		return next(cc)
	}
}
