// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userToRoleFieldNames          = builder.RawFieldNames(&UserToRole{})
	userToRoleRows                = strings.Join(userToRoleFieldNames, ",")
	userToRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(userToRoleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userToRoleInsertFeilds   = strings.Join(stringx.Remove(userToRoleFieldNames, "`id`", "`create_time`", "`is_delete`"), ",")
)

type (
	userToRoleModel interface {
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindRoleUUIDArrByUserUuid(ctx context.Context, userUuid string) (*[]string, error)
		CountUserUUIDArrByRoleUuid(ctx context.Context, roleUUID string) (int64, error)
		FindUserUUIDArrByRoleUuid(ctx context.Context,  args *FindUserUUIDArrByRoleUuidArgs) (*[]string, error)
		CountUserGroupByRoleUuid(ctx context.Context, roleUUIDArray *[]string) (*[]UserCount, error)
		ListRoleByUserUUidArray(ctx context.Context, userUUIDArray *[]string) (*[]RoleListItem, error)
		Insert(ctx context.Context, data *UserToRole) (sql.Result, error)
		FindOne(ctx context.Context, userUuid string, roleUuid string) (*UserToRole, error)
		Update(ctx context.Context, newData *UserToRole) error
		Delete(ctx context.Context, id int64) error
		TransDeleteByUserUUID(ctx context.Context, session sqlx.Session, userUUID string, updateTime time.Time) error
		TransAddRoleUUIDArrayByUserUUID(ctx context.Context, session sqlx.Session, userUUID string, roleUUIDArray []string, updateTime time.Time) error
	}

	defaultUserToRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserToRole struct {
		Id         int64     `db:"id"`
		UserUuid   string     `db:"user_uuid"`
		RoleUuid   string     `db:"role_uuid"`
		IsDelete   int64     `db:"is_delete"` // 是否删除: 0正常, 1删除
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}

	FindUserUUIDArrByRoleUuidArgs struct {
		RoleUUID string
		Page int64
		PageSize int64
	}

	UserCount struct {
		RoleUuid   string     `db:"role_uuid"`
		Count      int64     `db:"count"`
	}

	RoleListItem struct {
		UserUuid   string    `db:"user_uuid"`
		RoleUUID   string    `db:"role_uuid"`
		RoleName   string    `db:"role_name"`
	}
)

func newUserToRoleModel(conn sqlx.SqlConn) *defaultUserToRoleModel {
	return &defaultUserToRoleModel{
		conn:  conn,
		table: "`user_to_role`",
	}
}

func (m *defaultUserToRoleModel) CountUserUUIDArrByRoleUuid(ctx context.Context, roleUUID string) (count int64, err error) {
	query := fmt.Sprintf("select count(*) as count from %s where `is_delete` = 0 and `role_uuid` = ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	err = stmt.QueryRowCtx(ctx, &count, roleUUID)

	return
}

func (m *defaultUserToRoleModel)FindUserUUIDArrByRoleUuid(ctx context.Context, args *FindUserUUIDArrByRoleUuidArgs) (userUUIDArray *[]string, err error) {
	query := fmt.Sprintf("select `user_uuid` from %s where `is_delete` = 0 and `role_uuid` = ? limit ? offset ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	userUUIDArray = new([]string)
	offset:= (args.Page - 1) * args.PageSize
	err = stmt.QueryRowsCtx(ctx, userUUIDArray, args.RoleUUID, args.PageSize, offset)

	return
}

func (m *defaultUserToRoleModel) CountUserGroupByRoleUuid(ctx context.Context, roleUUIDArray *[]string) (resp *[]UserCount, err error) {
	if *roleUUIDArray == nil || len(*roleUUIDArray) == 0 {
		return
	}

	uuids:= "("
	var placeholder []interface{}
	for i,roleUUID := range *roleUUIDArray {
		if i >0 {
			uuids += ", "
		}
		placeholder = append(placeholder, roleUUID)
		uuids += "?"
	}
	uuids += ")"
	query := fmt.Sprintf("select `role_uuid`, count(*) as count from %s where `is_delete` = 0 and `role_uuid` in %s group by `role_uuid`", m.table, uuids)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}

	resp = new([]UserCount)
	err = stmt.QueryRowsCtx(ctx, resp, placeholder...)
	return
}

func (m *defaultUserToRoleModel)ListRoleByUserUUidArray(ctx context.Context, userUUIDArray *[]string) (resp *[]RoleListItem, err error) {
	if *userUUIDArray == nil || len(*userUUIDArray) == 0 {
		return
	}

	uuids:= "("
	var placeholder []interface{}
	for i,userUUID := range *userUUIDArray {
		if i >0 {
			uuids += ", "
		}
		placeholder = append(placeholder, userUUID)
		uuids += "?"
	}
	uuids += ")"
	// 表名是否要用*defaultRoleModel.tableName()来获取？
	query := fmt.Sprintf("select utr.`user_uuid`, r.`role_uuid`, r.`role_name` from %s as utr right join role as r on utr.`role_uuid` = r.`role_uuid` where utr.`is_delete` = 0 and r.`is_delete` = 0 and utr.`user_uuid` in %s", m.table, uuids)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return
	}

	resp = new([]RoleListItem)
	err = stmt.QueryRowsCtx(ctx, resp, placeholder...)
	return
}

func (m *defaultUserToRoleModel) FindRoleUUIDArrByUserUuid(ctx context.Context, userUuid string) (roleUUIDArray *[]string, err error) {
	query := fmt.Sprintf("select `role_uuid` from %s where `is_delete` = 0 and `user_uuid` = ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	roleUUIDArray = new([]string)
	err = stmt.QueryRowsCtx(ctx, roleUUIDArray, userUuid)

	return
}

func (m *defaultUserToRoleModel) Delete(ctx context.Context, id int64)  (err error) {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err !=nil {
		return
	}
	
	_, err = stmt.ExecCtx(ctx, id)
	return
}

func (m *defaultUserToRoleModel) FindOne(ctx context.Context, userUuid string, roleUuid string) (resp *UserToRole, err error) {
	query := fmt.Sprintf("select %s from %s where `user_uuid` = ? and `role_uuid` = ? limit 1", userToRoleRows, m.table)
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err!=nil {
		return
	}
	resp = new(UserToRole)
	err = stmt.QueryRowCtx(ctx, resp, userUuid, roleUuid)
	if err == sqlc.ErrNotFound {
		err = ErrNotFound
	}

	return
}

func (m *defaultUserToRoleModel) Insert(ctx context.Context, data *UserToRole) (res sql.Result, err error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userToRoleRowsExpectAutoSet)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err !=nil {
		return
	}
	res, err = stmt.ExecCtx(ctx, data.UserUuid, data.RoleUuid, data.IsDelete)
	return
}

func (m *defaultUserToRoleModel) Update(ctx context.Context, newData *UserToRole) (err error) {
	query := fmt.Sprintf("update %s set `is_delete` = ? where `user_uuid` = ? and `role_uuid` = ?", m.table)
	stmt, err:= m.conn.PrepareCtx(ctx, query)
	if err !=nil {
		return
	}
	res, err := stmt.ExecCtx(ctx, newData.IsDelete, newData.UserUuid, newData.RoleUuid)
	if err !=nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err !=nil{
		return
	}

	if rowsAffected == 0 {
		err = errors.New("no rows affected")
	}

	return
}

// 提供给logic开启事务用
func (m *defaultUserToRoleModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func (ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultUserToRoleModel) TransDeleteByUserUUID(ctx context.Context, session sqlx.Session, userUUID string, updateTime time.Time) (err error) {
	query := fmt.Sprintf("update %s set `is_delete`=1, `update_time`= ? where `user_uuid` = ?", m.table)
	stmt, err:= session.PrepareCtx(ctx, query)
	if err != nil{
		return
	}
	_, err = stmt.ExecCtx(ctx, updateTime, userUUID)
	return
}

func (m *defaultUserToRoleModel) TransAddRoleUUIDArrayByUserUUID(ctx context.Context, session sqlx.Session, userUUID string, roleUUIDArray []string, updateTime time.Time) (err error) {
	fields := ""
	var placeholder []interface{}
	for i, roleUUID:= range roleUUIDArray {
		if i > 0 {
			fields += ", "
		}
		fields += "(?, ?, ?)"
 		placeholder = append(placeholder, userUUID, roleUUID, updateTime)
	}
	
	query := fmt.Sprintf("insert into %s(%s) values %s on duplicate key update `is_delete`= 0", m.table, userToRoleInsertFeilds, fields)
	stmt, err:= session.PrepareCtx(ctx, query)
	if err != nil{
		return
	}
	_, err = stmt.ExecCtx(ctx, placeholder...)
	return
}
