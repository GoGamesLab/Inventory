package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/GoGamesLab/Inventory/pkg/container"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

func main() {
	Logger.Info("🧳 Inventory control start")

	i1 := container.NewInventory()

	it1 := container.Item{ID: 100, Name: "Item 1"}
	it2 := container.Item{ID: 200, Name: "Item 2"}
	it3 := container.Item{ID: 300, Name: "Item 3"}

	i1.AddItem(it1.ID, 10)
	i1.RemoveItem(it1.ID, 5)
	i1.AddItem(it2.ID, 20)
	i1.RemoveItem(it2.ID, 5)
	i1.AddItem(it3.ID, 30)
	i1.RemoveItem(it3.ID, 5)

	for i, quantity := range i1.Items {
		Logger.Info(fmt.Sprintf("🏷️ Item %d quantity %d", i, quantity))
	}
}
