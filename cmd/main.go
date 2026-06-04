package main

import (
	"fmt"
	"log/slog"
	"os"

	container "github.com/GoGamesLab/Inventory/pkg"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

func main() {
	Logger.Info("🧳 Inventory control start")

	i1 := container.NewInventory()

	it1 := container.ItemID(100)
	it2 := container.ItemID(100)
	it3 := container.ItemID(100)

	i1.AddItem(it1, 10)
	i1.RemoveItem(it1, 5)
	i1.AddItem(it2, 20)
	i1.RemoveItem(it2, 5)
	i1.AddItem(it3, 30)
	i1.RemoveItem(it3, 5)

	for i, quantity := range i1.Items {
		Logger.Info(fmt.Sprintf("🏷️ Item %d quantity %d", i, quantity))
	}
}
