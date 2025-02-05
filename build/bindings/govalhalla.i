%module govalhalla
%include <std_string.i>
%include <std_pair.i>


%{
#include <utility>
#include <string>
#include <memory>
#include <functional>
// #include <valhalla/tyr/actor.h>
// #include <valhalla/baldr/graphreader.h>
// #include <valhalla/proto/api.pb.h>
#include "govalhalla_actor.h"
// using namespace valhalla;
// using namespace valhalla::baldr;
// using namespace valhalla::tyr;

%}


// %ignore tile_extract_t;
// %ignore SearchFilter;
// %ignore rapidjson::get_optional;
// %ignore std::hash<valhalla::baldr::Location>;

// Expose std::pair<std::string, std::string> to hold the result 

typedef std::pair<std::string, std::string> ResponsePair;

%template(ResponsePair) std::pair<std::string, std::string>;
// %import <valhalla/baldr/graphreader.h>
// %import "valhalla/tyr/actor.h" 
%include "govalhalla_actor.h" 


// %typemap(in) const std::function<void()>* {
//     $1 = nullptr;
// }


// Expose the ActorWrapper class to 
// %typemap(type) std::uintptr_t "uintptr"   // Map to 's uintptr type
// %typemap(in) std::uintptr_t {
//     $1 = static_cast<std::uintptr_t>($input);
// }
// %typemap(out) std::uintptr_t {
//     $result = static_cast<std::uintptr_t>($1);
// }
// class actor_t {
// // public:
    

// };
// }}
// valhalla::tyr::
// %include "valhalla/proto/api.pb.h"
// Declare the namespaces and class before extending
// Include necessary headers for SWIG parsing
// %import <valhalla/proto/api.pb.h>
// %import <valhalla/baldr/graphreader.h>
// %include <boost/property_tree/ptree.hpp>


