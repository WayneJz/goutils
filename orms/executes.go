package orms

import (
	"fmt"
	"time"

	"github.com/WayneJz/goutils/errors"
	"github.com/jinzhu/gorm"
)

// ExecObject ... Gorm execution object:
// DB must not have instant-execution method, but can hold condition clauses beforehand,
// Entity should be an addressable value, just as gorm required,
// Where field is optional,
// Method is the final (instant-execution) method
type ExecObject struct {
	DB     *gorm.DB
	Entity interface{}
	Where  []interface{}
	Method string
}

// MultiExec ... Concurrently perform database manipulation, with timeout killed
func Executes(timeout time.Duration, obj ...ExecObject) error {
	errs := new(errors.Errors)
	doneChannel := make(chan int, len(obj))

	for index := range obj {
		go func(i int) {
			if err := exec(obj[i].DB, obj[i].Entity, obj[i].Where, obj[i].Method); err != nil {
				errs.Add(err)
			}
			doneChannel <- i
		}(index)
	}

	// Over such timeout, this function killed
	timer := time.NewTimer(timeout)
	receivedDone := 0

	// Waiting for all tasks done or timeout killed
	for {
		select {
		case <-doneChannel:
			receivedDone++
			if receivedDone == len(obj) {
				close(doneChannel)
				return errs.Return()
			}
		case <-timer.C:
			close(doneChannel)
			if receivedDone != len(obj) {
				errs.Add(fmt.Errorf("orm multi execution: timeout %v exceeded, killed", timeout))
			}
			return errs.Return()
		}
	}
}

// exec ... Execute gorm method, now supports for create, find, first, save, update, delete
func exec(db *gorm.DB, entity interface{}, where []interface{}, method string) error {
	if db == nil || entity == nil {
		return fmt.Errorf("gorm exec: passing nil database pointer or nil entity")
	}
	switch method {
	case "create", "Create":
		return db.Create(entity).Error
	case "find", "Find":
		return db.Find(entity, where...).Error
	case "first", "First":
		return db.First(entity, where...).Error
	case "save", "Save":
		return db.Save(entity).Error
	case "update", "Update":
		return db.Update(entity).Error
	case "delete", "Delete":
		return db.Delete(entity, where...).Error
	default:
		return fmt.Errorf("gorm exec: unsupported method %v", method)
	}
}
