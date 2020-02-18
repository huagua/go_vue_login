package serializer

import (
	"giligili/model"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// Train 车次序列化器
type Train struct {
	ID        		uint   `json:"id"`
	TrainNum        string `json:"train_num"`
	From            string `json:"from"`
	To		        string `json:"to"`
	StartTime		string `json:"start_time"`
	EndTime	        string `json:"end_time"`
	Duration        string `json:"duration"`
	Date		    string `json:"date"`
	Super			uint `json:"super"`
	First	        uint `json:"first"`
	Second          uint `json:"second"`
}

// Train 车次序列化器
type SearchTrain struct {
	ID        		int   `json:"id"`
	TrainNum        string `json:"train_num"`
	From            string `json:"from"`
	To		        string `json:"to"`
	StartTime		interface{} `json:"start_time"`
	Duration        interface{} `json:"duration"`
	Super			interface{} `json:"super"`
	First	        interface{} `json:"first"`
	Second          interface{} `json:"second"`
	HasChildren     bool `json:"hasChildren"`
}

type Ticket struct {
	Carriage int `json:"carriage"`
	Row      int `json:"row"`
	Which    string `json:"which"`
	Money    int `json:"money"`
}

// BuildUser 序列化用户
func BuildTicket(ticket model.Ticket) Ticket {
	return Ticket{
		Carriage: ticket.Carriage,
		Row:      ticket.Row,
		Which:    ticket.Which,
		Money:    ticket.Money,
	}
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUser 序列化用户
func BuildTrain(train model.Train) Train {
	return Train{
		ID:        train.ID,
		TrainNum:       train.TrainNum,
		From:          train.From,
		To:             train.To,
		StartTime:      train.StartTime,
		EndTime:        train.EndTime,
		Duration:      train.Duration,
		Date:           train.Date,
	}
}

func BuildSearchTrain(ticket model.SearchTicket, id int) SearchTrain {
	return SearchTrain{
		ID:        id,
		TrainNum:       ticket.TrainNum,
		From:          ticket.From,
		To:             ticket.To,
		StartTime:      ticket.StartTime,
		Duration:      ticket.Duration,
		Super:          ticket.Super,
		First:          ticket.First,
		Second:         ticket.Second,
		HasChildren:    true,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

// BuildUserResponse 序列化车次信息响应
func BuildTrainResponse(train model.Train) Response {
	return Response{
		Data: BuildTrain(train),
	}
}

// BuildUserResponse 序列化车次信息响应
func BuildSearchResponse(tickets []model.SearchTicket) Response {
	var data []SearchTrain

	for id,ticket := range tickets{
		data = append(data, BuildSearchTrain(ticket, id))
	}

	return Response{
		Data: data,
	}
}
