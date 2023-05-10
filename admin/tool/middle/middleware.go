package middle

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"opensearch-tool/tool/domain"
	"opensearch-tool/tool/open"
)

func CheckUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.RequestURI == "/" {
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
			ctx.Abort()
			return
		}

		auth := ctx.Request.Header.Get("Authorization")
		auth = strings.Trim(auth, " ")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token不存在"})
			ctx.Abort()
			return
		}
		decoderBytes, err := base64.StdEncoding.DecodeString(auth)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token不正确"})
			ctx.Abort()
			return
		}
		info := strings.Split(string(decoderBytes), "###")
		if len(info) != 3 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token不正确"})
			ctx.Abort()
			return
		}
		ctx.Set("opensearch:host", info[0])
		ctx.Set("opensearch:user", info[1])
		ctx.Set("opensearch:pwd", info[2])
	}
}

func Login(ctx *gin.Context) {
	var userInfo domain.UserInfo
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	decodePwd, err := base64.StdEncoding.DecodeString(userInfo.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "解密错误"})
		return
	}
	encodeStr := base64.StdEncoding.EncodeToString([]byte(userInfo.Host + "###" + userInfo.User + "###" + string(decodePwd)))
	ctx.JSON(http.StatusOK, gin.H{"msg": "成功", "code": 200, "token": encodeStr})
}

func Index(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)
	tool, err := open.NewOpenTool(userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	info, err := tool.Info()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "成功", "code": 200, "data": info})
}

func Refresh(ctx *gin.Context) {
	userInfo := getUserInfo(ctx)
	tool, err := open.NewOpenTool(userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	health, err := tool.CatHealth()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	nodes, err := tool.CatNodes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	indices, err := tool.CatIndices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	shards, err := tool.CatShards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data := make(map[string]interface{})
	data["health"] = health
	data["nodes"] = nodes
	data["indices"] = indices
	data["shards"] = shards
	ctx.JSON(http.StatusOK, gin.H{"msg": "成功", "code": 200, "data": data})
}

func getUserInfo(ctx *gin.Context) (userInfo domain.UserInfo) {
	userInfo.Host = ctx.GetString("opensearch:host")
	userInfo.User = ctx.GetString("opensearch:user")
	userInfo.Password = ctx.GetString("opensearch:pwd")
	return
}
