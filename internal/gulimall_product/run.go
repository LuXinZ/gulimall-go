package run

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gulimall-product/internal/gulimall_product/dao"
	"gulimall-product/internal/pkg/core"
	"gulimall-product/internal/pkg/middleware"
	"log"
	"net/http"
)

func Run() error {
	gormdb, _ := gorm.Open(mysql.Open("root:root@(127.0.0.1:3306)/gulimall_pms?charset=utf8mb4&parseTime=True&loc=Local"))
	dao.SetDefault(gormdb)
	cate := dao.Q.PmsCategory
	user, err := cate.Where(cate.CatID.Eq(1)).First()
	if err != nil {

	}
	// 创建 Gin 引擎
	r := gin.New()
	// 全局中间件
	// Logger 中间件将日志写到 gin.DefaultWriter，即使设置了 GIN_MODE=release
	// 默认设置 gin.DefaultWriter = os.Stdout
	// Recovery 中间件，从任何 panic 恢复，并返回一个 500 错误
	r.Use(gin.Recovery(), gin.Logger(), middleware.Cors, middleware.RequestID())
	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: ":50002", Handler: r}
	r.GET("/healthz", func(c *gin.Context) {

		core.WriteResponse(c, nil, user)
	})
	if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("can not start server %v", err)
	}

	return nil
}
