/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.3.0
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

// source: govalhalla.i

package govalhalla

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stddef.h>
#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
typedef _gostring_ swig_type_10;
typedef _gostring_ swig_type_11;
typedef _gostring_ swig_type_12;
typedef _gostring_ swig_type_13;
typedef _gostring_ swig_type_14;
typedef _gostring_ swig_type_15;
typedef _gostring_ swig_type_16;
typedef _gostring_ swig_type_17;
typedef _gostring_ swig_type_18;
typedef _gostring_ swig_type_19;
typedef _gostring_ swig_type_20;
typedef _gostring_ swig_type_21;
typedef _gostring_ swig_type_22;
typedef _gostring_ swig_type_23;
typedef _gostring_ swig_type_24;
typedef _gostring_ swig_type_25;
typedef _gostring_ swig_type_26;
typedef _gostring_ swig_type_27;
typedef _gostring_ swig_type_28;
typedef _gostring_ swig_type_29;
typedef _gostring_ swig_type_30;
typedef _gostring_ swig_type_31;
typedef _gostring_ swig_type_32;
typedef _gostring_ swig_type_33;
typedef _gostring_ swig_type_34;
typedef _gostring_ swig_type_35;
typedef _gostring_ swig_type_36;
typedef _gostring_ swig_type_37;
typedef _gostring_ swig_type_38;
typedef _gostring_ swig_type_39;
typedef _gostring_ swig_type_40;
typedef _gostring_ swig_type_41;
typedef _gostring_ swig_type_42;
typedef _gostring_ swig_type_43;
typedef _gostring_ swig_type_44;
extern void _wrap_Swig_free_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_govalhalla_19f8d4121edc110d(swig_intgo arg1);
extern uintptr_t _wrap_new_ResponsePair__SWIG_0_govalhalla_19f8d4121edc110d(void);
extern uintptr_t _wrap_new_ResponsePair__SWIG_1_govalhalla_19f8d4121edc110d(swig_type_1 arg1, swig_type_2 arg2);
extern uintptr_t _wrap_new_ResponsePair__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern void _wrap_ResponsePair_first_set_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_3 arg2);
extern swig_type_4 _wrap_ResponsePair_first_get_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern void _wrap_ResponsePair_second_set_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_5 arg2);
extern swig_type_6 _wrap_ResponsePair_second_get_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern void _wrap_delete_ResponsePair_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern uintptr_t _wrap_new_Actor__SWIG_0_govalhalla_19f8d4121edc110d(swig_type_7 arg1, _Bool arg2);
extern uintptr_t _wrap_new_Actor__SWIG_1_govalhalla_19f8d4121edc110d(swig_type_8 arg1);
extern uintptr_t _wrap_new_Actor__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern void _wrap_Actor_Cleanup_govalhalla_19f8d4121edc110d(uintptr_t arg1);
extern uintptr_t _wrap_Actor_Act__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Act__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Actor_Route__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_9 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Route__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_10 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Route__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_11 arg2);
extern uintptr_t _wrap_Actor_Locate__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_12 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Locate__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_13 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Locate__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_14 arg2);
extern uintptr_t _wrap_Actor_Matrix__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_15 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Matrix__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_16 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Matrix__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_17 arg2);
extern uintptr_t _wrap_Actor_OptimizedRoute__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_18 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_OptimizedRoute__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_19 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_OptimizedRoute__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_20 arg2);
extern uintptr_t _wrap_Actor_Isochrone__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_21 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Isochrone__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_22 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Isochrone__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_23 arg2);
extern uintptr_t _wrap_Actor_TraceRoute__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_24 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_TraceRoute__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_25 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_TraceRoute__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_26 arg2);
extern uintptr_t _wrap_Actor_TraceAttributes__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_27 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_TraceAttributes__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_28 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_TraceAttributes__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_29 arg2);
extern uintptr_t _wrap_Actor_Height__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_30 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Height__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_31 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Height__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_32 arg2);
extern uintptr_t _wrap_Actor_TransitAvailable__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_33 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_TransitAvailable__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_34 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_TransitAvailable__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_35 arg2);
extern uintptr_t _wrap_Actor_Expansion__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_36 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Expansion__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_37 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Expansion__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_38 arg2);
extern uintptr_t _wrap_Actor_Centroid__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_39 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Centroid__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_40 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Centroid__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_41 arg2);
extern uintptr_t _wrap_Actor_Status__SWIG_0_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_42 arg2, uintptr_t arg3, uintptr_t arg4);
extern uintptr_t _wrap_Actor_Status__SWIG_1_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_43 arg2, uintptr_t arg3);
extern uintptr_t _wrap_Actor_Status__SWIG_2_govalhalla_19f8d4121edc110d(uintptr_t arg1, swig_type_44 arg2);
extern void _wrap_delete_Actor_govalhalla_19f8d4121edc110d(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


func getSwigcptr(v interface { Swigcptr() uintptr }) uintptr {
	if v == nil {
		return 0
	}
	return v.Swigcptr()
}


type _ sync.Mutex

//export cgo_panic__govalhalla_19f8d4121edc110d
func cgo_panic__govalhalla_19f8d4121edc110d(p *byte) {
	s := (*[1024]byte)(unsafe.Pointer(p))[:]
	for i, b := range s {
		if b == 0 {
			panic(string(s[:i]))
		}
	}
	panic(string(s))
}


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_govalhalla_19f8d4121edc110d(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrResponsePair uintptr

func (p SwigcptrResponsePair) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrResponsePair) SwigIsResponsePair() {
}

func NewResponsePair__SWIG_0() (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_new_ResponsePair__SWIG_0_govalhalla_19f8d4121edc110d()))
	return swig_r
}

func NewResponsePair__SWIG_1(arg1 string, arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_new_ResponsePair__SWIG_1_govalhalla_19f8d4121edc110d(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func NewResponsePair__SWIG_2(arg1 ResponsePair) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := getSwigcptr(arg1)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_new_ResponsePair__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewResponsePair(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 0 {
		return NewResponsePair__SWIG_0()
	}
	if argc == 1 {
		return NewResponsePair__SWIG_2(a[0].(ResponsePair))
	}
	if argc == 2 {
		return NewResponsePair__SWIG_1(a[0].(string), a[1].(string))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrResponsePair) SetFirst(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ResponsePair_first_set_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrResponsePair) GetFirst() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_ResponsePair_first_get_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrResponsePair) SetSecond(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ResponsePair_second_set_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrResponsePair) GetSecond() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_ResponsePair_second_get_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func DeleteResponsePair(arg1 ResponsePair) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_ResponsePair_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))
}

type ResponsePair interface {
	Swigcptr() uintptr
	SwigIsResponsePair()
	SetFirst(arg2 string)
	GetFirst() (_swig_ret string)
	SetSecond(arg2 string)
	GetSecond() (_swig_ret string)
}

type SwigcptrActor uintptr

func (p SwigcptrActor) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrActor) SwigIsActor() {
}

func NewActor__SWIG_0(arg1 string, arg2 bool) (_swig_ret Actor) {
	var swig_r Actor
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Actor)(SwigcptrActor(C._wrap_new_Actor__SWIG_0_govalhalla_19f8d4121edc110d(*(*C.swig_type_7)(unsafe.Pointer(&_swig_i_0)), C._Bool(_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func NewActor__SWIG_1(arg1 string) (_swig_ret Actor) {
	var swig_r Actor
	_swig_i_0 := arg1
	swig_r = (Actor)(SwigcptrActor(C._wrap_new_Actor__SWIG_1_govalhalla_19f8d4121edc110d(*(*C.swig_type_8)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func NewActor__SWIG_2(arg1 uintptr) (_swig_ret Actor) {
	var swig_r Actor
	_swig_i_0 := arg1
	swig_r = (Actor)(SwigcptrActor(C._wrap_new_Actor__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewActor(a ...interface{}) Actor {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(string); !ok {
			goto check_1
		}
		return NewActor__SWIG_1(a[0].(string))
	}
check_1:
	if argc == 1 {
		return NewActor__SWIG_2(a[0].(uintptr))
	}
	if argc == 2 {
		return NewActor__SWIG_0(a[0].(string), a[1].(bool))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Cleanup() {
	_swig_i_0 := arg1
	C._wrap_Actor_Cleanup_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrActor) Act__SWIG_0(arg2 Valhalla_Api, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Act__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2))))
	return swig_r
}

func (arg1 SwigcptrActor) Act__SWIG_1(arg2 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Act__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (p SwigcptrActor) Act(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Act__SWIG_1(a[0].(Valhalla_Api))
	}
	if argc == 2 {
		return p.Act__SWIG_0(a[0].(Valhalla_Api), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Route__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Route__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Route__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Route__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_10)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Route__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Route__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_11)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Route(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Route__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Route__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Route__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Locate__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Locate__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_12)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Locate__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Locate__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_13)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Locate__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Locate__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_14)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Locate(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Locate__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Locate__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Locate__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Matrix__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Matrix__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_15)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Matrix__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Matrix__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_16)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Matrix__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Matrix__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_17)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Matrix(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Matrix__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Matrix__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Matrix__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) OptimizedRoute__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_OptimizedRoute__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_18)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) OptimizedRoute__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_OptimizedRoute__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_19)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) OptimizedRoute__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_OptimizedRoute__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_20)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) OptimizedRoute(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.OptimizedRoute__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.OptimizedRoute__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.OptimizedRoute__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Isochrone__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Isochrone__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_21)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Isochrone__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Isochrone__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_22)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Isochrone__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Isochrone__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_23)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Isochrone(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Isochrone__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Isochrone__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Isochrone__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) TraceRoute__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TraceRoute__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_24)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) TraceRoute__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TraceRoute__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_25)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) TraceRoute__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TraceRoute__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_26)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) TraceRoute(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.TraceRoute__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.TraceRoute__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.TraceRoute__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) TraceAttributes__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TraceAttributes__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_27)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) TraceAttributes__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TraceAttributes__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_28)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) TraceAttributes__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TraceAttributes__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_29)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) TraceAttributes(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.TraceAttributes__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.TraceAttributes__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.TraceAttributes__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Height__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Height__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_30)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Height__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Height__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_31)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Height__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Height__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_32)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Height(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Height__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Height__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Height__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) TransitAvailable__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TransitAvailable__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_33)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) TransitAvailable__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TransitAvailable__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_34)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) TransitAvailable__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_TransitAvailable__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_35)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) TransitAvailable(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.TransitAvailable__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.TransitAvailable__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.TransitAvailable__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Expansion__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Expansion__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_36)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Expansion__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Expansion__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_37)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Expansion__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Expansion__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_38)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Expansion(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Expansion__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Expansion__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Expansion__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Centroid__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Centroid__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_39)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Centroid__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Centroid__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_40)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Centroid__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Centroid__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_41)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Centroid(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Centroid__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Centroid__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Centroid__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrActor) Status__SWIG_0(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_, arg4 Valhalla_Api) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Status__SWIG_0_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_42)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Status__SWIG_1(arg2 string, arg3 Std_function_Sl_void_Sp__SP__Sg_) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Status__SWIG_1_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_43)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrActor) Status__SWIG_2(arg2 string) (_swig_ret ResponsePair) {
	var swig_r ResponsePair
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ResponsePair)(SwigcptrResponsePair(C._wrap_Actor_Status__SWIG_2_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0), *(*C.swig_type_44)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrActor) Status(a ...interface{}) ResponsePair {
	argc := len(a)
	if argc == 1 {
		return p.Status__SWIG_2(a[0].(string))
	}
	if argc == 2 {
		return p.Status__SWIG_1(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_))
	}
	if argc == 3 {
		return p.Status__SWIG_0(a[0].(string), a[1].(Std_function_Sl_void_Sp__SP__Sg_), a[2].(Valhalla_Api))
	}
	panic("No match for overloaded function call")
}

func DeleteActor(arg1 Actor) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_Actor_govalhalla_19f8d4121edc110d(C.uintptr_t(_swig_i_0))
}

type Actor interface {
	Swigcptr() uintptr
	SwigIsActor()
	Cleanup()
	Act(a ...interface{}) ResponsePair
	Route(a ...interface{}) ResponsePair
	Locate(a ...interface{}) ResponsePair
	Matrix(a ...interface{}) ResponsePair
	OptimizedRoute(a ...interface{}) ResponsePair
	Isochrone(a ...interface{}) ResponsePair
	TraceRoute(a ...interface{}) ResponsePair
	TraceAttributes(a ...interface{}) ResponsePair
	Height(a ...interface{}) ResponsePair
	TransitAvailable(a ...interface{}) ResponsePair
	Expansion(a ...interface{}) ResponsePair
	Centroid(a ...interface{}) ResponsePair
	Status(a ...interface{}) ResponsePair
}


type SwigcptrStd_function_Sl_void_Sp__SP__Sg_ uintptr
type Std_function_Sl_void_Sp__SP__Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_function_Sl_void_Sp__SP__Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrValhalla_Api uintptr
type Valhalla_Api interface {
	Swigcptr() uintptr;
}
func (p SwigcptrValhalla_Api) Swigcptr() uintptr {
	return uintptr(p)
}

