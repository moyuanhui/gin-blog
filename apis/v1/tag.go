package v1

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/gin-blog/models"
	"github.com/gin-blog/pkg/e"
	"github.com/gin-blog/pkg/setting"
	"github.com/gin-blog/util"
	"github.com/gin-gonic/gin"
)

// 获取多个标签
func GetTags(ctx *gin.Context) {
	name := ctx.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// 新增文章标签
func AddTag(ctx *gin.Context) {

}

// 修改文章标签
func EditTag(ctx *gin.Context) {

}

func DeleteTag(ctx *gin.Context) {

}
