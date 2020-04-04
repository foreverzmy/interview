// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: answer.proto

/*
Package answer is a generated protocol buffer package.

It is generated from these files:
	answer.proto

It has these top-level messages:
	Empty
	Answer
	AnswerList
	CreateAnswerResponse
	GetAnswerRequest
	GetAnswerListRequest
*/
package answer

import context "context"
import errors "errors"
import time "time"

import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import ptypes1 "github.com/golang/protobuf/ptypes"

import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type AnswerORM struct {
	Content   string    `gorm:"default:'';not null"`
	CreatedAt time.Time `gorm:"not null"`
	Id        int64     `gorm:"primary_key;auto_increment"`
	QuId      int64     `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (AnswerORM) TableName() string {
	return "answer"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Answer) ToORM(ctx context.Context) (AnswerORM, error) {
	to := AnswerORM{}
	var err error
	if prehook, ok := interface{}(m).(AnswerWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = t
	}
	to.QuId = m.QuId
	to.Content = m.Content
	if posthook, ok := interface{}(m).(AnswerWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AnswerORM) ToPB(ctx context.Context) (Answer, error) {
	to := Answer{}
	var err error
	if prehook, ok := interface{}(m).(AnswerWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if to.CreatedAt, err = ptypes1.TimestampProto(m.CreatedAt); err != nil {
		return to, err
	}
	if to.UpdatedAt, err = ptypes1.TimestampProto(m.UpdatedAt); err != nil {
		return to, err
	}
	to.QuId = m.QuId
	to.Content = m.Content
	if posthook, ok := interface{}(m).(AnswerWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Answer the arg will be the target, the caller the one being converted from

// AnswerBeforeToORM called before default ToORM code
type AnswerWithBeforeToORM interface {
	BeforeToORM(context.Context, *AnswerORM) error
}

// AnswerAfterToORM called after default ToORM code
type AnswerWithAfterToORM interface {
	AfterToORM(context.Context, *AnswerORM) error
}

// AnswerBeforeToPB called before default ToPB code
type AnswerWithBeforeToPB interface {
	BeforeToPB(context.Context, *Answer) error
}

// AnswerAfterToPB called after default ToPB code
type AnswerWithAfterToPB interface {
	AfterToPB(context.Context, *Answer) error
}

// DefaultCreateAnswer executes a basic gorm create call
func DefaultCreateAnswer(ctx context.Context, in *Answer, db *gorm1.DB) (*Answer, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateAnswer")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AnswerORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadAnswer executes a basic gorm read call
func DefaultReadAnswer(ctx context.Context, in *Answer, db *gorm1.DB) (*Answer, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadAnswer")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.New("DefaultReadAnswer requires a non-zero primary key")
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &AnswerORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := AnswerORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(AnswerORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type AnswerORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteAnswer(ctx context.Context, in *Answer, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteAnswer")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&AnswerORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type AnswerORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteAnswerSet(ctx context.Context, in []*Answer, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteAnswerSet")
	}
	var err error
	keys := []int64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.New("A non-zero ID value is required for a delete call")
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&AnswerORM{})).(AnswerORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&AnswerORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&AnswerORM{})).(AnswerORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type AnswerORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Answer, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Answer, *gorm1.DB) error
}

// DefaultStrictUpdateAnswer clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateAnswer(ctx context.Context, in *Answer, db *gorm1.DB) (*Answer, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateAnswer")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &AnswerORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type AnswerORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchAnswer executes a basic gorm update call with patch behavior
func DefaultPatchAnswer(ctx context.Context, in *Answer, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Answer, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultPatchAnswer")
	}
	var pbObj Answer
	var err error
	if hook, ok := interface{}(&pbObj).(AnswerWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadAnswer(ctx, &Answer{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(AnswerWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskAnswer(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(AnswerWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateAnswer(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(AnswerWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type AnswerWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Answer, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Answer, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Answer, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Answer, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultApplyFieldMaskAnswer patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAnswer(ctx context.Context, patchee *Answer, patcher *Answer, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Answer, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.New("Patchee inputs to DefaultApplyFieldMaskAnswer must be non-nil")
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"UpdatedAt" {
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if f == prefix+"QuId" {
			patchee.QuId = patcher.QuId
			continue
		}
		if f == prefix+"Content" {
			patchee.Content = patcher.Content
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAnswer executes a gorm list call
func DefaultListAnswer(ctx context.Context, db *gorm1.DB) ([]*Answer, error) {
	in := Answer{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &AnswerORM{}, &Answer{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []AnswerORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AnswerORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Answer{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AnswerORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AnswerORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]AnswerORM) error
}