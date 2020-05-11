package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"sync"
)

type ExecObject struct {
	db     *gorm.DB
	entity interface{}
	method string
}

func GormMultiExec(obj ...ExecObject) error {
	wg := sync.WaitGroup{}
	wg.Add(len(obj))

	var multiErrs []string
	for i := range obj {
		go func(index int, group *sync.WaitGroup) {
			if err := gormExec(obj[index].db, obj[index].entity, obj[index].method); err != nil {
				multiErrs = append(multiErrs, fmt.Sprintf("GormMultiExec: exec error at index %v, err=%#v",
					index, err))
			}
			group.Done()
		}(i, &wg)
	}
	wg.Wait()
	if len(multiErrs) == 0 {
		return nil
	}
	return fmt.Errorf("%v", strings.Join(multiErrs, "\n"))
}

func gormExec(db *gorm.DB, entity interface{}, method string) error {
	switch method {
	case "create", "Create":
		{
			return db.Create(entity).Error
		}
	case "find", "Find":
		{
			return db.Find(entity).Error
		}
	case "first", "First":
		{
			return db.First(entity).Error
		}
	case "save", "Save":
		{
			return db.Save(entity).Error
		}
	case "update", "Update":
		{
			return db.Update(entity).Error
		}
	case "delete", "Delete":
		{
			return db.Delete(entity).Error
		}
	default:
		return fmt.Errorf("gormExec: unsupported method")
	}
}
