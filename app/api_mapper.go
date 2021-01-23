package app
import 	"awacs_smart_api_service/controllers"

func mapUrls() {

	router.GET("/", controllers.PlaygroundHandler())

	router.POST("/query", controllers.GraphqlHandler())

}
