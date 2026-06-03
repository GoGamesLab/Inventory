package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/GoGamesLab/Inventory/pkg/container"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

func main() {
	Logger.Info("🧳 Invetory control start")

	i1 := container.NewInventory()

	it1 := container.Item{ID: 1, Name: "Item 1"}
	it2 := container.Item{ID: 2, Name: "Item 2"}
	it3 := container.Item{ID: 3, Name: "Item 3"}

	i1.AddItem(it1.ID, 10)
	i1.RemoveItem(it1.ID, 5)
	i1.AddItem(it2.ID, 20)
	i1.RemoveItem(it2.ID, 5)
	i1.AddItem(it3.ID, 30)
	i1.RemoveItem(it3.ID, 5)

	if _, ok := i1.AddItem(4, 40); ok {
		Logger.Info("🧨 Item 4 não deveria ser adicionado")
	}
	if _, ok := i1.RemoveItem(4, 5); ok {
		Logger.Info("🧨 Item 4 não deveria ser removido")
	}

	for i, quantity := range i1.Items {
		Logger.Info(fmt.Sprintf("🏷️ Item %d quantity %d", i, quantity))
	}
}
