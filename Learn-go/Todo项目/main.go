package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type Base struct{
// 	ID uint64 `json:"id" gorm:"primary_key"`
// 	CreatedAt time.Time `json:"createdAt" gorm:"DEFAULT CURRENT_TIMESTAMP"`
// 	UpdateAt time.Time `json:"updateAt" gorm:"DEFAULT CURRENT_TIMESTAMP ON UPDATE      CURRENT_TIMESTAMP"`
// }
type UserTodo struct {
	gorm.Model
	UserID              string    `gorm:"userid;size:50;not null;"`        //用户ID
	UserTodoName        string    `gorm:"usertodoname;size:50;not null;"`  //用户Todo Name
	UserTodoTitle       string    `gorm:"usertodotitle;size:50;not null;"` //用户Todo的标题
	UserTodoDescription string    `gorm:"description;size:200;not null;"`  //用户Todo的描述
	UserTodoDueTime     time.Time `gorm:"duetime"`                         //用户Todo截止时间
	UserTodoRemindTime  time.Time `gorm:"remindtime"`                      //用户Todo提前多久通知
	Status              bool      `gorm: "status"`                         //用户Todo状态
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:sds123@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接有问题！")
	}
	db.AutoMigrate(&UserTodo{})

}

//增加一个Todo
func addTodo(c *gin.Context) {
	loc, _ := time.LoadLocation("Europe/Berlin")
	duetime, _ := time.ParseInLocation("2000-Jan-01", c.PostForm("duetime"), loc)
	remindtime, _ := time.ParseInLocation("2000-Jan-01", c.PostForm("remindtime"), loc)
	fmt.Print(remindtime)
	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		panic("转换错误！")
	}
	todo := UserTodo{UserID: c.PostForm("userid"), UserTodoName: c.PostForm("usertodoname"), UserTodoTitle: c.PostForm("usertodotitle"),
		UserTodoDescription: c.PostForm("description"), UserTodoDueTime: duetime, UserTodoRemindTime: remindtime, Status: status}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!",
		"resourceId": todo.ID})

}

//取出所有Todo
func fetchAllTodo(c *gin.Context) {
	var todos []UserTodo
	var _todos []UserTodo
	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"stataus": http.StatusNotFound, "message": "未发现数据，请先添加数据！"})
		return
	}
	for _, item := range todos {
		_todos = append(_todos, UserTodo{Model: gorm.Model{ID: item.ID, CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt},
			UserID: item.UserID, UserTodoName: item.UserTodoName, UserTodoTitle: item.UserTodoTitle,
			UserTodoDescription: item.UserTodoDescription, UserTodoDueTime: item.UserTodoDueTime,
			UserTodoRemindTime: item.UserTodoRemindTime, Status: item.Status})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

//取出一条数据
func fetchOneTodo(c *gin.Context) {
	var todo UserTodo
	//var _todo UserTodo
	todoid := c.Param("userid")
	db.Where("user_id = ?", todoid).First(&todo)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "没有数据，取不出来！！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})

}

//更新一个todo
func updateTodo(c *gin.Context) {
	var todo UserTodo
	todoid := c.Param("userid")
	db.Where("user_id = ?", todoid).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "没有数据，请添加后在更新！！"})
		return
	}
	duetime, _ := time.Parse("2000-Jan-01", c.PostForm("duetime"))
	remindtime, _ := time.Parse("2000-Jan-01", c.PostForm("remindtime"))
	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		panic("转换错误！")
	}
	db.Model(&todo).Update(map[string]interface{}{"userid": c.PostForm("userid"), "usertodoname": c.PostForm("usertodoname"),
		"usertodotitle": c.PostForm("usertodotitle"), "description": c.PostForm("description"), "duetime": duetime,
		"remindtime": remindtime, "status": status})
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "更新成功！！"})
}

//删除一个todo
func deleteTodo(c *gin.Context) {
	var todo UserTodo
	todoid := c.Param("userid")
	db.Where("user_id = ?", todoid).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "没有数据，请添加后再删除！！"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "删除成功！！"})
}

func main() {
	router := gin.Default()
	//创建路由组
	v := router.Group("/api/todos")
	v.POST("/", addTodo)
	v.GET("/", fetchAllTodo)
	v.GET("/:userid", fetchOneTodo)
	v.PUT("/:userid", updateTodo)
	v.DELETE("/:userid", deleteTodo)

	router.Run(":8080")
	defer db.Close()

}
