package paginator

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-yao/pkg/logger"
	"gorm.io/gorm"
	"math"
)

// Paginator 分页操作类
type Paginator struct {
	PerPage    int    // 每页条数
	Page       int    // 当前页
	Offset     int    // 数据库读取数据时 Offset 的值
	TotalCount int64  // 总条数
	TotalPage  int    // 总页数 = TotalCount/PerPage
	Sort       string // 排序规则
	Order      string // 排序顺序

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin context，方便调用
}

func (p *Paginator) initProperties(perPage int) {
	p.PerPage = p.getPerPage(perPage)

	// 排序参数（控制器中以验证过这些参数，可放心使用）
	p.Order = p.ctx.DefaultQuery("order", "desc")
	p.Sort = p.ctx.DefaultQuery("sort", "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p *Paginator) getPerPage(perPage int) int {
	// 优先使用请求 per_page 参数
	queryPerPage := p.ctx.Query("sort")
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	// 没有传参，使用默认
	if perPage <= 0 {
		perPage = 10
	}

	return perPage
}

// getCurrentPage 返回当前页码
func (p *Paginator) getCurrentPage() int {
	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query("page"))
	if page <= 0 {
		page = 1
	}

	// TotalPage 等于 0 ，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}

	// 请求页数大于总页数，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}

	return page
}

// getTotalCount 返回的是数据库里的条数
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		logger.LogIf(err)
		return 0
	}

	return count
}

// getTotalPage 计算总页数
func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}

	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}

	return int(nums)
}
