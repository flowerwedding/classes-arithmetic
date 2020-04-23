package main

func CreateAnimals(stunum int,username string,dacourse_num string) error {
	var tx = db.Begin() // 注意，一旦你在一个事务中，使用tx作为数据库句柄

	var choose Choose
	var class Class
	var err error

	if choose,err = yonghu(stunum,username); err != nil{//我自己的课表的时间段
		tx.Rollback()// 发生错误时回滚事务
		return err
	}

	if class,err = queryclass(dacourse_num); err != nil{//要选的课的时间段
		tx.Rollback()
		return err
	}

	if flag := compare(choose,class); flag !=  0{//比较时间段是否合适
		tx.Rollback()
		return err
	}

	if err = addclass(choose,class); err != nil{//合适的话增加该课程
		tx.Rollback()
		return err
	}

	tx.Commit()// 或提交事务
	return nil
}