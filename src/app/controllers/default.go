package controllers

import (
	"app/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
	_"fmt"
	_"math"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	r := generateSchedule()

	// assign data to show on template
	c.Data["ration"] = r
	c.Data["menu"] = "home"

	// render template
	c.TplName = "index.html"
}

func generateSchedule() map[string]map[int]models.Ration {
	// initialize orm
	o := orm.NewOrm()

	// declare model
	var records []models.Ration

	var minCaloryPerDay ,minWaterPerDay int = 2500 ,2;

	var scheduledRation = make(map[string]map[int]models.Ration);

	var nextScheduleKey int = 0;

	// fetch all packets records
	o.QueryTable("ration").OrderBy("expiry_date").All(&records)

	// choose packets
	choosePackets(records,minCaloryPerDay,scheduledRation,nextScheduleKey,minWaterPerDay)
	return scheduledRation
}

func choosePackets(records []models.Ration,minCaloryPerDay int,scheduledRation map[string]map[int]models.Ration,nextScheduleKey int,minWaterPerDay int) map[string]map[int]models.Ration {
	var rationPerDay = make(map[int]models.Ration);
	var coloryPerDay,packetOnIndex,waterQuantity,waterQuantityOnIndex,counter int = 0,0,0,0,0;

	for _,val := range records {
		if val.Calories > 0 {
			// Check expiry date
			diff := int(val.ExpiryDate.Sub(time.Now()).Hours()/24);
			if  diff > -1 {
				// assign packet value
				packetOnIndex = findNearestCaloryPacket(records,val.ExpiryDate,minCaloryPerDay)
				rationPerDay[counter] = records[packetOnIndex]
				coloryPerDay += records[packetOnIndex].Calories
				records = RemoveIndex(records,packetOnIndex)
				counter++
			}
			if coloryPerDay >= minCaloryPerDay {
				break;
			}
		}
	}

	for _,val := range records {
		if val.Quantity > 0 {
			// assign packet value
			waterQuantityOnIndex = findNearestWaterQuantity(records,val.Quantity,minWaterPerDay)
			rationPerDay[counter] = records[waterQuantityOnIndex]
			waterQuantity += records[waterQuantityOnIndex].Quantity
			records = RemoveIndex(records,waterQuantityOnIndex)
			counter++

			if waterQuantity >= minWaterPerDay {
				break;
			}
		}
	}

	if coloryPerDay >= minCaloryPerDay && waterQuantity >= minWaterPerDay {
		scheduledRation[time.Now().AddDate(0,0,nextScheduleKey).Format("2006-01-02")] = rationPerDay
		nextScheduleKey++
		choosePackets(records,minCaloryPerDay,scheduledRation,nextScheduleKey,minWaterPerDay)
	}
	return scheduledRation
}

func findNearestCaloryPacket(records []models.Ration,expiryDate time.Time,minCaloryPerDay int) int {
	var distance,cdistance float64 = 0,0;
	distance = float64(expiryDate.Sub(time.Now()).Hours()/24)
	var idx int = 0;
	for key,val := range records {
		if val.Calories > 0 {
		    cdistance = float64(val.ExpiryDate.Sub(time.Now()).Hours()/24)
		    if(cdistance <= distance){
		        idx = key;
		        distance = cdistance;
		    }
		}
	}
	return int(idx);
}

func findNearestWaterQuantity(records []models.Ration,waterQuantity int,minWaterPerDay int) int {
	var distance,cdistance int = 0,0;
	if waterQuantity > minWaterPerDay {
		waterQuantity = minWaterPerDay
	}
	distance = minWaterPerDay - waterQuantity
	var idx int = 0;
	for key,val := range records {
		if val.Quantity > 0 {
		    cdistance = minWaterPerDay - val.Quantity
		    if(cdistance >= distance){
		        idx = key;
		        distance = cdistance;
		    }
		}
	}
	return int(idx);
}

func RemoveIndex(p []models.Ration, index int) []models.Ration {
    return append(p[:index], p[index+1:]...)
}

