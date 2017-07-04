package demo

type apiSlot struct {
	Id             int          `json:"Id"`
	Name           string       `json:"Name"`   //名称
	Width          int          `json:"Width"`  //宽,px
	Height         int          `json:"Height"` //高,px
	Site           *Site        `json:"Site"`
	ViewType       *apiViewType `json:"ViewType"`       // 展现形式
	CreativeType   []int        `json:"CreativeType"`   //支持格式 gif:10001,jpg:10002,png:10003,swf:20001
	Screen         int          `json:"Screen"`         //屏次 0,未知；1,第一屏；2，第二屏; 3,第三屏；4，第四屏；5，第五屏；101，非首屏；102，五屏以外，不能修改
	MobileSizeType int          `json:"MobileSizeType"` // 当 ViewType 为“移动固定/移动悬浮”广告位时使用；1=自适应尺寸；2=固定高度，宽度自适应
	Control        int          `json:"Control"`        // 广告是否可以关闭，0 不可关闭，1 可以关，移动悬浮（屏幕固定）支持

	Logo int `json:"Logo"` // 是否显示角标, 0显示，9不显示

	Description string `json:"Description"` // 说明
}

type Site struct {
	Name       string `bson:"Name"`
	Url        string `bson:"Url"`        //网站首页,请填写以http://或https://开头的完整URL
	CategoryId []int  `bson:"CategoryId"` //网站分类,格式：[父分类Id, 子分类Id],参考所给文档
	ICP        string `bson:"ICP"`        //ICP备案信息
	Summary    string `bson:"Summary"`    //简介
}

type apiViewType struct {
	Id  int `json:"Id"`       // 1:固定，2:漂浮，3:对联，4:文字链，5:弹窗，6:移动网页固定，7:移动悬浮（屏幕固定），11:视频贴片，12:视频暂停
	YId int `json:"Location"` // 1:居上，2:居下
}
