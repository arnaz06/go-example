package controller

import(
	"../ApiHelpers"
	"../Models"
	"github.com/gin-gonic/gin"
)




func AllArticle(c *gin.Context){
	var article []Models.Article
	err := Models.GetAllArticles(&article)
	if err != nil{
		ApiHelpers.RespondJSON(c, 404, article)
	}
	ApiHelpers.RespondJSON(c, 200, article)
}

func CreateArticle(c *gin.Context) {
	var article Models.Article
	c.BindJSON(&article)
	err := Models.CreateArticle(&article)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, article)
	} else {
		ApiHelpers.RespondJSON(c, 200, article)
	}
}

// func NewArticle(w http.ResponseWriter, r *http.Request){
// 	Ping()
// 	// fmt.Println(r.Body)
// 	var article model.Article
// 	if err := json.NewDecoder(r.Body).Decode(&article); err != nil{
// 		fmt.Printf("%+v\n", article)
// 		panic(err)
// 	}
// 	if err:= db.Create(&article); err != nil{
// 		panic(err)
// 	}
// 	// fmt.Println(Article)
// }

// func Article(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Article Endpoint Hit")
// }

// func UpdateArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Update Article Endpoint Hit")
// }

// func DeleteArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Delete Article Endpoint Hit")
// }


// func SearchArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Search Article Endpoint Hit")
// }
