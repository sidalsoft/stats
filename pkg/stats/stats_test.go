package stats

import (
	"github.com/sidalsoft/bank/v2/pkg/types"
	"reflect"
	"testing"
)

//func TestFilterByCategory_nil(t *testing.T) {
//	var payments []types.Payment
//	result := FilterByCategory(payments, "mobile")
//	if len(result) != 0 {
//		t.Error("result len!=0")
//	}
//}
//
//func TestFilterByCategory_empty(t *testing.T) {
//	payments := []types.Payment{}
//	result := FilterByCategory(payments, "mobile")
//	if len(result) != 0 {
//		t.Error("result len!=0")
//	}
//}
//
//func TestFilterByCategory_notFound(t *testing.T) {
//	payments := []types.Payment{
//		{ID: 1, Category: "auto"},
//		{ID: 2, Category: "food"},
//		{ID: 3, Category: "auto"},
//		{ID: 4, Category: "auto"},
//		{ID: 5, Category: "fun"},
//	}
//	result := FilterByCategory(payments, "mobile")
//	if len(result) != 0 {
//		t.Error("result len !=0")
//	}
//}
//
//func TestFilterByCategory_foundOne(t *testing.T) {
//	payments := []types.Payment{
//		{ID: 1, Category: "auto"},
//		{ID: 2, Category: "food"},
//		{ID: 3, Category: "auto"},
//		{ID: 4, Category: "auto"},
//		{ID: 5, Category: "fun"},
//	}
//	expected := []types.Payment{
//		{ID: 2, Category: "food"},
//	}
//	result := FilterByCategory(payments, "food")
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Error("invalid result")
//	}
//}
//
//func TestFilterByCategory_foundMultiple(t *testing.T) {
//	payments := []types.Payment{
//		{ID: 1, Category: "auto"},
//		{ID: 2, Category: "food"},
//		{ID: 3, Category: "auto"},
//		{ID: 4, Category: "auto"},
//		{ID: 5, Category: "fun"},
//	}
//	expected := []types.Payment{
//		{ID: 1, Category: "auto"},
//		{ID: 3, Category: "auto"},
//		{ID: 4, Category: "auto"},
//	}
//	result := FilterByCategory(payments, "auto")
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Errorf("Invalid result, expected : %v, actual %v", expected, result)
//	}
//}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 5},
		{ID: 4, Category: "auto", Amount: 6},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 7,
		"food": 2_000_00,
		"fun":  5_000_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result, expected : %v, actual %v", expected, result)
	}
}
