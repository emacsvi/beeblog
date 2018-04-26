package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	// register driver
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// set default database
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}
type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Category         string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	ReplayLastUserId int64
}

type Comment struct {
	Id        int64
	Tid       int64
	Name      string
	Content   string    `orm:"size(1000)"`
	ReplyTime time.Time `orm:"index"`
}

func AddCategory(name string) error {
	beego.Error(name)
	o := orm.NewOrm()
	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func AddTopic(title, category, content string) error {
	o := orm.NewOrm()
	t := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(t)
	return err
}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if cate != "" {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-Created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidN, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidN).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, category, content string) error {
	tidN, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	t := &Topic{Id: tidN}
	if o.Read(t) == nil {
		t.Title = title
		t.Category = category
		t.Content = content
		t.Updated = time.Now()
		_, err = o.Update(t)
		return err
	}

	return nil
}

func DelTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	t := &Topic{Id: tid}
	_, err = o.Delete(t)
	return err
}

func AddComment(tid, name, content string) error {
	tidN, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	c := &Comment{
		Tid:       tidN,
		Name:      name,
		Content:   content,
		ReplyTime: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(c)
	return err
}

func GetAllComments() ([]*Comment, error) {
	o := orm.NewOrm()
	c := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err := qs.OrderBy("-ReplyTime").All(&c)
	return c, err
}
func DelComment(id string) error {
	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	c := &Comment{Id: rid}
	_, err = o.Delete(c)
	return err
}
