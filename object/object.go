package object

import (
	"fmt"
	"time"
)

type Method string

const (
	MethodGet    Method = "get"
	MethodSet    Method = "set"
	MethodFind   Method = "find"
	MethodDelete Method = "delete"
	MethodWatch  Method = "watch"
)

type Key struct {
	Type      string
	Namespace string
	Name      string
}

func (k Key) String() string {
	return fmt.Sprintf("{Type: %s, Namespace: %s, Name: %s}", k.Type, k.Namespace, k.Name)
}

func (k *Key) Equal(key *Key) bool {
	if k.Type != key.Type {
		return false
	} else if k.Namespace != key.Namespace {
		return false
	} else if k.Name != key.Name {
		return false
	}
	return true
}

func (k *Key) Match(key *Key) bool {
	if k.Type != "" && k.Type != key.Type {
		return false
	} else if k.Namespace != "" && k.Namespace != key.Namespace {
		return false
	} else if k.Name != "" && k.Name != key.Name {
		return false
	}
	return true
}

type Meta struct {
	CreationTime time.Time
	Owner        string
	Method       Method
}

type Object struct {
	Meta  *Meta
	Key   *Key
	Value string
}

func (o Object) String() string {
	val := ""
	if len(o.Value) > 10 {
		val = fmt.Sprintf("'%s...'", o.Value[0:10])
	} else {
		val = fmt.Sprintf("'%s'", o.Value)
	}

	return fmt.Sprintf("{Key: %v, Value: %s}", o.Key, val)
}
