package main

func yonghu(stunum int,username string)(Choose,error){
	var choose Choose
	err := db.Where(Choose{Stunum: stunum}).Attrs(Choose{Username: username}).FirstOrCreate(&choose).Error//如果没有该记录就自动增加
	return choose,err
}

func queryclass(dacourse_num string)(Class,error){
	var class Class
	err := db.Model(&Class{}).Where("Dacourse_num = ?",dacourse_num).First(&class).Error // 在事务中做一些数据库操作（从这一点使用'tx'，而不是'db'）
	return class,err
}

func addclass(choose Choose,class Class)(error){
	var timeday = choose.Timeday+" # "+class.Day
	var timelesson = choose.Timelesson+" # "+class.Lesson
	var date = choose.Date+" "+class.Dacourse_num

	if choose.Timeday == ""{
		timeday = class.Day
		timelesson = class.Lesson
	}

	err := db.Model(&Choose{}).Where("Stunum = ?",choose.Stunum).Updates(Choose{Date:date,Timeday:timeday,Timelesson:timelesson}).Error
	return err
}