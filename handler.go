package main

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
  "strings"
)

// an useful struct to create a logger. It helps to make the dependencies with clear logic of handlers
type Handler struct {
  l *log.Logger
}

func NewHandler(l *log.Logger) *Handler{
  return &Handler{l}
}

func getAllSchedulesHandler(w http.ResponseWriter, r *http.Request){
  lp,err := getSchedules()
  if err != nil{
    http.Error(w,"Failed to get every order",http.StatusInternalServerError)
  }
  lp_json,err := json.Marshal(lp)
  if err != nil{
    http.Error(w,"failed to marshal json",http.StatusInternalServerError)
  }
  w.WriteHeader(http.StatusOK)
  w.Write(lp_json)
}  
func deleteScheduleByIdHandler(scheduleId uint,w http.ResponseWriter, r *http.Request){
     err := deleteSchedule(scheduleId)
     if err != nil{
        http.Error(w,"failed to delete an Order Info by Id",http.StatusInternalServerError)
     }
     w.WriteHeader(http.StatusOK)

   }

// this is a main handler, it will give other handlers depend of a method(request)
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  // get an url path in the beginning
  url := r.URL.Path
  
  if r.Method == http.MethodGet{
    if url == "/"{
    // if url doesn't contain any id then get every order
    getAllSchedulesHandler(w,r)
    }
    return
  }
  if r.Method == http.MethodPost{
    return
  }
  if r.Method == http.MethodPost{
    return
  }
  if r.Method == http.MethodDelete{
    if url == "/"{
      /// if url doesn't contain any id then return error
        http.Error(w,"Method  delete but you didn't enter any id",http.StatusBadRequest)
        return
    }  else {
    // otherwise delete an order by id
      tmp := strings.Trim(url,"/")
      orderIdConv,err := strconv.ParseUint(tmp,10,64)
      if err != nil{
        http.Error(w,"failed to convert URL Path in uint",http.StatusInternalServerError)
      }
      deleteScheduleByIdHandler(uint(orderIdConv),w,r)
    return

  }
 }
}

// func putNewOrderHandler(orderId uint,productId uint,w http.ResponseWriter, r *http.Request){
//   err := addProduct(orderId,productId)
//      if err != nil{
//         http.Error(w,"error in put new order handler",http.StatusInternalServerError)
//      }
//      w.WriteHeader(http.StatusOK)
// }
// // a handler to delete an order by Id
     


// }
// /// a handler to get an order by id 
// func getOrderByIdHandler(OrderId uint,w http.ResponseWriter, r *http.Request){
//      lp,err := getInfoOrderById(OrderId)
//      if err != nil{
//         http.Error(w,"failed to get an Order Info by Id",http.StatusInternalServerError)
//      }
//      lp_json,err := json.Marshal(lp)
//      if err != nil{
//          http.Error(w,"failed to marshal json",http.StatusInternalServerError)
//       }
//      w.WriteHeader(http.StatusOK)
//       w.Write(lp_json)
// }
// handler to get information about every order, if client didn't give any information about orders and used a method get. It's like a default response by a server

