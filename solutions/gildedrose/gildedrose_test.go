package gildedrose

import (
	"testing"
)

func TestUpdateQuality_regularItem(t *testing.T) {
	items := []*Item{
		{"foo", 10, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "foo" {
		t.Errorf("Expected name to be foo, got %s", items[0].Name)
	}
	if items[0].SellIn != 9 {
		t.Errorf("Expected sellIn to be 9, got %d", items[0].SellIn)
	}
	if items[0].Quality != 9 {
		t.Errorf("Expected quality to be 9, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_sellDatePassed(t *testing.T) {
	items := []*Item{
		{"foo", 0, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "foo" {
		t.Errorf("Expected name to be foo, got %s", items[0].Name)
	}
	if items[0].SellIn != -1 {
		t.Errorf("Expected sellIn to be -1, got %d", items[0].SellIn)
	}
	if items[0].Quality != 8 {
		t.Errorf("Expected quality to be 8, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_qualityIsNeverNegative(t *testing.T) {
	items := []*Item{
		{"foo", 10, 0},
	}

	UpdateQuality(items)

	if items[0].Name != "foo" {
		t.Errorf("Expected name to be foo, got %s", items[0].Name)
	}
	if items[0].SellIn != 9 {
		t.Errorf("Expected sellIn to be 9, got %d", items[0].SellIn)
	}
	if items[0].Quality != 0 {
		t.Errorf("Expected quality to be 0, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_agedBrie(t *testing.T) {
	items := []*Item{
		{"Aged Brie", 10, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Aged Brie" {
		t.Errorf("Expected name to be Aged Brie, got %s", items[0].Name)
	}
	if items[0].SellIn != 9 {
		t.Errorf("Expected sellIn to be 9, got %d", items[0].SellIn)
	}
	if items[0].Quality != 11 {
		t.Errorf("Expected quality to be 11, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_agedBrie_sellDatePassed(t *testing.T) {
	items := []*Item{
		&Item{"Aged Brie", 0, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Aged Brie" {
		t.Errorf("Expected name to be Aged Brie, got %s", items[0].Name)
	}
	if items[0].SellIn != -1 {
		t.Errorf("Expected sellIn to be -1, got %d", items[0].SellIn)
	}
	if items[0].Quality != 12 {
		t.Errorf("Expected quality to be 12, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_agedBrie_qualityIsNeverMoreThan50(t *testing.T) {
	items := []*Item{
		{"Aged Brie", 10, 50},
	}

	UpdateQuality(items)

	if items[0].Name != "Aged Brie" {
		t.Errorf("Expected name to be Aged Brie, got %s", items[0].Name)
	}
	if items[0].SellIn != 9 {
		t.Errorf("Expected sellIn to be 9, got %d", items[0].SellIn)
	}
	if items[0].Quality != 50 {
		t.Errorf("Expected quality to be 50, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_sulfuras(t *testing.T) {
	items := []*Item{
		{"Sulfuras, Hand of Ragnaros", 10, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Sulfuras, Hand of Ragnaros" {
		t.Errorf("Expected name to be Sulfuras, Hand of Ragnaros, got %s", items[0].Name)
	}
	if items[0].SellIn != 10 {
		t.Errorf("Expected sellIn to be 10, got %d", items[0].SellIn)
	}
	if items[0].Quality != 10 {
		t.Errorf("Expected quality to be 10, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_backstagePasses(t *testing.T) {
	items := []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", 20, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Backstage passes to a TAFKAL80ETC concert" {
		t.Errorf("Expected name to be Backstage passes to a TAFKAL80ETC concert, got %s", items[0].Name)
	}
	if items[0].SellIn != 19 {
		t.Errorf("Expected sellIn to be 19, got %d", items[0].SellIn)
	}
	if items[0].Quality != 11 {
		t.Errorf("Expected quality to be 11, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_backstagePasses10Days(t *testing.T) {
	items := []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", 9, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Backstage passes to a TAFKAL80ETC concert" {
		t.Errorf("Expected name to be Backstage passes to a TAFKAL80ETC concert, got %s", items[0].Name)
	}
	if items[0].SellIn != 8 {
		t.Errorf("Expected sellIn to be 8, got %d", items[0].SellIn)
	}
	if items[0].Quality != 12 {
		t.Errorf("Expected quality to be 12, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_backstagePasses5Days(t *testing.T) {
	items := []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", 4, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Backstage passes to a TAFKAL80ETC concert" {
		t.Errorf("Expected name to be Backstage passes to a TAFKAL80ETC concert, got %s", items[0].Name)
	}
	if items[0].SellIn != 3 {
		t.Errorf("Expected sellIn to be 3, got %d", items[0].SellIn)
	}
	if items[0].Quality != 13 {
		t.Errorf("Expected quality to be 13, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_backstagePassesExpired(t *testing.T) {
	items := []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", 0, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Backstage passes to a TAFKAL80ETC concert" {
		t.Errorf("Expected name to be Backstage passes to a TAFKAL80ETC concert, got %s", items[0].Name)
	}
	if items[0].SellIn != -1 {
		t.Errorf("Expected sellIn to be -1, got %d", items[0].SellIn)
	}
	if items[0].Quality != 0 {
		t.Errorf("Expected quality to be 0, got %d", items[0].Quality)
	}
}

func TestUpdateQuality_conjured(t *testing.T) {
	items := []*Item{
		{"Conjured Mana Cake", 10, 10},
	}

	UpdateQuality(items)

	if items[0].Name != "Conjured Mana Cake" {
		t.Errorf("Expected name to be Conjured Mana Cake, got %s", items[0].Name)
	}
	if items[0].SellIn != 9 {
		t.Errorf("Expected sellIn to be 9, got %d", items[0].SellIn)
	}
	if items[0].Quality != 8 {
		t.Errorf("Expected quality to be 8, got %d", items[0].Quality)
	}
}
