package main

import (
	"time"

	"github.com/gin-gonic/gin"
)




type Activity struct {
	Start string
}

type CreateActivityRequest struct {
	Name          string
	Start         time.Time
	End           time.Time
	PromoProducts []*Product
}


type Product struct {

}


func createPromoActivity(ctx *gin.Context) {
	req := new(CreateActivityRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(400, gin.H{
			"msg": "请求错误",
		})
		return
	}
	
	start := req.Start.Unix()
	end := req.End.Unix()
	if start > end || time.Now().Unix() > start {
		return  // BizException 非法参数 illegal argument
	} 
	if req.PromoProducts == nil {
		return // 
	}

	// 创建

}

func main() {
	r := gin.New()
	r.Run()
}

/*
商品详情
	价格 吊牌价

	类目属性信息

	规格信息

	商品状态 1 在售 2 下架 3 删除

	店铺 id
	价格 秒杀价


订单初始化 order
	基本信息
	校验活动和商品状态
	扣减库存

		缓存
			db 性能差 索引 updates 行锁（硬盘不如内存快）

			redis 
				先 get decrease 不是原子
				decrease 原子，也不可以 -1 -1 100件 1000 人 -900
					有10 个人取消，+10 ，-890，这也没法消费了
					
				扣减库存，脚本 lua

	从后端获取秒杀价格
	订单号生成
	
*/


// func test() {
// 	ticker := time.NewTicker(5 * time.Minute)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			SyncLocalCache()  // 每隔五分钟刷新本地缓存数据
// 		}
// 	}
// }