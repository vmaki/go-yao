package paginator

import (
	"github.com/gin-gonic/gin"
	"go-yao/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int   // 当前页
	PerPage     int   // 每页条数
	TotalPage   int   // 总页数
	TotalCount  int64 // 总条数
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, perPage int) Paging {
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage)

	// 查询数据库
	err := p.query.Preload(clause.Associations). // 读取关联
							Order(p.Sort + " " + p.Order). // 排序
							Limit(p.PerPage).
							Offset(p.Offset).
							Find(data).
							Error
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
	}
}
