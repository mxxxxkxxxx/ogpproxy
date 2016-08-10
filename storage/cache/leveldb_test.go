package cache

import "testing"

func TestLevelDBHandler_Write_Read(t *testing.T) {
	handler := GetLevelDBHandler()

	key   := "test_key"
	value := "test_value"
	err   := handler.Write(key, []byte(value))
	if err != nil {
		t.Fatalf("Failed to write: err=[%s]", err)
	}

	data, err := handler.Read(key)
	if err != nil {
		t.Fatalf("Failed to read: err=[%s]", err)
	}

	if value != string(data) {
		t.Fail()
	}
}

func TestLevelDBHandler_Write_Failed_WithoutKey(t *testing.T) {
	handler := GetLevelDBHandler()

	key   := ""
	value := "test_value"
	err   := handler.Write(key, []byte(value))
	if err == nil {
		t.Fatal("Succeeded to write without key")
	}
}

func TestLevelDBHandler_Read_Failed_WithoutKey(t *testing.T) {
	handler := GetLevelDBHandler()

	_, err := handler.Read("")
	if err == nil {
		t.Fatal("Succeeded to read without key")
	}
}