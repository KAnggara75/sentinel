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
 * @author KAnggara75 on Thu 06/11/25 05.54
 * @project backend http
 * https://github.com/KAnggara75/sentinel/tree/main/backend/internal/delivery/http
 */

package http

import (
	"github.com/gofiber/fiber/v3"

	"github.com/KAnggara75/sentinel/backend/internal/config"
	"github.com/KAnggara75/sentinel/backend/internal/delivery/http/middleware"
	"github.com/KAnggara75/sentinel/backend/internal/service"
)

func NewServer(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Sentinel",
	})

	authService := service.NewAuthService(cfg)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api")

	api.Post("/login", authService.Login)
	api.Post("/logout", middleware.Protected(cfg.JWTSecret), authService.Logout)

	return app
}
