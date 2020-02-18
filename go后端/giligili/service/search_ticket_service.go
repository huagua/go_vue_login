package service

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"github.com/gin-gonic/gin"
)

// UserRegisterService 管理用户注册服务
type SearchTicketService struct {
	From            string `form:"From" json:"From"`
	To		        string `form:"To" json:"To"`
	Date		    string `form:"Date" json:"Date"`
}



// Login 用户登录函数
func (service *SearchTicketService) Search(c *gin.Context) serializer.Response {
	//var trains []model.Train
	//
	//model.DB.Where(&model.Train{From:service.From, To:service.To, Date:service.Date}).Find(&trains)
	//
	//if  len(trains) == 0 {
	//	return serializer.ParamErr("无符合车次", nil)
	//}
	//
	//return serializer.BuildSearchResponse(trains[:])
	var tickets []model.SearchTicket

	trains,_ := cache.RedisClient.LRange(cache.GetSearchKey(service.From, service.To),0, -1).Result()
	for _,train := range trains{
		super,_ := cache.RedisClient.HMGet(cache.GetTicketKey(service.Date, train, "super"), "left").Result()
		first,_ := cache.RedisClient.HMGet(cache.GetTicketKey(service.Date, train, "first"), "left").Result()
		second,_ := cache.RedisClient.HMGet(cache.GetTicketKey(service.Date, train, "second"), "left").Result()
		startTime,_ := cache.RedisClient.HMGet(cache.GetTrainKey(train, service.Date), "startTime").Result()
		duration,_ := cache.RedisClient.HMGet(cache.GetTrainKey(train, service.Date), "duration").Result()
		var ticket = model.SearchTicket{
			From:      service.From,
			To:        service.To,
			TrainNum:  train,
			Super:     super,
			First:     first,
			Second:    second,
			StartTime: startTime,
			Duration:  duration,
		}

		tickets = append(tickets, ticket)
	}

	return serializer.BuildSearchResponse(tickets)

	//trains, _:=cache.RedisClient.Keys(cache.GetSearchKey(service.From, service.To)).Result()
	//for _,train := range trains{
	//	tmp := strings.Split(train, ":")
	//	super,_ := cache.RedisClient.HMGet(train, "left").Result()
	//	first,_ := cache.RedisClient.HMGet(train, "left").Result()
	//	second,_ := cache.RedisClient.HMGet(train, "left").Result()
	//
	//	var ticket = model.SearchTicket{
	//		From:     tmp[1],
	//		To:       tmp[2],
	//		TrainNum: tmp[4],
	//		Super:    super,
	//		First:    first,
	//		Second:   second,
	//	}
	//
	//	tickets = append(tickets, ticket)
	//}
	//
	//return serializer.BuildSearchResponse(tickets)
}
