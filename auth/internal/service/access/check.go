package access

import (
	"auth/internal/model"
	"context"
	"fmt"
)

func (s serv) Check(ctx context.Context, url string, claims *model.UserClaims) (bool, error) {
	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get accessible roles")
	}

	accessRole, ok := accessibleMap[url]

	if !ok {
		return false, nil
	}

	if accessRole == claims.Role {
		return true, nil
	}

	return false, fmt.Errorf("access denied")
}

var accessibleRoles map[string]string

// Возвращает мапу с адресом эндпоинта и ролью, которая имеет доступ к нему
func (s serv) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]string)

		// Лезем в базу за данными о доступных ролях для каждого эндпоинта
		// Можно кэшировать данные, чтобы не лезть в базу каждый раз

		// Например, для эндпоинта /note_v1.NoteV1/Get доступна только роль admin
		accessibleRoles[model.ExamplePath] = "USER"
		accessibleRoles[model.ChatPath] = "USER"
	}

	return accessibleRoles, nil
}
