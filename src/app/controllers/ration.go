package controllers

import (
	"app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
	"fmt"
	"strconv"
)

func init() {
    // register model
    orm.RegisterModel(new(models.Ration))
}

// RationController operations for Ration
type RationController struct {
	beego.Controller
}

// URLMapping ...
func (c *RationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
	c.Mapping("inventory", c.List)
	c.Mapping("Add", c.Add)
	c.Data["menu"] = "ration"
}

// GET ...
// @Title Ration Form
// @Description create Ration
func (c *RationController) Add() {
	c.TplName = "ration/add.html"
	beego.ReadFromRequest(&c.Controller)
}

// GET
// @Title Show inventory
// @Param null
func (c *RationController) List() {
	o := orm.NewOrm()
	var ration []models.Ration
	o.QueryTable("ration").All(&ration)
	c.Data["ration"] = ration
	c.TplName = "ration/list.html"
}

// Post ...
// @Title Create
// @Description create Ration
// @Param	body		body 	models.Ration	true		"body for Ration content"
// @Success 201 {object} models.Ration
// @Failure 403 body is empty
func (c *RationController) Post() {
	o := orm.NewOrm()

	ration := models.Ration{}
	flash := beego.NewFlash()
	
	// Get form value.
	packet_type := c.GetString("packet_type")
	packet_content := c.GetString("packet_content")
	calories, _ := c.GetInt("calories")
	expiry_date,_ := time.Parse("2006-01-02", c.GetString("expiry_date"))
	quantity, _ := c.GetInt("quantity")

	ration.PacketType = packet_type
	ration.PacketContent = packet_content
	ration.Calories = calories
	ration.ExpiryDate = expiry_date
	ration.Quantity = quantity

	valid := validation.Validation{}
	if packet_type == "food" {
	    valid.Required(ration.PacketType, "PacketType")
	    valid.Required(ration.PacketContent, "Packet Content")
	    valid.Required(ration.Calories, "Calories")
	    valid.Required(ration.ExpiryDate, "ExpiryDate")
	}else{
		valid.Required(ration.Quantity, "Quantity")
	}

	// validate form
    b, _ := valid.Valid(&ration)
 
    if !b {
        // validation does not pass
        var errs string = "";
        for _, err := range valid.Errors {
            errs += err.Key+" "+err.Message
        }
        flash.Error(errs)
        flash.Store(&c.Controller)
        c.Redirect("/ration/add", 302)
        return
    }else{
		// insert
	    id, err := o.Insert(&ration)
	    fmt.Printf("ID: %d, ERR: %v\n", id, err)

		if id!=0 {
			c.Redirect("/inventory", 302)
			return
		}
    }
}

// Delete ...
// @Title Delete
// @Description delete the Ration
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
func (c *RationController) Delete() {
	o := orm.NewOrm()
	id,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	ration := models.Ration{Id:id}
	// delete
    uId, _ := o.Delete(&ration)
	if uId!=0 {
		c.Redirect("/inventory", 302)
		return
	}
}
