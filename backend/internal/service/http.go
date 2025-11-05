/*
 * Copyright (c) 2025 KAnggara75
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara75 on Thu 06/11/25 05.55
 * @project backend service
 * https://github.com/KAnggara75/sentinel/tree/main/backend/internal/service
 */

package service

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	"github.com/KAnggara75/sentinel/backend/internal/config"
)

type AuthService struct {
	cfg *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{cfg: cfg}
}

func (s *AuthService) Login(c fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var r req
	if err := c.Bind().JSON(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	if r.Email != "admin@sentinel.local" || r.Password != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": r.Email,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenStr, _ := token.SignedString([]byte(s.cfg.JWTSecret))

	return c.JSON(fiber.Map{"token": tokenStr})
}

func (s *AuthService) Logout(c fiber.Ctx) error {
	// TODO: implement blacklist token (Redis or in-memory cache)
	return c.JSON(fiber.Map{"message": "logged out"})
}
