// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPeopleLogs(t *testing.T) {
	t.Parallel()

	query := PeopleLogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPeopleLogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeopleLogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PeopleLogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeopleLogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeopleLogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeopleLogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PeopleLogExists(ctx, tx, o.PeopleLogId)
	if err != nil {
		t.Errorf("Unable to check if PeopleLog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PeopleLogExists to return true, but got false.")
	}
}

func testPeopleLogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	peopleLogFound, err := FindPeopleLog(ctx, tx, o.PeopleLogId)
	if err != nil {
		t.Error(err)
	}

	if peopleLogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPeopleLogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PeopleLogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPeopleLogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PeopleLogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPeopleLogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	peopleLogOne := &PeopleLog{}
	peopleLogTwo := &PeopleLog{}
	if err = randomize.Struct(seed, peopleLogOne, peopleLogDBTypes, false, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}
	if err = randomize.Struct(seed, peopleLogTwo, peopleLogDBTypes, false, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peopleLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peopleLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PeopleLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPeopleLogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	peopleLogOne := &PeopleLog{}
	peopleLogTwo := &PeopleLog{}
	if err = randomize.Struct(seed, peopleLogOne, peopleLogDBTypes, false, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}
	if err = randomize.Struct(seed, peopleLogTwo, peopleLogDBTypes, false, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peopleLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peopleLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func peopleLogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func peopleLogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PeopleLog) error {
	*o = PeopleLog{}
	return nil
}

func testPeopleLogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PeopleLog{}
	o := &PeopleLog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, peopleLogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PeopleLog object: %s", err)
	}

	AddPeopleLogHook(boil.BeforeInsertHook, peopleLogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	peopleLogBeforeInsertHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.AfterInsertHook, peopleLogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	peopleLogAfterInsertHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.AfterSelectHook, peopleLogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	peopleLogAfterSelectHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.BeforeUpdateHook, peopleLogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	peopleLogBeforeUpdateHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.AfterUpdateHook, peopleLogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	peopleLogAfterUpdateHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.BeforeDeleteHook, peopleLogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	peopleLogBeforeDeleteHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.AfterDeleteHook, peopleLogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	peopleLogAfterDeleteHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.BeforeUpsertHook, peopleLogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	peopleLogBeforeUpsertHooks = []PeopleLogHook{}

	AddPeopleLogHook(boil.AfterUpsertHook, peopleLogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	peopleLogAfterUpsertHooks = []PeopleLogHook{}
}

func testPeopleLogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeopleLogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(peopleLogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeopleLogToOnePersonUsingPeopleIdNew(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PeopleLog
	var foreign Person

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, personDBTypes, true, personColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Person struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.PeopleIdNew, foreign.PeopleId)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.PeopleIdNew().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.PeopleId, foreign.PeopleId) {
		t.Errorf("want: %v, got %v", foreign.PeopleId, check.PeopleId)
	}

	slice := PeopleLogSlice{&local}
	if err = local.L.LoadPeopleIdNew(ctx, tx, false, (*[]*PeopleLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PeopleIdNew == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PeopleIdNew = nil
	if err = local.L.LoadPeopleIdNew(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PeopleIdNew == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPeopleLogToOnePersonUsingPeopleIdOld(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PeopleLog
	var foreign Person

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, personDBTypes, true, personColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Person struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.PeopleIdOld, foreign.PeopleId)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.PeopleIdOld().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.PeopleId, foreign.PeopleId) {
		t.Errorf("want: %v, got %v", foreign.PeopleId, check.PeopleId)
	}

	slice := PeopleLogSlice{&local}
	if err = local.L.LoadPeopleIdOld(ctx, tx, false, (*[]*PeopleLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PeopleIdOld == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PeopleIdOld = nil
	if err = local.L.LoadPeopleIdOld(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PeopleIdOld == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPeopleLogToOneSetOpPersonUsingPeopleIdNew(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeopleLog
	var b, c Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peopleLogDBTypes, false, strmangle.SetComplement(peopleLogPrimaryKeyColumns, peopleLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Person{&b, &c} {
		err = a.SetPeopleIdNew(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PeopleIdNew != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PeopleIdNewPeopleLogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.PeopleIdNew, x.PeopleId) {
			t.Error("foreign key was wrong value", a.PeopleIdNew)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PeopleIdNew))
		reflect.Indirect(reflect.ValueOf(&a.PeopleIdNew)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.PeopleIdNew, x.PeopleId) {
			t.Error("foreign key was wrong value", a.PeopleIdNew, x.PeopleId)
		}
	}
}

func testPeopleLogToOneRemoveOpPersonUsingPeopleIdNew(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeopleLog
	var b Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peopleLogDBTypes, false, strmangle.SetComplement(peopleLogPrimaryKeyColumns, peopleLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPeopleIdNew(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePeopleIdNew(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.PeopleIdNew().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.PeopleIdNew != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.PeopleIdNew) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PeopleIdNewPeopleLogs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPeopleLogToOneSetOpPersonUsingPeopleIdOld(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeopleLog
	var b, c Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peopleLogDBTypes, false, strmangle.SetComplement(peopleLogPrimaryKeyColumns, peopleLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Person{&b, &c} {
		err = a.SetPeopleIdOld(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PeopleIdOld != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PeopleIdOldPeopleLogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.PeopleIdOld, x.PeopleId) {
			t.Error("foreign key was wrong value", a.PeopleIdOld)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PeopleIdOld))
		reflect.Indirect(reflect.ValueOf(&a.PeopleIdOld)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.PeopleIdOld, x.PeopleId) {
			t.Error("foreign key was wrong value", a.PeopleIdOld, x.PeopleId)
		}
	}
}

func testPeopleLogToOneRemoveOpPersonUsingPeopleIdOld(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeopleLog
	var b Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peopleLogDBTypes, false, strmangle.SetComplement(peopleLogPrimaryKeyColumns, peopleLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPeopleIdOld(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePeopleIdOld(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.PeopleIdOld().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.PeopleIdOld != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.PeopleIdOld) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PeopleIdOldPeopleLogs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPeopleLogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeopleLogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeopleLogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeopleLogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PeopleLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	peopleLogDBTypes = map[string]string{`PeopleLogId`: `INTEGER`, `PeopleIdNew`: `INTEGER`, `PeopleIdOld`: `INTEGER`, `GnNew`: `TEXT`, `GNOLD`: `TEXT`, `SnNew`: `TEXT`, `SNOLD`: `TEXT`, `BirthdayNew`: `DATE`, `BirthdayOld`: `DATE`, `Action`: `TEXT`, `Timestamp`: `DATE`}
	_                = bytes.MinRead
)

func testPeopleLogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(peopleLogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(peopleLogAllColumns) == len(peopleLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPeopleLogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(peopleLogAllColumns) == len(peopleLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PeopleLog{}
	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeopleLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peopleLogDBTypes, true, peopleLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PeopleLog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(peopleLogAllColumns, peopleLogPrimaryKeyColumns) {
		fields = peopleLogAllColumns
	} else {
		fields = strmangle.SetComplement(
			peopleLogAllColumns,
			peopleLogPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PeopleLogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
