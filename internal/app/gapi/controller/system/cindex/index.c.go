package cindex

import (
	"github.com/DicoderCn/ginp"
	"github.com/gin-gonic/gin"
)

func IndexView(c *ginp.ContextPlus) {
	c.HTML(200, "index.html", gin.H{})
}
