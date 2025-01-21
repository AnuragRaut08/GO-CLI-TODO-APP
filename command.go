package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type CmdFlags struct {
    Add    string
    Del    int
    Edit   string
    Toggle int
    List   bool
}

func NewCmdFlags() *CmdFlags {
    cf := CmdFlags{}
   
    flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
    flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
    flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
    flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
    flag.BoolVar(&cf.List, "list", false, "List all todos")
   
    flag.Parse()
   
    return &cf
}
