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
 * @author KAnggara75 on Thu 06/11/25 05.52
 * @project backend sentinel
 * https://github.com/KAnggara75/sentinel/tree/main/backend/cmd/sentinel
 */

package main

import (
	"log/slog"

	"github.com/KAnggara75/sentinel/backend/internal/config"
	"github.com/KAnggara75/sentinel/backend/internal/delivery/http"
)

func main() {
	cfg := config.Load()

	app := http.NewServer(cfg)

	slog.Info("starting sentinel service", "port", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
