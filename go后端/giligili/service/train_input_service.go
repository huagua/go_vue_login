package service

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
)

// UserRegisterService 管理用户注册服务
type TrainInputService struct {
	TrainNum  string `form:"TrainNum" json:"TrainNum"`
	From      string `form:"From" json:"From"`
	To        string `form:"To" json:"To"`
	StartTime string `form:"StartTime" json:"StartTime"`
	EndTime   string `form:"EndTime" json:"EndTime"`
	Duration  string `form:"Duration" json:"Duration"`
	Date      string `form:"Date" json:"Date"`
	Super     uint   `form:"Super" json:"Super"`
	First     uint   `form:"First" json:"First"`
	Second    uint   `form:"Second" json:"Second"`
}



// valid 验证表单
func (service *TrainInputService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.Train{}).Where("train_num = ?", service.TrainNum).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "已有该车次",
		}
	}
	addFree(service)
	addCache(service)

	return nil
}

func addCache(service *TrainInputService) {
	cache.RedisClient.LPush(cache.GetSearchKey(service.From, service.To), service.TrainNum)

	train := make(map[string]interface{})
	train["startTime"] = service.StartTime
	train["duration"] = service.Duration
	train["user"] = ""
	cache.RedisClient.HMSet(cache.GetTrainKey(service.TrainNum, service.Date), train)

	super := make(map[string]interface{})
	super["left"] = service.Super
	super["money"] = 390
	cache.RedisClient.HMSet(cache.GetTicketKey(service.Date, service.TrainNum,"super"),super)

	first := super
	first["left"] = service.First
	first["Money"] = 250
	cache.RedisClient.HMSet(cache.GetTicketKey(service.Date, service.TrainNum,"first"),first)

	second := super
	second["left"] = service.Second
	second["money"] = 150
	cache.RedisClient.HMSet(cache.GetTicketKey(service.Date, service.TrainNum, "second"),second)
}

func addFree(service *TrainInputService) {
	insertRow(service.TrainNum, service.Date, "Super", service.Super, 1, 1, 10)
	insertRow(service.TrainNum, service.Date, "First", service.First, 2, 3, 12)
	insertRow(service.TrainNum, service.Date, "Second", service.Second, 4, 6, 20)
}

func insertRow(nums string, date string, rank string, tmp uint, from int, to int, rows int) {
	for i := from; i <= to; i++ {
		for j := 1; j <= rows; j++ {
			freeTicket := model.FreeTicket{
				TrainNum: nums,
				Date:     date,
				Carriage: i,
				Row:      j,
				Which:    "A",
				Rank:     rank,
			}
			for tmp > 0 {
				ticket := freeTicket
				model.DB.Create(&ticket)
				freeTicket.Which = nextWhich(freeTicket.Which)
				tmp--
				if freeTicket.Which == "" {
					break
				}
			}
			if tmp == 0 {
				break
			}
		}
	}
}

func nextWhich(Which interface{}) string {
	switch Which {
	case "A":
		return "B"
	case "B":
		return "C"
	case "C":
		return "E"
	case "E":
		return "F"
	case "F":
		return ""
	}
	return ""
}

// Register 用户注册
func (service *TrainInputService) Input() serializer.Response {
	train := model.Train{
		TrainNum:  service.TrainNum,
		From:      service.From,
		To:        service.To,
		StartTime: service.StartTime,
		EndTime:   service.EndTime,
		Duration:  service.Duration,
		Date:      service.Date,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 创建用户
	if err := model.DB.Create(&train).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}

	return serializer.BuildTrainResponse(train)
}
