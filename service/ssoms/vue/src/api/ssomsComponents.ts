// Code generated by goctl. DO NOT EDIT.

export interface Option {
	label: string
	value: string
	extra?: string
}

export interface OptionWithDisabled {
	label: string
	value: string
	extra?: string
	disabled: boolean
}

export interface Statistic {
	month: string
	type: string
	value: number
}

export interface StatisticReply {
	roleAmount: number // 角色数量
	userAmount: number // 用户数量
	systemAmount: number // 系统数量
	menuAmount: number // 菜单数量
	actionAmount: number // 操作数量
	permissionAmount: number // 授权数量
	statistics: Array<Statistic> // 菜单&操作统计
}

export interface UserListReq {
}

export interface UserListReqParams {
	name: string
	mobile: string
	page: number
	pageSize: number
}

export interface User {
	uuid: string
	name: string
	mobile: string
	avatar: string
	gender: number
	status: number
	roles: Array<Option>
}

export interface UserListReply {
	total: number
	list: Array<User>
}

export interface UserForm {
	uuid?: string
	name: string
	mobile: string
	password?: string
	avatar: string
	gender: number
	birth: string
	introduction: string
	status: number
}

export interface AddUserReply {
	uuid: string
}

export interface UserDetailReq {
}

export interface UserDetailReqParams {
	uuid: string
}

export interface UpdateUserReq {
	name: string
	mobile: string
	password?: string
	avatar: string
	gender: number
	birth: string
	introduction: string
	status: number
}

export interface UpdateUserReqParams {
	uuid: string
}

export interface UpdateUserReply {
	success: boolean
}

export interface DeleteUserReq {
}

export interface DeleteUserReqParams {
	uuid: string
}

export interface DeleteUserReply {
	success: boolean
}

export interface UserFilterOptionsReq {
}

export interface UserFilterOptionsReqParams {
	name: string
}

export interface UserFilterOptionsReply {
	options: Array<Option>
}

export interface AssignedRolesReq {
}

export interface AssignedRolesReqParams {
	uuid: string
}

export interface AssignedRolesReply {
	uuid: string
	name: string
	assigned: Array<Option>
	options: Array<Option>
}

export interface AssignRoleReq {
	uuid: string
	roleUUIDArray: Array<string>
}

export interface AssignRoleReply {
	success: boolean
}

export interface UserPermissionsReq {
	uuid: string
	type: number // 类型: 角色
}

export interface UserPermissionsReply {
	system: Array<Option>
}

export interface AvatarUploadReq {
	avatar: string
}

export interface AvatarUploadReply {
	success: boolean
}

export interface InfoEditReq {
	introduction: string
}

export interface InfoEditReply {
	success: boolean
}

export interface PasswordResetReq {
	oldPassword: string
	password: string
	confirmPassword: string
}

export interface PasswordResetReply {
	success: boolean
}

export interface RoleListReq {
}

export interface RoleListReqParams {
	roleName: string
	page: number
	pageSize: number
}

export interface Role {
	roleUUID: string // 角色uuid
	roleName: string // 角色名称
	summary: string // 角色概述
	inheritances: Array<Option> // 继承的角色
	usersAmount: number // 拥有的用户数量
}

export interface RoleListReply {
	total: number
	list: Array<Role>
}

export interface RoleForm {
	roleUUID?: string
	roleName: string
	summary: string
}

export interface AddRoleReply {
	roleUUID: string
}

export interface RoleDetailReq {
}

export interface RoleDetailReqParams {
	roleUUID: string
}

export interface UpdateRoleReq {
	roleName: string
	summary: string
}

export interface UpdateRoleReqParams {
	roleUUID: string
}

export interface UpdateRoleReply {
	success: boolean
}

export interface DeleteRoleReq {
}

export interface DeleteRoleReqParams {
	roleUUID: string
}

export interface DeleteRoleReply {
	success: boolean
}

export interface OptionsReply {
	options: Array<Option>
}

export interface AssignedUsersReq {
}

export interface AssignedUsersReqParams {
	roleUUID: string
	page: number
	pageSize: number
}

export interface UserOtion {
	uuid: string
	name: string
	mobile: string
	avatar: string
	gender: number
	status: number
	isDelete: number
}

export interface AssignedUsersReply {
	total: number
	list: Array<UserOtion>
	roleName: string
}

export interface AssignUserReq {
	userUUID: string
}

export interface AssignUserReqParams {
	roleUUID: string
}

export interface AssignUserReply {
	success: boolean
}

export interface DeassignUserReq {
	userUUID: string
}

export interface DeassignUserReqParams {
	roleUUID: string
}

export interface DeassignUserReply {
	success: boolean
}

export interface InheritancesReq {
}

export interface InheritancesReqParams {
	roleUUID: string
}

export interface InheritancesReply {
	roles: Array<Option>
}

export interface AddInheritanceReq {
	extendedRoleUUIDArray: Array<string>
}

export interface AddInheritanceReqParams {
	roleUUID: string
}

export interface AddInheritanceReply {
	success: boolean
}

export interface SystemListReq {
}

export interface SystemListReqParams {
	systemCode: string
	systemName: string
}

export interface System {
	uuid: string
	systemCode: string
	systemName: string
	domain: string
	sort: number
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	extra?: string // 扩展字段
	icon: string // 图标
	status: number
}

export interface SystemListReply {
	list: Array<System>
}

export interface SystemForm {
	uuid?: string
	systemCode: string
	systemName: string
	domain: string
	sort: number
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	icon: string // 图标
	status: number
}

export interface AddSystemReply {
	uuid: string
}

export interface SystemDetailReq {
}

export interface SystemDetailReqParams {
	uuid: string
}

export interface UpdateSystemReq {
	systemCode: string
	systemName: string
	domain: string
	sort: number
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	icon: string // 图标
	status: number
}

export interface UpdateSystemReqParams {
	uuid: string
}

export interface UpdateSystemReply {
	success: boolean
}

export interface DeleteSystemReq {
}

export interface DeleteSystemReqParams {
	uuid: string
}

export interface DeleteSystemReply {
	success: boolean
}

export interface ObjectListReq {
}

export interface ObjectListReqParams {
	topKey: string // 传systemCode
	objectName: string
	key: string
}

export interface Object {
	uuid: string
	objectName: string
	identifier: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1操作对象, 2菜单，3操作(接口)
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	extra?: string // 扩展字段
	icon: string // 图标
	status: number
	pUUID?: string
	children?: Array<Object>
}

export interface ObjectOption {
	value: string
	label: string
	pUUID?: string
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	children?: Array<ObjectOption> // 子菜单
	apis: Array<ObjectOption> // 操作
}

export interface MenuOptionsReq {
}

export interface MenuOptionsReqParams {
	excludeHide: boolean // 是否排查隐藏菜单
}

export interface MenuOption {
	value: string
	label: string
	pUUID?: string
	children?: Array<MenuOption> // 子菜单
}

export interface MenuOptionsReply {
	list: Array<MenuOption>
}

export interface ObjectListReply {
	list: Array<Object>
}

export interface ObjectForm {
	uuid?: string
	objectName: string
	identifier: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	extra?: string // 扩展字段
	icon: string // 图标
	status: number
	pUUID?: string
	topKey?: string // 传systemCode
}

export interface AddObjectReply {
	uuid: string
}

export interface ImportObjectReq {
	uuid: string
	objectName: string
	identifier?: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1操作对象, 2模块，3菜单组，4菜单，5操作(接口)
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	extra?: string // 扩展字段
	icon?: string // 图标
	status: number
	pUUID?: string
	topKey: string // 传systemCode
}

export interface ImportObjectReply {
	status: string // 导入结果状态
	msg: string // 导入结果信息
}

export interface ObjectDetailReq {
}

export interface ObjectDetailReqParams {
	uuid: string
}

export interface UpdateObjectReq {
	objectName: string
	identifier: string
	key: string // 操作对象的systemCode, 菜单的path, 操作的uri
	sort: number
	type: number // 类型, 1系统, 2菜单, 3操作(接口)
	subType: number // 子类型, 菜单时: (1菜单，2菜单组, 3隐藏菜单)
	extra?: string // 扩展字段
	icon: string // 图标
	status: number
	pUUID?: string
	topKey?: string // 更新时无法修改该值, 但是更新前需要使用该值校验对象是否已存在
}

export interface UpdateObjectReqParams {
	uuid: string
}

export interface UpdateObjectReply {
	success: boolean
}

export interface DeleteObjectReq {
}

export interface DeleteObjectReqParams {
	uuid: string
}

export interface DeleteObjectReply {
	success: boolean
}

export interface RoleOperationsReq {
}

export interface RoleOperationsReqParams {
	topKey: string // 传systemCode
}

export interface RoleOperationsReply {
	list?: Array<ObjectOption>
}

export interface RolePermissionsReq {
}

export interface RolePermissionsReqParams {
	roleUUID: string
	topKey: string
}

export interface RolePermissionsReply {
	roleName: string
	menuUUIDArray: Array<string>
	actionUUIDArray: Array<string>
}

export interface GrantReq {
	topKey: string
	menuUUIDArray: Array<string>
	actionUUIDArray: Array<string>
}

export interface GrantReqParams {
	roleUUID: string
}

export interface GrantReply {
	success: boolean
}

