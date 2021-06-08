package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "time"
  // "strconv"
  // "log"
  // "github.com/streadway/amqp"
  // "encoding/json"
  // "os"
)


//struct for a database connection

type dbConnection struct {
  DB *gorm.DB
}

  type Model struct {
    ID        uint       `gorm:"primary_key auto_increment:true;column:id" json:"id"`
 CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
 UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
 DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
type Schedule struct {
  Model
  name string 
  type_name string
  start_date time.Time
  interval_in_days uint
  frequency uint
}
type ScheduleDaily struct { Model
  ScheduleID uint `gorm:"references:Schedule"`
  Time_of_day time.Time
  Start_date time.Time
  end_date time.Time 
}
type ScheduleMonthly struct {
  Model
  ScheduleID uint `gorm:"references:Schedule"`
  Day_of_month time.Time
  Start_date time.Time
  End_date time.Time 
}
type ScheduleSpecific struct {
  Model
  ScheduleID uint `gorm:"references:Schedule"`
  Start_date time.Time
  End_date time.Time 
}

type InfoSchedule struct {
  Name string
  Type_name string
  Interval_in_days uint
  Start_date time.Time
  end_date time.Time 
}
  
//func to connect to a database 
func connectToDataBase()(db *gorm.DB,err error){
  dsn := "root:root@tcp(127.0.0.1:3306)/schedule?charset=utf8mb4&parseTime=True&loc=Local"
  db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil,err
  }
  return db,nil
}
// get all schedules
func getSchedules() (schedule []*Schedule,err error){
  db,err := connectToDataBase()
  if err != nil{
    return nil,err
  }
 db.Model(&Schedule{}).Select("*").Scan(&schedule)
 db.Model(&ScheduleDaily{}).Select("*").Scan(&schedule)


  return schedule,nil
}
// delete an order 
func deleteSchedule(ScheduleId uint)(err error){
  db,err := connectToDataBase()
  if err != nil{
    return err
  }
  db.Delete(&Schedule{},ScheduleId)

  return nil
}

// add new schedule
func addSchedule(schedule *[]Schedule) (err error) {
  db,err := connectToDataBase()
  if err != nil{
    return err
  }
  result := db.Create(&schedule)
  if result.Error != nil{
    return result.Error
  }
  return nil 
}
// update schedule
func updateSchedule(schedule_id uint,Name string,Type_name string,Start_date time.Time,Interval_in_days uint,Frequency uint) (err error) {
  db,err := connectToDataBase()
  if err != nil{
    return err
  }
  var schedule Schedule
  schedule.ID = schedule_id
  db.First(&schedule)

  return nil
}

