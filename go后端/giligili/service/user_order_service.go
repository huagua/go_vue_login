package service

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserLoginService 管理用户登录的服务
type UserOrderService struct {
	TrainNum string `form:"trainNum" json:"trainNum"`
	Name string `form:"name" json:"name"`
	IdNum string `form:"idNum" json:"idNum"`
	Date string `form:"date1" json:"date1"`
	Rank string `form:"rank" json:"rank"`
	From string `form:"from" json:"from"`
	To string `form:"to" json:"to"`
}

// Login 用户登录函数
func (service *UserOrderService) Order(c *gin.Context) serializer.Response {
	var train model.Train
	var ticket model.FreeTicket

	//从缓存中寻找是否有该用户
	if cache.RedisClient.SIsMember("user", service.Name).Val() == false {
		return serializer.ParamErr("无该用户", nil)
	}

	//查看缓存中是否有该车次
	if res,_ :=cache.RedisClient.Exists(cache.GetTrainKey(service.TrainNum, service.Date)).Result(); res == 0 {
		return serializer.ParamErr("无该车次", nil)
	}

	//在缓存中查看该车次的已订票用户是否包含该用户
	if exist,_ := cache.RedisClient.HExists(cache.GetTrainKey(service.TrainNum, service.Date), service.Name).Result(); exist == true {
		return serializer.ParamErr("该用户已购票", nil)
	}

	//在缓存中查取余票信息
	num,_:= cache.RedisClient.HGet(cache.GetTicketKey(service.Date, service.TrainNum, service.Rank), "left").Result()
	if num == "0" {
		return serializer.ParamErr("该类车票已售罄", nil)
	}

	//将缓存中的余票减去一张
	left,_ := strconv.Atoi(num)
	cache.RedisClient.HSet(cache.GetTicketKey(service.Date, service.TrainNum, service.Rank), "left", left-1)

	//hash表记录已经购票的用户
	cache.RedisClient.HSet(cache.GetTrainKey(service.TrainNum, service.Date), service.Name, service.Rank)

	//从数据库中的空闲车票中找到一张空闲票，并从空闲表中删除
	model.DB.Where(&model.FreeTicket{TrainNum:service.TrainNum, Date:service.Date, Rank:service.Rank}).First(&ticket)
	//数据库中的剩余票减去一张
	model.DB.Delete(&ticket)

	return serializer.BuildTrainResponse(train)
}
