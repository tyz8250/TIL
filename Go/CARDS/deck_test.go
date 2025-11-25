package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card of Ace of Spades,but got %v",d[0])
	}
	if d[len(d)-1] != "Four of Clabs"{
		t.Errorf("Expected last card of Four of Clabs,but got %v",d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // 既存のテスト用ファイルを消す

	deck := newDeck() // 新しいデッキを作る
	deck.saveToFile("_decktesting") // ファイルに保存する

	loadedDeck := newDeckFromFile("_decktesting")  // ファイルから読み込む

	// 読み込んだデッキの枚数が16枚か確認
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting") //後処理としてファイルを削除

}