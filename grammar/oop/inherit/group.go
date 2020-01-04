package inherit

import "sync"

//广告信息
type Group struct {
	mu     sync.RWMutex
	Gid *string              //广告业务ID
	Status *int64            //广告投放状态. 1:未投放 2:投放中 3:余额不足 4:投放过期
	AdType *int64		 //广告类型
	BidMode *int64           //竞价方式. 1:cpm, 2:cpc (第一期只做cpm)
	Price *int64             //广告主出价
	InteractionType *int64   //交互类型. 1:打开网页 2:下载应用
	AppPackage *string       //下载应用包名
	AppSize *int64           //应用包大小
	LandingPage *string      //目标地址
	ThirdPlatformType *int64 //第三方检测平台类型(支持激活回调)
	ThirdPlatformURL *string //第三方检测平台地址(支持激活回调)
	TimeInterval *string     //投放时段.
	UserTotalImps *int64     //用户总曝光频次(保留一定时间)
	UserDayImps *int64       //用户日曝光频次
	UserTotalClicks *int64   //用户总点击频次(保留一定时间)
	UserDayClicks *int64     //用户日点击频次
	AdDayImps *int64         //广告日曝光控制
	AdDayClicks *int64       //广告日点击控制
	Targets map[int64]string //定向设置
}

type Groups struct {
	groups map[string]Group
}

