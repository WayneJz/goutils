package orms

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/waynejz/goutils/errors"
)

// ExecObject ... Gorm execution object, DB must not have instant-execute method,
// Entity should be an addressable value, just as gorm required
type ExecObject struct {
	DB     *gorm.DB
	Entity interface{}
	Method string
}

// MultiExec ... Concurrently perform database manipulation, with timeout killed
func MultiExec(timeout time.Duration, obj ...ExecObject) error {
	multiErr := new(errors.MultiError)
	doneChannel := make(chan int, len(obj))

	for i := range obj {
		go func(index int) {
			if err := gormExec(obj[index].DB, obj[index].Entity, obj[index].Method); err != nil {
				multiErr.Add(err)
			}
			doneChannel <- index
		}(i)
	}

	// Over such timeout, this function killed
	timer := time.NewTimer(timeout)
	receivedDone := 0

	for {
		select {
		case <-doneChannel:
			receivedDone++
			if receivedDone == len(obj) {
				close(doneChannel)
				return multiErr.Return()
			}
		case <-timer.C:
			close(doneChannel)
			if receivedDone != len(obj) {
				multiErr.Add(fmt.Errorf("orm multi execution: timeout %v exceeded, killed", timeout))
			}
			return multiErr.Return()
		}
	}
}

// Execute gorm method, now supports for create, find, first, save, update, delete
func gormExec(db *gorm.DB, entity interface{}, method string) error {
	switch method {
	case "create", "Create":
		return db.Create(entity).Error
	case "find", "Find":
		return db.Find(entity).Error
	case "first", "First":
		return db.First(entity).Error
	case "save", "Save":
		return db.Save(entity).Error
	case "update", "Update":
		return db.Update(entity).Error
	case "delete", "Delete":
		return db.Delete(entity).Error
	default:
		return fmt.Errorf("gorm exec: unsupported method %v", method)
	}
}
