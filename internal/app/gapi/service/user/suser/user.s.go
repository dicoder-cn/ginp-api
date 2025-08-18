package suser

import (
	"fmt"
	"ginp-api/internal/app/gapi/entity"
	"ginp-api/internal/app/gapi/model/user/muser"
	scommon "ginp-api/internal/app/gapi/service/system/common"

	"ginp-api/internal/db/mysql"
	"ginp-api/pkg/where"
)

var User *muser.Model

func Model() *muser.Model {
	if User == nil {
		dbRead := mysql.GetReadDb()
		dbWrite := mysql.GetWriteDb()
		User = muser.NewModel(dbRead, dbWrite)
	}
	return User
}

// 修改信息
func UpdateUserInfo(user *entity.User, emailCode string) error {
	if user == nil || user.ID == 0 {
		return fmt.Errorf("user is nil or user.id is 0")
	}
	//查询出旧的信息
	oldUser, err := Model().FindOne(where.New(muser.FieldID, "=", user.ID).Conditions())
	if oldUser == nil || oldUser.ID <= 0{
		return fmt.Errorf("get user failed: %v", err)
	}

	//如果修改了用户名
	if user.Username != "" && user.Username != oldUser.Username {
		//查询用户名是否存在
		oldUsernameUser, _ := Model().FindOne(where.New(muser.FieldUsername, "=", user.Username).Conditions())
		if oldUsernameUser != nil && oldUsernameUser.ID > 0 {
			return fmt.Errorf("username is exist")
		}
	}


	//如果密码不为空
	var passwordChanged bool
	if user.Password != "" {
		// 验证新密码是否与旧密码不同
		passwordValid, _ := VerifyPassword(user.Password, oldUser.Password)
		if err != nil || !passwordValid {
			passwordChanged = true
		}
	}
	
	if passwordChanged || (user.Email != "" && user.Email != oldUser.Email) {
		//如果邮箱验证码不为空
		if emailCode != "" {
			//验证邮箱验证码
			err := scommon.EmailInstance.VerifyCode(user.Email, emailCode)
			if err {
				return fmt.Errorf("email code verify failed: %v", err)
			}
		} else {
			return fmt.Errorf("email code is empty")
		}

		//修改密码
		if passwordChanged {
			hashedPassword, err := HashPassword(user.Password)
			if err != nil {
				return fmt.Errorf("failed to hash password: %w", err)
			}
			user.Password = hashedPassword
		}

		//修改邮箱
		if user.Email != "" && user.Email != oldUser.Email {	
			//查询新邮箱是否存在
			oldEmailUser, _ := Model().FindOne(where.New(muser.FieldEmail, "=", user.Email).Conditions())
			if oldEmailUser != nil && oldEmailUser.ID > 0 {
				return fmt.Errorf("email is exist")
			}
		}
	}


	wheres := where.New(muser.FieldID, "=", user.ID).Conditions()
	err = Model().Update(wheres, user)
	if err != nil {
		return fmt.Errorf("update user failed: %v", err.Error())
	}
	//修改成功
	return nil
}
