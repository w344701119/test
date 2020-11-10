package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ZtProject struct {
	Id              int       `orm:"column(id);auto" description:"ID"`
	IsCat           string    `orm:"column(isCat)"`
	CatID           uint32    `orm:"column(catID)"`
	Type            string    `orm:"column(type);size(20)" description:"类型"`
	Parent          uint32    `orm:"column(parent)" description:"父ID"`
	Name            string    `orm:"column(name);size(90)" description:"名称"`
	Code            string    `orm:"column(code);size(45)" description:"编号"`
	Begin           time.Time `orm:"column(begin);type(date)" description:"开始"`
	End             time.Time `orm:"column(end);type(date)" description:"结束"`
	Days            uint16    `orm:"column(days)" description:"工期"`
	Status          string    `orm:"column(status);size(10)" description:"状态"`
	Statge          string    `orm:"column(statge)" description:"阶段"`
	Pri             string    `orm:"column(pri)" description:"优先级"`
	Activity        string    `orm:"column(activity)" description:"项目活跃程度"`
	Important       int8      `orm:"column(important)" description:"是否是重点项目"`
	Desc            string    `orm:"column(desc)" description:"描述"`
	OpenedBy        string    `orm:"column(openedBy);size(30)" description:"创建人"`
	OpenedDate      time.Time `orm:"column(openedDate);type(datetime)" description:"创建日期"`
	OpenedVersion   string    `orm:"column(openedVersion);size(20)" description:"创建版本号"`
	ClosedBy        string    `orm:"column(closedBy);size(30)" description:"关闭人"`
	ClosedDate      time.Time `orm:"column(closedDate);type(datetime)" description:"关闭日期"`
	CanceledBy      string    `orm:"column(canceledBy);size(30)" description:"取消人"`
	CanceledDate    time.Time `orm:"column(canceledDate);type(datetime)" description:"取消日期"`
	PO              string    `orm:"column(PO);size(30)" description:"产品负责人"`
	PM              string    `orm:"column(PM);size(30)" description:"项目负责人"`
	QD              string    `orm:"column(QD);size(30)" description:"测试负责人"`
	RD              string    `orm:"column(RD);size(30)" description:"发布负责人"`
	DD              string    `orm:"column(DD);size(30)" description:"设计负责人"`
	TD              string    `orm:"column(TD);size(30)" description:"终端开发负责人"`
	MD              string    `orm:"column(MD);size(30)" description:"后台开发负责人"`
	Team            string    `orm:"column(team);size(90)" description:"团队"`
	Acl             string    `orm:"column(acl)" description:"访问控制:open公开,private私有,custom:自定义"`
	Whitelist       string    `orm:"column(whitelist)" description:"白名单"`
	Attribute       string    `orm:"column(attribute);size(20)" description:"项目属性"`
	Order           uint32    `orm:"column(order)" description:"排序"`
	Deleted         string    `orm:"column(deleted)" description:"是否删除"`
	Genre           int8      `orm:"column(genre)" description:"项目类型"`
	Stage           string    `orm:"column(stage);size(20)" description:"项目阶段"`
	StageDate       time.Time `orm:"column(stageDate);type(date)" description:"阶段截止时间"`
	Remark          string    `orm:"column(remark);null" description:"备注"`
	Img             string    `orm:"column(img);size(255)" description:"图片"`
	ScreenShot      int8      `orm:"column(screenShot)" description:"甘特图是否截图"`
	Schedule        int8      `orm:"column(schedule)" description:"是否排期0:否,1:是"`
	TestWorkload    float32   `orm:"column(testWorkload)" description:"测试计划工作量"`
	StoryBegin      time.Time `orm:"column(storyBegin);type(date)" description:"需求整理计划结开始时间"`
	StoryEnd        time.Time `orm:"column(storyEnd);type(date)" description:"需求整理计划结束时间"`
	DevBegin        time.Time `orm:"column(devBegin);type(date)" description:"开发计划开始时间"`
	DevEnd          time.Time `orm:"column(devEnd);type(date)" description:"开发计划结束时间"`
	DebugBegin      time.Time `orm:"column(debugBegin);type(date)" description:"联调计划开始时间"`
	DebugEnd        time.Time `orm:"column(debugEnd);type(date)" description:"联调计划结束时间"`
	AcceptanceBegin time.Time `orm:"column(acceptanceBegin);type(date)" description:"验收计划开始时间"`
	AcceptanceEnd   time.Time `orm:"column(acceptanceEnd);type(date)" description:"验收计划结束时间"`
	TestBegin       time.Time `orm:"column(testBegin);type(date)" description:"测试计划开始时间"`
	TestEnd         time.Time `orm:"column(testEnd);type(date)" description:"测试计划结束时间"`
	SubmitTime      time.Time `orm:"column(submitTime);type(datetime)" description:"提测时间"`
	TestFinish      time.Time `orm:"column(testFinish);type(datetime)" description:"完成测试时间"`
}

func (t *ZtProject) TableName() string {
	return "zt_project"
}

func init() {
	orm.RegisterModel(new(ZtProject))
}

// AddZtProject insert a new ZtProject into database and returns
// last inserted Id on success.
func AddZtProject(m *ZtProject) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZtProjectById retrieves ZtProject by Id. Returns error if
// Id doesn't exist
func GetZtProjectById(id int) (v *ZtProject, err error) {
	o := orm.NewOrm()
	v = &ZtProject{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllZtProject retrieves all ZtProject matches certain condition. Returns empty list if
// no records exist
func GetAllZtProject(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZtProject))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ZtProject
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateZtProject updates ZtProject by Id and returns error if
// the record to be updated doesn't exist
func UpdateZtProjectById(m *ZtProject) (err error) {
	o := orm.NewOrm()
	v := ZtProject{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZtProject deletes ZtProject by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZtProject(id int) (err error) {
	o := orm.NewOrm()
	v := ZtProject{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZtProject{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
