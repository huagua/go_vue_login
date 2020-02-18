package cache

import (
	"fmt"
)

const (
	User = "user"
)

func GetTicketKey(date string,trainNum string, rank string) string{
	return fmt.Sprintf("tickets:%s:%s:%s", date, trainNum, rank)
}

func GetTrainKey(trainNum string, date string) string{
	return fmt.Sprintf("trains:%s:%s", trainNum, date)
}

func GetSearchKey(from string, to string) string{
	return fmt.Sprintf("tickets:%s:%s", from, to)
}



