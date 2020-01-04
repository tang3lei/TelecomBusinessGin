package view

import (
	"github.com/gin-gonic/gin"
	"time"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
	"upc.edu.cn/telecom_business/telecom_business_api/model"
	"upc.edu.cn/telecom_business/telecom_business_api/tools"
)

func ListDeal(c *gin.Context) {
	arg, _ := NewQueryArg(c)
	arg.Equal("phone_number", "phone_number")

	var res []model.Deal
	err := db.Deal().ApplyQuery(arg, &res)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	SendSucResp(c, res)
}

func MonthlyDeals(c *gin.Context) {
	beginMonth := tools.GetQueryString(c, "begin_month", "")
	endMonth := tools.GetQueryString(c, "end_month", "")
	beginTime, err := time.Parse("2006-01", beginMonth)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	endTime, err := time.Parse("2006-01", endMonth)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	var deals []model.Deal
	endTime = endTime.AddDate(0, 1, 0)
	err = db.Deal().GetByMap(map[string]interface{}{
		"create_at >= ?": beginTime,
		"create_at <= ?": endTime,
	}, &deals)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	resList := make(map[int][]model.Deal)
	for _, d := range deals {
		if d.Type > 0 {
			resList[d.Type] = append(resList[d.Type], d)
		}
	}
	SendSucResp(c, resList)
}

func DailyDeals(c *gin.Context) {
	beginDaily := tools.GetQueryString(c, "begin_daily", "")
	endDaily := tools.GetQueryString(c, "end_daily", "")
	beginTime, err := time.Parse("2006-01-02", beginDaily)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	endTime, err := time.Parse("2006-01-02", endDaily)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	var deals []model.Deal
	endTime = endTime.AddDate(0, 0, 1)
	err = db.Deal().GetByMap(map[string]interface{}{
		"create_at >= ?": beginTime,
		"create_at <= ?": endTime,
	}, &deals)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	resList := make(map[string][]model.Deal)
	for _, d := range deals {
		resList[d.JobName] = append(resList[d.JobName], d)
	}
	SendSucResp(c, resList)
}

type tempRes struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data []int  `json:"data"`
}

func UserDailyDeals(c *gin.Context) {
	beginDaily := tools.GetQueryString(c, "begin_daily", "")
	endDaily := tools.GetQueryString(c, "end_daily", "")
	beginTime, err := time.Parse("2006-01-02", beginDaily)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	endTime, err := time.Parse("2006-01-02", endDaily)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	var deals []model.Deal
	endTime = endTime.AddDate(0, 0, 1)
	err = db.Deal().GetByMap(map[string]interface{}{
		"create_at >= ?": beginTime,
		"create_at <= ?": endTime,
	}, &deals)

	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	pnMp := make(map[string][]model.Deal)
	for _, d := range deals {
		pnMp[d.PhoneNumber] = append(pnMp[d.PhoneNumber], d)
	}
	caL := []string{}
	b := beginTime
	for b.Before(endTime) {
		caL = append(caL, b.Format("2006-01-02"))
		b = b.AddDate(0, 0, 1)
	}
	var srLs []tempRes
	for number, list := range pnMp {
		sr := tempRes{
			Name: number,
			Type: "line",
			Data: nil,
		}
		tmMp := make(map[string][]model.Deal)
		for _, d := range list {
			tmMp[d.CreateAt.Format("2006-01-02")] = append(pnMp[d.CreateAt.Format("2006-01-02")], d)
		}
		b := beginTime
		for b.Before(endTime) {
			var cost int
			if ls, ok := tmMp[b.Format("2006-01-02")]; ok {
				for _, c := range ls {
					cost += int(c.Cost)
				}
			}
			sr.Data = append(sr.Data, cost)
			b = b.AddDate(0, 0, 1)
		}
		srLs = append(srLs, sr)
	}
	SendSucResp(c, map[string]interface{}{
		"ca":   caL,
		"srls": srLs,
	})
}

func UpdateDeal(c *gin.Context) {
	var deal model.Deal
	var err error
	err = c.BindJSON(&deal)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}

	if deal.Id == 0 {
		deal.CreateAt = time.Now()
		deal.UpdateAt = deal.CreateAt
		err = db.Deal().Create(&deal)
	} else {
		err = db.Deal().Update(&deal)
	}
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}

	SendSucResp(c, deal)
}
