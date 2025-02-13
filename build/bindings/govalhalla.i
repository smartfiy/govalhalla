%module govalhalla

// Include standard SWIG interfaces


%{
#include <utility>
#include <string>
#include <memory>
#include <functional>
#include "valhalla/tyr/actor.h"
#include <valhalla/baldr/graphreader.h>
#include <valhalla/proto/api.pb.h>
#include "valhalla/tyr/govalhalla_actor.h"
%}

%include <std_string.i>
%include <std_pair.i>
// Handle std::function
%typemap(gotype) const std::function<void()>* "uintptr"
%typemap(in,numinputs=0) const std::function<void()>* {
    $1 = nullptr;
}

// Handle valhalla::Api
%typemap(gotype) valhalla::Api* "uintptr"
%typemap(in,numinputs=0) valhalla::Api* {
    $1 = nullptr;
}

// Create the response pair template
%template(ResponsePair) std::pair<std::string, std::string>;

// Include the actor header
%include "govalhalla_actor.h"

// CGO preamble
%insert(go_begin) %{
/*
#cgo LDFLAGS: -L -lgovalhalla
*/
%}