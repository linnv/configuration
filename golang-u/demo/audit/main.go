package main

import "fmt"

type Auditer interface {
	Audit() error
}

var AuditList map[string]Auditer

type A struct {
	A int
	Auditer
}

func (this A) Audit() error {
	fmt.Printf("this.A: approve  works %v\n", this.A)
	return nil
}

type B struct {
	B int
}

func (this B) Audit() error {
	fmt.Printf("this.B: disapprove works %v\n", this.B)
	return nil
}

func init() {
	println("init")
	AuditList = map[string]Auditer{}
}

func Regisgter(auditType string, auditInstance Auditer) {
	AuditList[auditType] = auditInstance
}

func Executer(auditType string) error {
	return AuditList[auditType].Audit()
	// audit, ok := AuditList[auditType]
	// if ok {
	// 	if err := audit.Audit(); err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("runs normally\n")
	// 	return nil
	// }
	// return errors.New("unexpected get instance fail")
}

func main() {
	a := A{A: 1000}
	b := B{B: 9999}
	Regisgter("approve", a)
	Regisgter("disapprove", b)
	if err := Executer("approve"); err != nil {
		panic(err.Error())
	}
	return
}
