package middleware

import (
	"catplus-server/database"
	"catplus-server/model"
	"log"
)

func Totall() {

	var totalIncome int
	var totalExpense int

	db := database.GetDB()

	// 计算总收入
	if err := db.Model(&model.Item{}).Select("SUM(amount)").Where("kind = ? AND tag_ids IS NOT NULL", "income").Row().Scan(&totalIncome); err != nil {
		// 处理错误
		log.Println("Failed to calculate total income:", err)
		return
	}

	// 计算总支出
	if err := db.Model(&model.Item{}).Select("SUM(amount)").Where("kind = ? AND tag_ids IS NOT NULL", "expense").Row().Scan(&totalExpense); err != nil {
		// 处理错误
		log.Println("Failed to calculate total income:", err)
		return
	}

	log.Println("Total Income:", totalIncome)
	log.Println("Total Expense:", totalExpense)

}
