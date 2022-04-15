package api

import (
	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Param body body controllers.LoginParams true "body参数"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "bitbucket.org/udevs/example_api_gateway/api/docs"
	v1 "bitbucket.org/udevs/example_api_gateway/api/handlers/v1"
	"bitbucket.org/udevs/example_api_gateway/config"
	"bitbucket.org/udevs/example_api_gateway/pkg/logger"
	"bitbucket.org/udevs/example_api_gateway/services"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	router.GET("/config", handlerV1.GetConfig)

	apiV1 := router.Group("/v1")
	apiV1.GET("/ping", handlerV1.Ping)

	// profession
	apiV1.POST("/company", handlerV1.CreateCompany)
	apiV1.GET("/company/:company_id", handlerV1.GetCompany)
	apiV1.GET("/company", handlerV1.GetAllCompanies)
	apiV1.PUT("/company/:company_id",handlerV1.UpdateCompany)
	apiV1.DELETE("/company/:company_id",handlerV1.DeleteCompany)
	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
