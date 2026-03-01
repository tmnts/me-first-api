package main

import "github.com/gin-gonic/gin"

func main() {
	// 1. Создаем сервер (роутер)
	r := gin.Default()

	// 1.1 Оформляем страницу 404 по-ямайски
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Me-First-API!",
			"status":  "Me a go document di whole tingg ME A GO BE DEVMAN",
			"hint":    "U a go try /hello or /dog/sharik an MAGIC A GO HAPPEN",
		})
	})

	// 2. Определяем Base URL и Endpoint
	// Когда кто-то зайдет на GET /hello
	r.GET("/hello", func(c *gin.Context) {
		// 3. Формируем JSON ответ (200 OK)
		c.JSON(200, gin.H{
			"message": "Привет! Это твой первый API",
			"status":  "working",
		})
	})

	// 4. Endpoint с параметром (например, имя собаки)
	// GET /dog/sharik
	r.GET("/dog/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"dog_name": name,
			"action":   "waiting for hood (rich)",
		})
	})

	// POST-метод: принимаем данные о новой собаке
	r.POST("/dog", func(c *gin.Context) {
		// Создаем структуру для входящих данных
		var json struct {
			Name  string `json:"name"`
			Breed string `json:"breed"`
		}

		// Пробуем "привязать" пришедший JSON к нашей структуре
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "Вот скажи мне, американец. В чем разница между GET и POST? Вот и брат говорит, что большая разница. А на деле - кривой ты джсон прислал, американец."})
			return
		}

		// Если все ок, отвечаем статусом 201 (Created)
		c.JSON(201, gin.H{
			"status":  "собака добавлена в базу",
			"name":    json.Name,
			"breed":   json.Breed,
			"message": "U A GO BE PACKMAN",
		})
	})

	// 5. Запуск сервера на порту 8080
	r.Run(":8080")
}
