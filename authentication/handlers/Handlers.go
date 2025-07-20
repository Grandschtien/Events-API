package handlers

import (
	refreshTokens "events-api/core/db/RefreshTokensDB"
	users "events-api/core/db/usersDB"
)

type AuthHandlers struct {
	UsersDB         *users.UsersDB
	RefreshTokensDB *refreshTokens.RefreshTokensDB
}
