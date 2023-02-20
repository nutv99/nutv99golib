package ppp

type Book999 struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}


type Book111 struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetSchoolAddress() string {
    return "กรุงเทพ";
}

func Aaaa() string {
	message := "Hello, "
	return message
}

func listBook2sHandler(c *gin.Context) {
	var book2s []Book2
	if result := Db.Find(&book2s); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &book2s)
}
